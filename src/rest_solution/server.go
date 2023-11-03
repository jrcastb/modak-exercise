package main

import (
	"modak-exercise/src/rest_solution/api"
	"net/http"
)

func NewServer(addr string) *http.Server {
	api.InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
