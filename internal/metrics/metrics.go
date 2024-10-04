package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	SuccessfulRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_successfmul_requests_count",
			Help: "Total number of successful requests",
		},
	)
	UnsuccessfulRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_unsuccessful_requests_count",
			Help: "Total number of unsuccessful requests",
		},
	)
)

func Init() {
	prometheus.MustRegister(SuccessfulRequests)
	prometheus.MustRegister(UnsuccessfulRequests)
}
