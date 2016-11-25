package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "logstash"
)

type Collector interface {
	Collect(ch chan<- prometheus.Metric) (err error)
}
