package server

import (
	"IoT_Service/server/internal"
	"IoT_Service/server/services/gRPC"
	"IoT_Service/server/services/restApi"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// App structures
type App struct {
	APIServer  *restApi.Server
	GRPCServer *gRPC.Server
	Error      error
}

func NewApp() *App {
	return &App{}
}

// Initialize load the configFile.yml file and create services
func (app *App) Initialize() error {
	// Create services
	app.APIServer = restApi.New()
	app.GRPCServer = gRPC.NewServer()

	// Load internal file
	cfg := &internal.Services{}
	f, err := os.Open("configFile.yml")
	defer f.Close()

	// Parse internal file to the structures and check for errors
	if err != nil {
		return internal.WrapError("Error on Application Init(): error on 'openFile('configFile.yml')'", err)
	}
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return internal.WrapError("Error on Application Init(): Error on 'decode yml'", err)
	}
	err = cfg.GRPC.Check()
	if err != nil {
		return internal.WrapError("Error on Application Init(): Error on 'configFile.yml section gRPC'", err)
	}
	err = cfg.API.Check()
	if err != nil {
		return internal.WrapError("Error on Application Init(): Error on 'configFile.yml section api'", err)
	}
	err = cfg.DB.Check()
	if err != nil {
		return internal.WrapError("Error on Application Init(): Error on 'configFile.yml section db'", err)
	}

	// insert config foreach services
	app.APIServer.Config = cfg.API
	app.GRPCServer.ConfigGrpc = cfg.GRPC
	app.GRPCServer.ConfigDB = cfg.DB
	// Initialize services
	if err := app.GRPCServer.Initialize(); err != nil {
		return internal.WrapError("Error on Initialize 'GRPC-Server'", err)
	}
	if err := app.APIServer.Initialize(); err != nil {
		return internal.WrapError("Error on Initialize 'API-Server'", err)
	}

	return nil
}

// Start starts all created services
func (app *App) Start() error {
	gRPCError := make(chan error)
	apiError := make(chan error)

	//
	// GRPC RUN
	//
	go func() {
		defer close(gRPCError)
		err := app.GRPCServer.Run()
		if err != nil {
			gRPCError <- err
			return
		}
	}()
	//
	// API RUN
	//
	go func() {
		defer close(apiError)
		err := app.APIServer.Run()
		if err != nil {
			apiError <- err
			return
		}
	}()

	app.Error = internal.WrapError("Error on 'grpcServer' ", <-gRPCError)
	if app.Error != nil {
		return app.Error
	}
	app.Error = internal.WrapError("Error on 'apiServer' ", <-gRPCError)
	if app.Error != nil {
		return app.Error
	}

	return nil
}

// Stop stops all created services
func (app *App) Stop() {
	log.Print("Stop services...")
	if err := app.GRPCServer.Stop(); err != nil {
		log.Print(internal.WrapError("Error on 'gRPC-Server-Stop'", err))
	}
	if err := app.APIServer.Stop(); err != nil {
		log.Print(internal.WrapError("Error on 'gRPC-Server-Stop'", err))
	}
}
