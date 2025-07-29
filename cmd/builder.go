package cmd

import (
	"br/com/agr/nfe/api"
	"br/com/agr/nfe/domain/services"
	"br/com/agr/nfe/infrastructure/apm"
	"br/com/agr/nfe/infrastructure/logger"
	"br/com/agr/nfe/resources/configs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.elastic.co/apm/module/apmhttp"
)

const baseURL = "agr-svc-nfe"

type Builder struct {
	mongoIF      interface{}
	postgresIF   interface{}
	nfeServiceIF services.NfeServiceIF
}

func NewBuilder() *Builder {
	return &Builder{}
}

type AppBuilder interface {
	setConfig(apmt *apm.ApmTransaction) error
	setDatabases(apmt *apm.ApmTransaction) error
	setServices(apmt *apm.ApmTransaction) error
	setWebServer(apmt *apm.ApmTransaction) chan error
}

type BuildDirector struct {
	builder AppBuilder
}

func NewAppBuilder(sb AppBuilder) *BuildDirector {
	return &BuildDirector{
		builder: sb,
	}
}

func (sbd *BuildDirector) Build(apmt *apm.ApmTransaction) (chan error, error) {

	logger.Infof(apmt.Ctx, "Staring Build...")

	var err error = nil

	defer func() {
		if err != nil {
			logger.Errorf(apmt.Ctx, "Build failed: %v", err.Error())
		}
	}()

	//mantenha a ordem de chamada
	//uma depende da outra

	if err := sbd.builder.setConfig(apmt); err != nil {
		return nil, err
	}

	if err := sbd.builder.setDatabases(apmt); err != nil {
		return nil, err
	}

	if err := sbd.builder.setServices(apmt); err != nil {
		return nil, err
	}

	errChan := sbd.builder.setWebServer(apmt)

	logger.Infof(apmt.Ctx, "Build completed successfully")

	return errChan, nil
}

func (e *Builder) setConfig(apmt *apm.ApmTransaction) error {
	logger.Infof(apmt.Ctx, "Loading configuration...")

	_, err := configs.LoadConfig(apmt)
	if err != nil {
		logger.Errorf(apmt.Ctx, "Error loading configuration: %v", err.Error())
		return err
	}

	logger.Infof(apmt.Ctx, "Configuration loaded successfully")
	return nil
}

func (e *Builder) setWebServer(apmt *apm.ApmTransaction) chan error {
	logger.Infof(apmt.Ctx, "Starting Web Server...")

	errChan := make(chan error, 1)

	apiServices := &api.ApiSevices{
		NfeServiceIF: &e.nfeServiceIF,
	}

	webServer := api.NewServer(apiServices)

	chiRouter := chi.NewRouter()
	webServer.SetupRoutes(baseURL, chiRouter)

	go func() {
		if err := http.ListenAndServe(configs.Cfg.Port, apmhttp.Wrap(chiRouter)); err != nil {
			logger.Errorf(apmt.Ctx, "Error starting web server: %v", err.Error())
			errChan <- err
		}
	}()
	logger.Infof(apmt.Ctx, "NFE - API Running on port %s", configs.Cfg.Port)

	return errChan
}

func (e *Builder) setServices(apmt *apm.ApmTransaction) error {
	e.nfeServiceIF = services.NewNfeService()

	logger.Infof(apmt.Ctx, "Services set up successfully")
	return nil
}

func (e *Builder) setDatabases(apmt *apm.ApmTransaction) error {
	// logger.Infof(apmt.Ctx, "Connecting to databases...")
	// conn, err := mongo.InitMongoConnection(apmt)
	// if err != nil {
	// 	logger.Errorf(apmt.Ctx, "Error connecting to MongoDB database: %v", err.Error())
	// 	return err
	// }

	// e.mongoIF = mongo.NewMongoDB(conn)

	// logger.Infof(apmt.Ctx, "Connected to MongoDB database successfully")

	// conn, err := postgres.InitPostgresConnection(apmt)
	// if err != nil {
	// 	logger.Errorf(apmt.Ctx, "Error connecting to PostgresDB database: %v", err.Error())
	// 	return err
	// }

	// e.postgresDBIF = postgres.NewPostgresDB(conn)

	// logger.Infof(apmt.Ctx, "Connected to postgresDB database successfully")
	return nil
}
