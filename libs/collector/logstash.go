package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	// Namespace const string
	Namespace = "logstash"
)

// Collector interface implement Collect function
type Collector interface {
	Collect(ch chan<- prometheus.Metric) (err error)
}
