package gRPC

import (
	"IoT_Service/proto/service"
	"IoT_Service/server/internal"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Server     *grpc.Server
	Handler    *RoutesHandler
	Listener   net.Listener
	ConfigGrpc internal.GRPC
	ConfigDB   internal.DB
}

func NewServer() *Server {
	return &Server{Server: grpc.NewServer(),
		Handler: NewHandler(),
	}
}

func (s *Server) Initialize() error {
	if err := s.Handler.Initialize(s.ConfigDB); err != nil {
		return internal.WrapError(" Error on 'Handler-Initialize' ", err)
	}
	IoT.RegisterIoTServiceServer(s.Server, s.Handler)
	reflection.Register(s.Server)

	return nil
}

func (s *Server) Run() error {
	serveError := make(chan error)
	err := error(nil)
	s.Listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", s.ConfigGrpc.Host, s.ConfigGrpc.Port))

	if err != nil {
		return err
	}
	go func() {
		defer close(serveError)
		err := s.Server.Serve(s.Listener)
		if err != nil {
			serveError <- err
			return
		}
	}()

	return <-serveError
}

func (s *Server) Stop() error {
	s.Server.GracefulStop()
	if err := s.Handler.DB.Close(); err != nil {
		return internal.WrapError(" Error on 'DB-Close' ", err)
	}

	return nil
}
