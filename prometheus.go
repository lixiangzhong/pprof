package pprof

import "github.com/prometheus/client_golang/prometheus/promhttp"

func Prometheus(register RouterRegister) {
	register.Handle("/metrics", promhttp.Handler())
}
