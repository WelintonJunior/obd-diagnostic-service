package grpcservice

import (
	"context"

	pb "github.com/WelintonJunior/obd-diagnostic-service/proto"
)

type DiagnosticService struct {
	pb.UnimplementedDiagnosticsServer
}

func (s *DiagnosticService) Ping(ctx context.Context, in *pb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "pong"}, nil
}
