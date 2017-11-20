package main

import (
	"flag"
	"fmt"
	"github.com/BonnierNews/logstash_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	scrapeDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: collector.Namespace,
			Subsystem: "exporter",
			Name:      "scrape_duration_seconds",
			Help:      "logstash_exporter: Duration of a scrape job.",
		},
		[]string{"collector", "result"},
	)
)

type LogstashCollector struct {
	collectors map[string]collector.Collector
}

func NewLogstashCollector(logstash_endpoint string) (error, *LogstashCollector) {
	_, nodeStatsCollector := collector.NewNodeStatsCollector(logstash_endpoint)

	return nil, &LogstashCollector{
		collectors: map[string]collector.Collector{
			"node": nodeStatsCollector,
		},
	}
}

func listen(exporter_bind_address string) {
	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
	})

	log.Infoln("Starting server on", exporter_bind_address)
	if err := http.ListenAndServe(exporter_bind_address, nil); err != nil {
		log.Fatalf("Cannot start Logstash exporter: %s", err)
	}
}

func (coll LogstashCollector) Describe(ch chan<- *prometheus.Desc) {
	scrapeDurations.Describe(ch)
}

func (coll LogstashCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(coll.collectors))
	for name, c := range coll.collectors {
		go func(name string, c collector.Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
	scrapeDurations.Collect(ch)
}

func execute(name string, c collector.Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := c.Collect(ch)
	duration := time.Since(begin)
	var result string

	if err != nil {
		log.Errorf("ERROR: %s collector failed after %fs: %s", name, duration.Seconds(), err)
		result = "error"
	} else {
		log.Debugf("OK: %s collector succeeded after %fs.", name, duration.Seconds())
		result = "success"
	}
	scrapeDurations.WithLabelValues(name, result).Observe(duration.Seconds())
}

func init() {
	prometheus.MustRegister(version.NewCollector("logstash_exporter"))
}

func main() {
	var (
		showVersion           = flag.Bool("version", false, "Print version information.")
		logstash_endpoint     = flag.String("logstash.endpoint", "http://localhost:9600", "The protocol, host and po on which logstash metrics API listens")
		exporter_bind_address = flag.String("exporter.bind_address", ":9198", "Exporter bind address")
	)
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("logstash_exporter"))
		os.Exit(0)
	}

	_, logstashCollector := NewLogstashCollector(*logstash_endpoint)
	prometheus.MustRegister(logstashCollector)

	log.Infoln("Starting Logstash exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	listen(*exporter_bind_address)
}
