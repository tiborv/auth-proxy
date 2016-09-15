package routes

import (
	"net/http"
	"net/http/httputil"
)

const proxyPath = "/api/proxy"

func init() {
	mux.HandleFunc(proxyPath, proxy)
}

func proxy(w http.ResponseWriter, r *http.Request) {

	director := func(req *http.Request) {
		req = r
		req.URL.Scheme = "http"
		req.URL.Host = "content.kk.no"
		req.URL.Path = "/"
		req.Header.Add("Host", req.URL.Host)

	}
	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ServeHTTP(w, r)
}
