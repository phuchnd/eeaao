package grpc

//import (
//	"fmt"
//	"github.com/phuchnd/simple-go-boilerplate/internal/config"
//	"github.com/phuchnd/simple-go-boilerplate/internal/cron"
//	"github.com/phuchnd/simple-go-boilerplate/internal/cron/jobs"
//	mysqldb "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
//	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories"
//	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
//	"github.com/phuchnd/simple-go-boilerplate/internal/generators"
//	"github.com/phuchnd/simple-go-boilerplate/internal/health"
//	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
//	grpc2 "github.com/phuchnd/simple-go-boilerplate/internal/service/grpc"
//	"github.com/phuchnd/simple-go-boilerplate/server/grpc/endpoints"
//	"github.com/phuchnd/simple-go-boilerplate/server/grpc/middlewares"
//	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
//	"google.golang.org/grpc"
//	"net"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
////go:generate mockery --name=IServer --case=snake --disable-version-string
//type IServer interface {
//	Start() error
//	Stop() error
//}
//
//type serverImpl struct {
//	grpcServer pb.SimpleGoBoilerplateServiceServer
//	runner     cron.IJobRunner
//
//	serverCfg *config.ServerConfig
//	logger    logging.Logger
//
//	healthCheckSvc health.IHealthCheck
//	isReady        bool
//	quit           chan os.Signal
//}
//
//func NewServer(logger logging.Logger, cfgProvider config.IConfig, db mysqldb.IMySqlDB) (IServer, error) {
//	serverCfg := cfgProvider.GetServerConfig()
//
//	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverCfg.GRPCPort))
//	if err != nil {
//		return nil, err
//	}
//	sf, err := generators.NewSnowflakeIDGenerator()
//	if err != nil {
//		return nil, err
//	}
//	idGenerator, err := entities.NewIDGenerator(sf)
//	if err != nil {
//		return nil, err
//	}
//
//	dbConfig := cfgProvider.GetDBConfig()
//	bookRepo, err := repositories.NewBookRepository(db.DB(), idGenerator, dbConfig.MySQL)
//	if err != nil {
//		return nil, err
//	}
//	serviceHandler := grpc2.NewGRPCService(bookRepo)
//	grpcEndpoints := endpoints.MakeEndPoints(serviceHandler,
//		middlewares.Tracing(),
//		middlewares.RequestLogging(),
//		middlewares.PanicRecovery(),
//	)
//	healthCheckSvc := health.NewHealthCheck(db, logger)
//	server := &serverImpl{
//		grpcServer: NewGRPCServer(grpcEndpoints),
//		runner:     initJobRunner(logger, cfgProvider),
//
//		serverCfg:      serverCfg,
//		logger:         logger,
//		healthCheckSvc: healthCheckSvc,
//	}
//	// Init jobs
//	server.runner.Start()
//
//	baseServer := grpc.NewServer()
//	pb.RegisterSimpleGoBoilerplateServiceServer(baseServer, server.grpcServer)
//
//	logger.Infof("gRPC server is running on :%d", serverCfg.GRPCPort)
//	if err := baseServer.Serve(grpcListener); err != nil {
//		logger.Errorf("failed to start gRPC server", "err", err)
//		return nil, err
//	}
//
//	return server, nil
//}
//
//func initJobRunner(logger logging.Logger, cfgProvider config.IConfig) cron.IJobRunner {
//	cronSystem := cron.NewCron(logger)
//	runner := cron.NewJobRunner(logger, cronSystem)
//	cfg := cfgProvider.GetCronSimpleExampleConfig()
//	simpleExampleJob := jobs.NewCronSimpleExample(logger, cfg)
//	_ = runner.RegisterJob(simpleExampleJob)
//
//	return runner
//}
//
//func (s *serverImpl) Start() error {
//	if err := s.healthCheckSvc.Check(); err != nil {
//		s.logger.Errorf("GRPC server initialization failed, dependency not ready: %s", err.Error())
//		return err
//	}
//
//	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.serverCfg.GRPCPort))
//	if err != nil {
//		s.logger.Errorf("GRPC server initialization failed: %s", err.Error())
//		return err
//	}
//	s.isReady = s.healthCheckSvc.IsReady()
//
//	go func() {
//		baseServer := grpc.NewServer()
//		pb.RegisterSimpleGoBoilerplateServiceServer(baseServer, s.grpcServer)
//		s.logger.Infof("GRPC server started successfully at port %d", s.serverCfg.GRPCPort)
//
//		if err := baseServer.Serve(grpcListener); err != nil {
//			s.logger.Errorf("failed to start GRPC server %s", err.Error())
//		}
//	}()
//
//	s.quit = make(chan os.Signal)
//	signal.Notify(s.quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
//	<-s.quit
//
//	s.logger.Info("received terminated signal, server is shutting down")
//	return s.Stop()
//}
//
//func (s *serverImpl) Stop() error {
//	return nil
//}
