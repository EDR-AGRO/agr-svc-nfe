package api

import (
	"br/com/agr/nfe/api/v1/handlers"
	"br/com/agr/nfe/domain/services"
	"br/com/agr/nfe/infrastructure/httpclient"
	"br/com/agr/nfe/resources/configs"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ApiServer struct {
	services *ApiSevices
	handlers *ApiHandlers
}

type ApiSevices struct {
	NfeServiceIF *services.NfeServiceIF
}

type ApiHandlers struct {
	NfeRouterHandler *handlers.NfeRouterHandler
}

func NewServer(services *ApiSevices) *ApiServer {
	apiServer := &ApiServer{
		services: services,
	}

	apiServer.SetHandlers()

	return apiServer
}

func (api *ApiServer) SetHandlers() {
	api.handlers = &ApiHandlers{
		NfeRouterHandler: handlers.NewNfeRouterHandler(*api.services.NfeServiceIF),
	}
}

func (api *ApiServer) SetupRoutes(envBaseUrl string, chiRouter *chi.Mux) {
	api.setupCORS(chiRouter)

	envBaseUrl = fmt.Sprintf("/%s", envBaseUrl)

	api.setHeartBeat(envBaseUrl, chiRouter)
	api.registerRoutes(envBaseUrl, chiRouter)
	api.registerCommonAPI(envBaseUrl, chiRouter)
}

func (api *ApiServer) registerCommonAPI(envBaseUrl string, subrouter chi.Router) {
	subrouter.Group(func(r chi.Router) {
		if configs.Cfg.ShowSwagger == 1 {
			r.Mount(envBaseUrl+"/swagger", httpSwagger.WrapHandler)
		}
	})
}

func (api *ApiServer) registerRoutes(envBaseUrl string, subrouter chi.Router) {
	subrouter.Group(func(r chi.Router) {
		r.Use(httpclient.Logger)
		r.Post(envBaseUrl+"/v1/nfe", api.handlers.NfeRouterHandler.EmitNfe)
	})
}

func (api *ApiServer) setupCORS(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func (api *ApiServer) setHeartBeat(envBaseUrl string, r *chi.Mux) {
	r.Use(middleware.Heartbeat(envBaseUrl + "/health/liveness"))
	r.Use(middleware.Heartbeat(envBaseUrl + "/health/readiness"))
}
