package restApi

import (
	"IoT_Service/proto/service"
	"IoT_Service/server/internal"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	Router     *mux.Router
	grpcClient IoT.IoTServiceClient
	grpcCtx    context.Context
	Config     internal.API
}

func New() *Server {
	return &Server{Router: mux.NewRouter(),
		grpcCtx: context.Background(),
	}
}

func (s *Server) InitializeRoutes() error {
	err := s.Router.HandleFunc("/device", s.createDevice).Methods("POST").GetError()
	err = s.Router.HandleFunc("/device", s.updateDevice).Methods("PATCH").GetError()
	err = s.Router.HandleFunc("/device", s.deleteDevice).Methods("DELETE").GetError()
	err = s.Router.HandleFunc("/device", s.getDevice).Methods("GET").GetError()
	err = s.Router.HandleFunc("/devices", s.getDevices).Methods("GET").GetError()

	err = s.Router.HandleFunc("/sensor", s.createSensor).Methods("POST").GetError()
	err = s.Router.HandleFunc("/sensor", s.updateSensor).Methods("PATCH").GetError()
	err = s.Router.HandleFunc("/sensor", s.deleteSensor).Methods("DELETE").GetError()
	err = s.Router.HandleFunc("/sensor", s.getSensor).Methods("GET").GetError()
	err = s.Router.HandleFunc("/sensors", s.getSensors).Methods("GET").GetError()

	err = s.Router.HandleFunc("/record", s.createRecord).Methods("POST").GetError()
	err = s.Router.HandleFunc("/record", s.updateRecord).Methods("PATCH").GetError()
	err = s.Router.HandleFunc("/record", s.deleteRecord).Methods("DELETE").GetError()
	err = s.Router.HandleFunc("/record", s.getRecord).Methods("GET").GetError()
	err = s.Router.HandleFunc("/records", s.getRecords).Methods("GET").GetError()

	err = s.Router.HandleFunc("/record_data", s.createRecordData).Methods("POST").GetError()
	err = s.Router.HandleFunc("/record_data", s.deleteRecordData).Methods("DELETE").GetError()
	err = s.Router.HandleFunc("/record_data", s.getRecordData).Methods("GET").GetError()

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Initialize() error {
	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%s", s.Config.GrpcHost, s.Config.GrpcPort), grpc.WithInsecure())
	if err != nil {
		return internal.WrapError(" Error on 'Connect to GRPC-Server' ", err)
	}
	s.grpcClient = IoT.NewIoTServiceClient(grpcConn)
	if err := s.InitializeRoutes(); err != nil {
		return internal.WrapError(" Error on 'InitializeRoutes' ", err)
	}

	return nil
}

func (s *Server) Run() error {
	serveError := make(chan error)

	go func() {
		defer close(serveError)
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port), s.Router)
		if err != nil {
			serveError <- err
			return
		}
	}()

	return <-serveError
}

func (s *Server) Stop() error {
	return nil
}
