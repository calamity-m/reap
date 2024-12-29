package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/calamity-m/reap/pkg/errs"
	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/calamity-m/reap/services/sow/internal/persistence"
	"github.com/calamity-m/reap/services/sow/internal/service"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type SowGRPCServer struct {
	// Configured logger allowing for output logs
	log *slog.Logger
	// Service that GRPC client calls will be passed
	// onto for logic handling
	service *service.FoodRecordService
	// List of server options to be applied on Run
	grpcOpts []grpc.ServerOption
	// Address to listen to grpc clients on
	addr string
	// Required to match the GRPC service interface
	sow.UnimplementedFoodRecordingServiceServer
}

func (s *SowGRPCServer) GetRecord(ctx context.Context, req *sow.GetRecordRequest) (*sow.GetRecordResponse, error) {
	s.log.DebugContext(ctx, "entered get record")

	if req == nil {
		return nil, errors.Join(status.Error(codes.InvalidArgument, "request cannot be nil"), errs.ErrBadRequest)
	}

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.Join(status.Errorf(codes.InvalidArgument, "id of %q is not a valid uuid", id), errs.ErrBadRequest)
	}

	record, err := s.service.Get(ctx, id)

	if err != nil {
		if errors.Is(err, errs.ErrInvalidRequest) {
			s.log.ErrorContext(ctx, "invalid request occured", slog.Any("err", err))
			return nil, errors.Join(status.Error(codes.InvalidArgument, "invalid request"), err)
		}

		s.log.ErrorContext(ctx, "some error occured while fetching record", slog.Any("err", err))
		return nil, errors.Join(status.Error(codes.Internal, "internal server error"), err)
	}

	return &sow.GetRecordResponse{Record: record}, nil
}

func (s *SowGRPCServer) GetRecords(wanted *sow.GetRecordsRequest, stream grpc.ServerStreamingServer[sow.GetRecordsResponse]) error {
	s.log.DebugContext(stream.Context(), "entered get records")

	// Verify our provided parameters
	if wanted == nil {
		return errors.Join(status.Error(codes.InvalidArgument, "wanted cannot be nil"), errs.ErrBadRequest)
	}

	userId, err := uuid.Parse(wanted.GetWanted().GetUserId())
	if err != nil {
		return errors.Join(status.Errorf(codes.InvalidArgument, "id of %q is not a valid uuid", userId), errs.ErrBadRequest)
	}

	// Try and get any matching records
	records, err := s.service.GetFiltered(stream.Context(), userId, &sow.Record{})
	if err != nil {
		if errors.Is(err, errs.ErrInvalidRequest) {
			s.log.ErrorContext(stream.Context(), "invalid request occured", slog.Any("err", err))
			return errors.Join(status.Error(codes.InvalidArgument, "invalid request"), err)
		}

		s.log.ErrorContext(stream.Context(), "some error occured while fetching record", slog.Any("err", err))
		return errors.Join(status.Error(codes.Internal, "internal server error"), err)
	}

	// Attempt to send our records through to the requesting client
	for _, record := range records {
		if err := stream.Send(&sow.GetRecordsResponse{Record: record}); err != nil {
			s.log.ErrorContext(stream.Context(), "failed to send through stream", slog.Any("err", err))
			return err
		}
	}

	return nil
}

func (s *SowGRPCServer) CreateRecord(ctx context.Context, req *sow.CreateRecordRequest) (*sow.CreateRecordResponse, error) {
	s.log.DebugContext(ctx, "entered create record")

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *SowGRPCServer) UpdateRecord(ctx context.Context, req *sow.UpdateRecordRequest) (*sow.UpdateRecordResponse, error) {
	s.log.DebugContext(ctx, "entered update record")

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *SowGRPCServer) DeleteRecord(ctx context.Context, req *sow.DeleteRecordRequest) (*sow.DeleteRecordResponse, error) {
	s.log.DebugContext(ctx, "entered delete record")

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *SowGRPCServer) Run() error {
	// Channel we'll use to signal for finish
	exit := make(chan error)

	// create the grpc server that we can later serve on
	grpcServer := grpc.NewServer(s.grpcOpts...)

	// register ourselves
	sow.RegisterFoodRecordingServiceServer(grpcServer, s)

	// enable reflection. Should be configurable later
	reflection.Register(grpcServer)

	go func() {
		// At the end of our function. If no errors were otherwise
		// pushed to this channel, it notifies as a successful shutdown
		defer close(exit)

		sig := make(chan os.Signal, 2)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive a Interrupt or Kill
		<-sig

		grpcServer.GracefulStop()
	}()

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.log.Info(fmt.Sprintf("Starting sow GRPC server on %s", s.addr))
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to start/close sow server due to: %w", err)
	}

	return <-exit
}

func NewSowServer(addr string, logger *slog.Logger) (*SowGRPCServer, error) {
	// This needs configuration to be added, for picking the store

	store := persistence.NewMemoryFoodStore()

	foodService, err := service.NewFoodRecorderService(logger, store)

	if err != nil {
		logger.Error("failed to create sow food service and store")
		return nil, errors.New("failed to create server")
	}

	server := &SowGRPCServer{
		log:     logger,
		service: foodService,
		addr:    addr,
	}

	return server, nil
}
