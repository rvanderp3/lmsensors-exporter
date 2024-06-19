package pkg

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

var (
	guages map[string]*prometheus.Gauge
)

func reportGuageMetric(name string, value int64) {
	var gauge prometheus.Gauge
	name = strings.ReplaceAll(name, "-", "_")
	if _gauge, ok := guages[name]; !ok {
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
		})
		guages[name] = &gauge
		prometheus.MustRegister(gauge)
	} else {
		gauge = *_gauge
	}
	gauge.Set(float64(value))
}

func init() {
	guages = make(map[string]*prometheus.Gauge)
	// Serve the default Prometheus metrics at /metrics
	http.HandleFunc("/metrics", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	collect()
	promHandler := promhttp.Handler()
	promHandler.ServeHTTP(w, r)
}
