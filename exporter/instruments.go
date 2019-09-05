package exporter

import (
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var registry *prometheus.Registry

func init() {
	registry = prometheus.NewRegistry()
	registry.MustRegister(dockerContainerState)
}

var dockerContainerState = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "docker_container_state",
		Help: "Current state of Docker container. 1->running; 2->created, 3->paused, 4->restarting, 5->removing, 6->exited, 7->dead",
	},
	[]string{"instance", "name"},
)

func DockerContainerMetrics(ch <-chan Update) {
	go func() {
		for {
			if ch == nil {
				break
			}

			update := <-ch
			containerName := strings.TrimLeft(update.Name, "/")
			dockerContainerState.WithLabelValues(update.Instance, containerName).Set(float64(StateValue(update.State)))
		}
	}()
}

func HandleHTTP(addr string) error {
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)
	return http.ListenAndServe(addr, nil)
}
