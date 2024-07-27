// Code generated by RedSock CLI. DO NOT EDIT

package grpc

import (
	"context"
	"net"
	"net/http"
	"strings"
	"sync"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka/servers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	"github.com/godverv/makosh/internal/config"
	"github.com/godverv/makosh/internal/interceptors"
)

type Server struct {
	serverMux cmux.CMux

	grpcServer *grpc.Server
	imps       []Implementation

	serverAddress string
	m             sync.Mutex
}

type Implementation interface {
	Register(server grpc.ServiceRegistrar)
	RegisterGw(ctx context.Context, mux *runtime.ServeMux, addr string) error
}

func NewServer(
	cfg config.Config,
	server *servers.GRPC,
	imps ...Implementation,
) (*Server, error) {

	var opts []grpc.ServerOption

	opts = append(opts, interceptors.GrpcInterceptor(cfg.GetEnvironment().AuthToken))

	grpcServer := grpc.NewServer(opts...)

	for _, imp := range imps {
		imp.Register(grpcServer)
	}

	return &Server{
		grpcServer:    grpcServer,
		serverAddress: "0.0.0.0:" + server.GetPortStr(),
		imps:          imps,
	}, nil
}

func (s *Server) Start(_ context.Context) error {
	s.m.Lock()
	defer s.m.Unlock()

	listener, err := net.Listen("tcp", s.serverAddress)
	if err != nil {
		return errors.Wrap(err, "error opening listener")
	}
	s.serverMux = cmux.New(listener)

	go s.startGRPC()

	go s.startGateway()

	go func() {
		serveErr := s.serverMux.Serve()
		if serveErr != nil {
			if !strings.Contains(serveErr.Error(), "closed network connection") {
				logrus.Errorf("error service server %s", serveErr)
			}
		}
	}()

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	s.grpcServer.GracefulStop()
	logrus.Infof("Server at %s is stopped", s.serverAddress)

	return nil
}

func (s *Server) startGRPC() {
	grpcListener := s.serverMux.Match(cmux.HTTP2())

	logrus.Infof("Starting server at %s", s.serverAddress)

	err := s.grpcServer.Serve(grpcListener)
	if err != nil {
		logrus.Errorf("error starting grpc server: %s", err)
	}
}

func (s *Server) startGateway() {
	httpListener := s.serverMux.Match(cmux.HTTP1Fast())

	httpMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{}))

	ctx := context.Background()
	for _, imp := range s.imps {
		err := imp.RegisterGw(ctx, httpMux, s.serverAddress)
		if err != nil {
			logrus.Errorf("error registering grpc2http handler: %s", err)
		}
	}

	server := &http.Server{
		Addr:    s.serverAddress,
		Handler: coverWithCors(httpMux),
	}

	err := server.Serve(httpListener)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("error starting gateway: %s", err)
		}
	}
}

func coverWithCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string) bool {
	//if viper.GetString("cors") == "*" {
	return true
	//}
	//if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
	//	return true
	//}
	//return false
}
