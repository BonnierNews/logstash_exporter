package collector

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var noQueueJSON = []byte(`
{
  "host": "69e7fa935209",
  "version": "5.6.4",
  "http_address": "0.0.0.0:9600",
  "id": "4607a9c2-8517-4e95-a7ce-742e39264a95",
  "name": "69e7fa935209",
  "jvm": {
    "threads": {
      "count": 20,
      "peak_count": 20
    },
    "mem": {
      "heap_used_percent": 20,
      "heap_committed_in_bytes": 275677184,
      "heap_max_in_bytes": 1056309248,
      "heap_used_in_bytes": 215538696,
      "non_heap_used_in_bytes": 75598880,
      "non_heap_committed_in_bytes": 79458304,
      "pools": {
        "survivor": {
          "peak_used_in_bytes": 17432576,
          "used_in_bytes": 14406616,
          "peak_max_in_bytes": 17432576,
          "max_in_bytes": 17432576,
          "committed_in_bytes": 17432576
        },
        "old": {
          "peak_used_in_bytes": 88768424,
          "used_in_bytes": 88768424,
          "peak_max_in_bytes": 899284992,
          "max_in_bytes": 899284992,
          "committed_in_bytes": 118652928
        },
        "young": {
          "peak_used_in_bytes": 139591680,
          "used_in_bytes": 112363656,
          "peak_max_in_bytes": 139591680,
          "max_in_bytes": 139591680,
          "committed_in_bytes": 139591680
        }
      }
    },
    "gc": {
      "collectors": {
        "old": {
          "collection_time_in_millis": 108,
          "collection_count": 2
        },
        "young": {
          "collection_time_in_millis": 630,
          "collection_count": 7
        }
      }
    },
    "uptime_in_millis": 24099
  },
  "process": {
    "open_file_descriptors": 63,
    "peak_open_file_descriptors": 63,
    "max_file_descriptors": 1048576,
    "mem": {
      "total_virtual_in_bytes": 3948072960
    },
    "cpu": {
      "total_in_millis": 37720,
      "percent": 21,
      "load_average": {
        "1m": 0.94,
        "5m": 0.22,
        "15m": 0.08
      }
    }
  },
  "pipeline": {
    "events": {
      "duration_in_millis": 0,
      "in": 0,
      "filtered": 0,
      "out": 0,
      "queue_push_duration_in_millis": 0
    },
    "plugins": {
      "inputs": [
        {
          "id": "5681ae93b83a24a100eacdb291ca4679effa35bf-1",
          "events": {
            "out": 0,
            "queue_push_duration_in_millis": 0
          },
          "name": "stdin"
        }
      ],
      "filters": [],
      "outputs": [
        {
          "id": "5681ae93b83a24a100eacdb291ca4679effa35bf-2",
          "events": {
            "duration_in_millis": 0,
            "in": 0,
            "out": 0
          },
          "name": "stdout"
        }
      ]
    },
    "reloads": {
      "last_error": null,
      "successes": 0,
      "last_success_timestamp": null,
      "last_failure_timestamp": null,
      "failures": 0
    },
    "queue": {
      "type": "memory"
    },
    "id": "main"
  },
  "reloads": {
    "successes": 0,
    "failures": 0
  },
  "os": {}
}
`)

var queueJSON = []byte(`
{
  "pipeline" : {
    "events" : {
      "duration_in_millis" : 1955,
      "in" : 100,
      "filtered" : 100,
      "out" : 100,
      "queue_push_duration_in_millis" : 71
    },
    "plugins" : {
      "inputs" : [ {
        "id" : "729b0efdc657715a4a59103ab2643c010fc46e77-1",
        "events" : {
          "out" : 100,
          "queue_push_duration_in_millis" : 71
        },
        "name" : "beats"
      } ],
      "filters" : [ {
        "id" : "729b0efdc657715a4a59103ab2643c010fc46e77-2",
        "events" : {
          "duration_in_millis" : 64,
          "in" : 100,
          "out" : 100
        },
        "matches" : 100,
        "patterns_per_field" : {
          "message" : 1
        },
        "name" : "grok"
      } ],
      "outputs" : [ {
        "id" : "729b0efdc657715a4a59103ab2643c010fc46e77-3",
        "events" : {
          "duration_in_millis" : 1724,
          "in" : 100,
          "out" : 100
        },
        "name" : "stdout"
      } ]
    },
    "reloads" : {
      "last_error" : null,
      "successes" : 2,
      "last_success_timestamp" : "2017-05-25T02:40:40.974Z",
      "last_failure_timestamp" : null,
      "failures" : 0
    },
    "queue" : {
      "events" : 0,
      "type" : "persisted",
      "capacity" : {
        "page_capacity_in_bytes" : 262144000,
        "max_queue_size_in_bytes" : 8589934592,
        "max_unread_events" : 12
      },
      "data" : {
        "path" : "/path/to/data/queue",
        "free_space_in_bytes" : 89280552960,
        "storage_type" : "hfs"
      }
    },
    "id" : "main"
  }
}
`)

type MockHTTPHandler struct {
	ReturnJSON []byte
	Endpoint   string
}

func (m *MockHTTPHandler) Get() (http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(m.ReturnJSON)),
	}

	return *response, nil
}

func TestPipelineNoQueueStats(t *testing.T) {
	var response NodeStatsResponse

	m := &MockHTTPHandler{ReturnJSON: noQueueJSON}
	getNodeStats(m, &response)

	if response.Pipeline.Queue.Capacity.MaxUnreadEvents == 12 {
		t.Fail()
	}
}

func TestPipelineQueueStats(t *testing.T) {
	var response NodeStatsResponse

	m := &MockHTTPHandler{ReturnJSON: queueJSON}
	getNodeStats(m, &response)

	if response.Pipeline.Queue.Capacity.MaxUnreadEvents != 12 {
		t.Fail()
	}
}
