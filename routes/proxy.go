package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/tiborv/prxy/db"
)

const (
	proxyPath  = "/api/proxy/"
	headerName = "Auth-Token"
)

func init() {
	mux.HandleFunc(proxyPath, proxy)
}

func StripPrefix(req *http.Request) {
	req.URL.Path = req.URL.Path[len(proxyPath):]
}

func GetPathSlugToken(req *http.Request) (string, string, string) {
	urlSplit := strings.Split(req.URL.Path, "/")
	return urlSplit[0], "/" + strings.Join(urlSplit[1:], "/"), req.Header.Get(headerName)
}

func proxy(w http.ResponseWriter, r *http.Request) {
	StripPrefix(r)
	slug, path, token := GetPathSlugToken(r)
	service, serviceNotFound := db.FindServiceBySlug(slug)
	authorized, tokenNotFound := db.ServiceHasToken(service.Slug, token)
	if !authorized || tokenNotFound != nil || serviceNotFound != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Not authorized")
		return
	}

	r.Host = service.Host
	director := func(req *http.Request) {
		req = r
		r.URL.Scheme = service.Scheme
		req.URL.Host = service.Host
		req.URL.Path = path

	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)
}
