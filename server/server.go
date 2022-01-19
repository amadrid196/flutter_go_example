package server

import "net/http"

type Country struct {
	Name     string
	Lenguaje string
}

var countrie []*Country = []*Country{}

func New(addr string) http.Server {
	initRoutes()
	return http.Server{
		Addr: addr,
	}
}
