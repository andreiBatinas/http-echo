package http

import (
	"fmt"
	_http "net/http"

	"alluvial/interview/config"
	"alluvial/interview/infrastructure/log"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerOptions struct {
	Port int
	Log  *log.Logger
}

type Server struct {
	apiGroup *echo.Group
	config   *ServerOptions
	handler  *echo.Echo
	log      *log.Logger
}

// RouterInterface router implementation
type RouterInterface interface {
	Handler()
}

func (http *Server) SetupDefaultEndpoints() {
	http.handler.GET("/metrics", echoprometheus.NewHandler())

	http.handler.GET("/health", func(c echo.Context) error {
		return c.String(_http.StatusOK, "ok")
	})
}

// SubscribeRouter subscribe a router group to HTTP Handler
func (http *Server) SubscribeRouter(ri RouterInterface) {
	ri.Handler()
}

// API add routes with prefix "api"
func (http *Server) API(r *Router) *echo.Group {
	httpGroup := http.apiGroup.Group(r.Name)
	return httpGroup
}

func (http *Server) Start() {
	if config.Configuration.Debug == true {
		http.handler.Use(middleware.Logger())
	}

	http.handler.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: middleware.DefaultCORSConfig.AllowOrigins}))
	http.handler.Use(echoprometheus.NewMiddleware("app"))

	http.SetupDefaultEndpoints()

	http.handler.Logger.Fatal(http.handler.Start(fmt.Sprintf(":%d", http.config.Port)))

}

func NewHttpServer(opts ServerOptions) *Server {
	e := echo.New()

	srv := &Server{
		apiGroup: e.Group("/eth"),
		config:   &opts,
		handler:  e,
		log:      opts.Log,
	}
	return srv
}
