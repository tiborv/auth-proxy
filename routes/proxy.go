package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const proxyPath = "/api/proxy"

func init() {
	mux.HandleFunc(proxyPath, proxy)
}

func proxy(w http.ResponseWriter, r *http.Request) {
	r.Host = "content.kk.no"

	director := func(req *http.Request) {
		req = r
		req.URL.Scheme = "http"
		req.URL.Host = "content.kk.no"
		req.URL.Path = "/tags/helse"
		fmt.Println(req.Host)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
