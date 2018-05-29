package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

// NodeStatsCollector type
type NodeStatsCollector struct {
	endpoint string

	JvmThreadsCount     *prometheus.Desc
	JvmThreadsPeakCount *prometheus.Desc

	MemHeapUsedInBytes         *prometheus.Desc
	MemHeapCommittedInBytes    *prometheus.Desc
	MemHeapMaxInBytes          *prometheus.Desc
	MemNonHeapUsedInBytes      *prometheus.Desc
	MemNonHeapCommittedInBytes *prometheus.Desc

	MemPoolPeakUsedInBytes  *prometheus.Desc
	MemPoolUsedInBytes      *prometheus.Desc
	MemPoolPeakMaxInBytes   *prometheus.Desc
	MemPoolMaxInBytes       *prometheus.Desc
	MemPoolCommittedInBytes *prometheus.Desc

	GCCollectionTimeInMillis *prometheus.Desc
	GCCollectionCount        *prometheus.Desc

	ProcessOpenFileDescriptors    *prometheus.Desc
	ProcessMaxFileDescriptors     *prometheus.Desc
	ProcessMemTotalVirtualInBytes *prometheus.Desc
	ProcessCPUTotalInMillis       *prometheus.Desc

	PipelineDuration       *prometheus.Desc
	PipelineEventsIn       *prometheus.Desc
	PipelineEventsFiltered *prometheus.Desc
	PipelineEventsOut      *prometheus.Desc

	PipelinePluginEventsDuration *prometheus.Desc
	PipelinePluginEventsIn       *prometheus.Desc
	PipelinePluginEventsOut      *prometheus.Desc
	PipelinePluginMatches        *prometheus.Desc
	PipelinePluginFailures       *prometheus.Desc

	PipelineQueueEvents          *prometheus.Desc
	PipelineQueuePageCapacity    *prometheus.Desc
	PipelineQueueMaxQueueSize    *prometheus.Desc
	PipelineQueueMaxUnreadEvents *prometheus.Desc

	PipelineDeadLetterQueueSizeInBytes *prometheus.Desc
}

// NewNodeStatsCollector function
func NewNodeStatsCollector(logstashEndpoint string) (Collector, error) {
	const subsystem = "node"

	return &NodeStatsCollector{
		endpoint: logstashEndpoint,

		JvmThreadsCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "jvm_threads_count"),
			"jvm_threads_count",
			nil,
			nil,
		),

		JvmThreadsPeakCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "jvm_threads_peak_count"),
			"jvm_threads_peak_count",
			nil,
			nil,
		),

		MemHeapUsedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_heap_used_bytes"),
			"mem_heap_used_bytes",
			nil,
			nil,
		),

		MemHeapCommittedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_heap_committed_bytes"),
			"mem_heap_committed_bytes",
			nil,
			nil,
		),

		MemHeapMaxInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_heap_max_bytes"),
			"mem_heap_max_bytes",
			nil,
			nil,
		),

		MemNonHeapUsedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_nonheap_used_bytes"),
			"mem_nonheap_used_bytes",
			nil,
			nil,
		),

		MemNonHeapCommittedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_nonheap_committed_bytes"),
			"mem_nonheap_committed_bytes",
			nil,
			nil,
		),

		MemPoolUsedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_pool_used_bytes"),
			"mem_pool_used_bytes",
			[]string{"pool"},
			nil,
		),

		MemPoolPeakUsedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_pool_peak_used_bytes"),
			"mem_pool_peak_used_bytes",
			[]string{"pool"},
			nil,
		),

		MemPoolPeakMaxInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_pool_peak_max_bytes"),
			"mem_pool_peak_max_bytes",
			[]string{"pool"},
			nil,
		),

		MemPoolMaxInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_pool_max_bytes"),
			"mem_pool_max_bytes",
			[]string{"pool"},
			nil,
		),

		MemPoolCommittedInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "mem_pool_committed_bytes"),
			"mem_pool_committed_bytes",
			[]string{"pool"},
			nil,
		),

		GCCollectionTimeInMillis: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "gc_collection_duration_seconds_total"),
			"gc_collection_duration_seconds_total",
			[]string{"collector"},
			nil,
		),

		GCCollectionCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "gc_collection_total"),
			"gc_collection_total",
			[]string{"collector"},
			nil,
		),

		ProcessOpenFileDescriptors: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "process_open_filedescriptors"),
			"process_open_filedescriptors",
			nil,
			nil,
		),

		ProcessMaxFileDescriptors: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "process_max_filedescriptors"),
			"process_max_filedescriptors",
			nil,
			nil,
		),

		ProcessMemTotalVirtualInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "process_mem_total_virtual_bytes"),
			"process_mem_total_virtual_bytes",
			nil,
			nil,
		),

		ProcessCPUTotalInMillis: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "process_cpu_total_seconds_total"),
			"process_cpu_total_seconds_total",
			nil,
			nil,
		),

		PipelineDuration: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "pipeline_duration_seconds_total"),
			"pipeline_duration_seconds_total",
			[]string{"pipeline"},
			nil,
		),

		PipelineEventsIn: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "pipeline_events_in_total"),
			"pipeline_events_in_total",
			[]string{"pipeline"},
			nil,
		),

		PipelineEventsFiltered: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "pipeline_events_filtered_total"),
			"pipeline_events_filtered_total",
			[]string{"pipeline"},
			nil,
		),

		PipelineEventsOut: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "pipeline_events_out_total"),
			"pipeline_events_out_total",
			[]string{"pipeline"},
			nil,
		),

		PipelinePluginEventsDuration: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "plugin_duration_seconds_total"),
			"plugin_duration_seconds",
			[]string{"pipeline", "plugin", "plugin_id", "plugin_type"},
			nil,
		),

		PipelinePluginEventsIn: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "plugin_events_in_total"),
			"plugin_events_in",
			[]string{"pipeline", "plugin", "plugin_id", "plugin_type"},
			nil,
		),

		PipelinePluginEventsOut: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "plugin_events_out_total"),
			"plugin_events_out",
			[]string{"pipeline", "plugin", "plugin_id", "plugin_type"},
			nil,
		),

		PipelinePluginMatches: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "plugin_matches_total"),
			"plugin_matches",
			[]string{"pipeline", "plugin", "plugin_id", "plugin_type"},
			nil,
		),

		PipelinePluginFailures: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "plugin_failures_total"),
			"plugin_failures",
			[]string{"pipeline", "plugin", "plugin_id", "plugin_type"},
			nil,
		),

		PipelineQueueEvents: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "queue_events"),
			"queue_events",
			[]string{"pipeline"},
			nil,
		),

		PipelineQueuePageCapacity: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "queue_page_capacity_bytes"),
			"queue_page_capacity_bytes",
			[]string{"pipeline"},
			nil,
		),

		PipelineQueueMaxQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "queue_max_size_bytes"),
			"queue_max_size_bytes",
			[]string{"pipeline"},
			nil,
		),

		PipelineQueueMaxUnreadEvents: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "queue_max_unread_events"),
			"queue_max_unread_events",
			[]string{"pipeline"},
			nil,
		),

		PipelineDeadLetterQueueSizeInBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "dead_letter_queue_size_bytes"),
			"dead_letter_queue_size_bytes",
			[]string{"pipeline"},
			nil,
		),
	}, nil
}

// Collect function implements nodestats_collector collector
func (c *NodeStatsCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("Failed collecting node metrics", desc, err)
		return err
	}
	return nil
}

func (c *NodeStatsCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	stats, err := NodeStats(c.endpoint)
	if err != nil {
		return nil, err
	}

	ch <- prometheus.MustNewConstMetric(
		c.JvmThreadsCount,
		prometheus.GaugeValue,
		float64(stats.Jvm.Threads.Count),
	)

	ch <- prometheus.MustNewConstMetric(
		c.JvmThreadsPeakCount,
		prometheus.GaugeValue,
		float64(stats.Jvm.Threads.PeakCount),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemHeapUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.HeapUsedInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemHeapMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.HeapMaxInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemHeapCommittedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.HeapCommittedInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemNonHeapUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.NonHeapUsedInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemNonHeapCommittedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.NonHeapCommittedInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakUsedInBytes),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.UsedInBytes),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakMaxInBytes),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.MaxInBytes),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolCommittedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.CommittedInBytes),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakUsedInBytes),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Young.UsedInBytes),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakMaxInBytes),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Young.MaxInBytes),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolCommittedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Young.CommittedInBytes),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakUsedInBytes),
		"survivor",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolUsedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Survivor.UsedInBytes),
		"survivor",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolPeakMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Old.PeakMaxInBytes),
		"survivor",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolMaxInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Survivor.MaxInBytes),
		"survivor",
	)

	ch <- prometheus.MustNewConstMetric(
		c.MemPoolCommittedInBytes,
		prometheus.GaugeValue,
		float64(stats.Jvm.Mem.Pools.Survivor.CommittedInBytes),
		"survivor",
	)

	ch <- prometheus.MustNewConstMetric(
		c.GCCollectionTimeInMillis,
		prometheus.CounterValue,
		0.001 * float64(stats.Jvm.Gc.Collectors.Old.CollectionTimeInMillis),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.GCCollectionCount,
		prometheus.GaugeValue,
		float64(stats.Jvm.Gc.Collectors.Old.CollectionCount),
		"old",
	)

	ch <- prometheus.MustNewConstMetric(
		c.GCCollectionTimeInMillis,
		prometheus.CounterValue,
		0.001 * float64(stats.Jvm.Gc.Collectors.Young.CollectionTimeInMillis),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.GCCollectionCount,
		prometheus.GaugeValue,
		float64(stats.Jvm.Gc.Collectors.Young.CollectionCount),
		"young",
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessOpenFileDescriptors,
		prometheus.GaugeValue,
		float64(stats.Process.OpenFileDescriptors),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessMaxFileDescriptors,
		prometheus.GaugeValue,
		float64(stats.Process.MaxFileDescriptors),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessMemTotalVirtualInBytes,
		prometheus.GaugeValue,
		float64(stats.Process.Mem.TotalVirtualInBytes),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessCPUTotalInMillis,
		prometheus.CounterValue,
		float64(stats.Process.CPU.TotalInMillis/1000),
	)

	// For backwards compatibility with Logstash 5
	pipelines := make(map[string]Pipeline)
	if len(stats.Pipelines) == 0 {
		pipelines["main"] = stats.Pipeline
	} else {
		pipelines = stats.Pipelines
	}

	for pipelineID, pipeline := range pipelines {
		ch <- prometheus.MustNewConstMetric(
			c.PipelineDuration,
			prometheus.CounterValue,
			float64(pipeline.Events.DurationInMillis/1000),
			pipelineID,
		)

		ch <- prometheus.MustNewConstMetric(
			c.PipelineEventsIn,
			prometheus.CounterValue,
			float64(pipeline.Events.In),
			pipelineID,
		)

		ch <- prometheus.MustNewConstMetric(
			c.PipelineEventsFiltered,
			prometheus.CounterValue,
			float64(pipeline.Events.Filtered),
			pipelineID,
		)

		ch <- prometheus.MustNewConstMetric(
			c.PipelineEventsOut,
			prometheus.CounterValue,
			float64(pipeline.Events.Out),
			pipelineID,
		)

		for _, plugin := range pipeline.Plugins.Inputs {
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsIn,
				prometheus.CounterValue,
				float64(plugin.Events.In),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"input",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsOut,
				prometheus.CounterValue,
				float64(plugin.Events.Out),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"input",
			)
		}

		for _, plugin := range pipeline.Plugins.Filters {
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsDuration,
				prometheus.CounterValue,
				float64(plugin.Events.DurationInMillis/1000),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"filter",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsIn,
				prometheus.CounterValue,
				float64(plugin.Events.In),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"filter",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsOut,
				prometheus.CounterValue,
				float64(plugin.Events.Out),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"filter",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginMatches,
				prometheus.CounterValue,
				float64(plugin.Matches),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"filter",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginFailures,
				prometheus.CounterValue,
				float64(plugin.Failures),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"filter",
			)
		}

		for _, plugin := range pipeline.Plugins.Outputs {
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsIn,
				prometheus.CounterValue,
				float64(plugin.Events.In),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"output",
			)
			ch <- prometheus.MustNewConstMetric(
				c.PipelinePluginEventsOut,
				prometheus.CounterValue,
				float64(plugin.Events.Out),
				pipelineID,
				plugin.Name,
				plugin.ID,
				"output",
			)
		}

		if pipeline.Queue.Type != "memory" {
			ch <- prometheus.MustNewConstMetric(
				c.PipelineQueueEvents,
				prometheus.CounterValue,
				float64(pipeline.Queue.Events),
				pipelineID,
			)

			ch <- prometheus.MustNewConstMetric(
				c.PipelineQueuePageCapacity,
				prometheus.CounterValue,
				float64(pipeline.Queue.Capacity.PageCapacityInBytes),
				pipelineID,
			)

			ch <- prometheus.MustNewConstMetric(
				c.PipelineQueueMaxQueueSize,
				prometheus.CounterValue,
				float64(pipeline.Queue.Capacity.MaxQueueSizeInBytes),
				pipelineID,
			)

			ch <- prometheus.MustNewConstMetric(
				c.PipelineQueueMaxUnreadEvents,
				prometheus.CounterValue,
				float64(pipeline.Queue.Capacity.MaxUnreadEvents),
				pipelineID,
			)
		}

		if pipeline.DeadLetterQueue.QueueSizeInBytes != 0 {
			ch <- prometheus.MustNewConstMetric(
				c.PipelineDeadLetterQueueSizeInBytes,
				prometheus.GaugeValue,
				float64(pipeline.DeadLetterQueue.QueueSizeInBytes),
				pipelineID,
			)
		}
	}

	return nil, nil
}
