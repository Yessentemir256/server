package app

import (
	"net/http"

	"github.com/Yessentemir256/server/pkg/banners"
)

// Server представляет собой логический сервер нашего приложения.
type Server struct {
	mux        *http.ServeMux
	bannersSvc *banners.Service
}

// NewServer - функция-конструктор для создания сервера.
func NewServer(mux *http.ServeMux, bannersSvc *banners.Service) *Server {
	return &Server{mux: mux}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

// Init инициализирует сервер (регистрирует все Handler'ы)
func (s *Server) Init() {
	s.mux.HandleFunc("/banners.getAll", s.handleGetAllBanners)
	s.mux.HandleFunc("/banners.getById", s.handleGetBannerByID)
	s.mux.HandleFunc("/banners.save", s.handleSaveBanner)
	s.mux.HandleFunc("/banners.removeById", s.handleRemoveByID)
}

func (s *Server) handleGetAllBanners(writer http.ResponseWriter, request *http.Request) {
	// Логика обработки запроса "banners.getAll"
}

func (s *Server) handleGetBannerByID(writer http.ResponseWriter, request *http.Request) {
	// Логика обработки запроса "banners.getById"
}

func (s *Server) handleSaveBanner(writer http.ResponseWriter, request *http.Request) {
	// Логика обработки запроса "banners.save"
}

func (s *Server) handleRemoveByID(writer http.ResponseWriter, request *http.Request) {
	// Логика обработки запроса "banners.removeById"
}
