package router

import (
	"backend/internal/services/buildcv"
	"backend/req/http/build"
	"backend/req/http/success"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/build", build.Handler)
	mux.HandleFunc("/api/success", success.Handler)
	mux.HandleFunc("/api/buildcv", buildcv.Handler)
}
