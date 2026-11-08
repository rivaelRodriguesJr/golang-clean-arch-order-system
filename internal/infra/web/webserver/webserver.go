package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebServer struct {
	Router        chi.Router
	Methods       map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Methods:       make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(method, pattern string, handler http.HandlerFunc) {
	s.Methods[method+"|"+pattern] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for key, handler := range s.Methods {
		splited := strings.Split(key, "|")
		method := splited[0]
		path := splited[1]

		s.Router.MethodFunc(method, path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
