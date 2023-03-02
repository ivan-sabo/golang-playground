package grpc

import (
	"github.com/ivan-sabo/golang-playground/api/grpc/proto"
	"github.com/ivan-sabo/golang-playground/api/grpc/server"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func NewGRPCServer(conn *gorm.DB) (*grpc.Server, error) {
	gsrv := grpc.NewServer()

	championshipSrv, err := server.NewGrpcServer(conn)
	if err != nil {
		return nil, err
	}
	proto.RegisterChampionshipServiceServer(gsrv, championshipSrv)

	return gsrv, nil
}
