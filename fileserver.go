package auth

import "net/http"

type FileServer struct {
	handler http.Handler
}

func (f *FileServer) Initialize() http.Handler {
	f.handler = http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
	return f.handler
}
