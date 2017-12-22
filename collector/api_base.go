package collector

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
)

// HTTPHandler type
type HTTPHandler struct {
	Endpoint string
}

// Get method for HTTPHandler
func (h *HTTPHandler) Get() (http.Response, error) {
	response, err := http.Get(h.Endpoint)
	if err != nil {
		return http.Response{}, err
	}

	return *response, nil
}

// HTTPHandlerInterface interface
type HTTPHandlerInterface interface {
	Get() (http.Response, error)
}

func getMetrics(h HTTPHandlerInterface, target interface{}) error {
	response, err := h.Get()
	if err != nil {
		log.Errorf("Cannot retrieve metrics: %s", err)
		return nil
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Errorf("Cannot close response body: %v", err)
		}
	}()

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		log.Errorf("Cannot parse Logstash response json: %s", err)
	}

	return nil
}
