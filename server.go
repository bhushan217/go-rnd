package main

import(
	"net/http"
)

type Server struct{
	*dbStore
}

func NewServer(store *dbStore) *Server{
	return &Server{store}
}
