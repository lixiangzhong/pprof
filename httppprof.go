package pprof

import (
	"net/http"
	"net/http/pprof"
	"path"

	"github.com/gin-gonic/gin"
)

//RouterRegister http.ServerMux
type RouterRegister interface {
	Handle(string, http.Handler)
}

func RegisterHTTPHandler(server RouterRegister, prefix ...string) {
	var root string
	if len(prefix) > 0 {
		root = prefix[0]
		root = "/" + root
		root = path.Clean(root)
	}
	server.Handle(root+"/debug/pprof/", http.HandlerFunc(pprof.Index))
	server.Handle(root+"/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	server.Handle(root+"/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	server.Handle(root+"/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	server.Handle(root+"/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
}

func RegisterGinHandler(router gin.IRouter) {
	router.GET("/debug/pprof/", gin.WrapH(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
	router.GET("/debug/pprof/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
	router.GET("/debug/pprof/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
	router.GET("/debug/pprof/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
}
