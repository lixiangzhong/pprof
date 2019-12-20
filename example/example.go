package main

import (
	"net/http"

	"github.com/lixiangzhong/pprof"
)

func main() {
	m := http.NewServeMux()
	pprof.RegisterHTTPHandler(m)
	http.ListenAndServe(":9999", m)
}
