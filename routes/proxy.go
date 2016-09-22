package routes

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/tiborv/auth-proxy/models"
)

const (
	proxyPath  = "/api/proxy/"
	headerName = "Authorization"
)

func init() {
	mux.HandleFunc(proxyPath, proxy)
}

func StripPrefix(req *http.Request) {
	req.URL.Path = req.URL.Path[len(proxyPath):]
}

func GetPathSlugToken(req *http.Request) (string, string, string) {
	tokenSplit := strings.Split(req.Header.Get(headerName), " ")
	urlSplit := strings.Split(req.URL.Path, "/")
	return urlSplit[0], "/" + strings.Join(urlSplit[1:], "/"), tokenSplit[1]
}

func proxy(w http.ResponseWriter, r *http.Request) {
	StripPrefix(r)
	slug, path, token := GetPathSlugToken(r)
	service, serviceNotFound := models.FindServiceBySlug(slug)
	authorized, tokenNotFound := models.ServiceHasToken(service.Slug, token)
	if !authorized || tokenNotFound != nil || serviceNotFound != nil {
		HttpResponse{Status: http.StatusForbidden, Msg: "Not authorized"}.Send(w)
		return
	}

	reqTime := time.Now()
	models.LogRequest(token, slug, reqTime, r.Header)

	r.Host = service.Host
	director := func(req *http.Request) {
		req = r
		r.URL.Scheme = service.Scheme
		req.URL.Host = service.Host
		req.URL.Path = path

	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(w, r)

	models.LogResponse(token, slug, time.Now().Sub(reqTime), w)
}
