package server

import (
	"context"

	"github.com/ivan-sabo/golang-playground/api/grpc/mapper"
	"github.com/ivan-sabo/golang-playground/api/grpc/proto"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

var _ proto.ChampionshipServiceServer = (*championshipServer)(nil)

type championshipServer struct {
	proto.UnimplementedChampionshipServiceServer

	championshipRepo domain.ChampionshipRepo
}

func NewGrpcServer(conn *gorm.DB) (srv *championshipServer, err error) {
	srv = &championshipServer{
		championshipRepo: repository.NewChampionshipMySQLRepo(conn),
	}

	return srv, nil
}

func (s *championshipServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	newChampionship, err := domain.NewChampionship("", req.Championship.GetName())
	if err != nil {
		return nil, err
	}

	c, err := s.championshipRepo.CreateChampionship(newChampionship)
	if err != nil {
		return nil, err
	}

	return &proto.CreateResponse{
		Championship: &proto.Championship{
			Id:   c.ID.String(),
			Name: c.Name,
		},
	}, nil
}

func (s *championshipServer) GetSingle(ctx context.Context, req *proto.GetSingleRequest) (*proto.GetSingleResponse, error) {
	c, err := s.championshipRepo.GetChampionship(req.GetUUID())
	if err != nil {
		return nil, err
	}

	return &proto.GetSingleResponse{
		Championship: &proto.Championship{
			Id:   c.ID.String(),
			Name: c.Name,
		},
	}, nil
}

func (s *championshipServer) GetList(ctx context.Context, req *proto.GetListRequest) (*proto.GetListResponse, error) {
	// @todo: implement filter
	dcs, err := s.championshipRepo.GetChampionships(domain.ChampionshipFilter{})
	if err != nil {
		return nil, err
	}

	cs := mapper.NewChampionships(dcs)

	return &proto.GetListResponse{
		Championships: cs,
	}, nil
}

func (s *championshipServer) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	dc, err := mapper.ChampionshipToEntity(req.Championship)
	if err != nil {
		return nil, err
	}

	r, err := s.championshipRepo.UpdateChampionship(dc.ID.String(), dc)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateResponse{
		Championship: mapper.NewChampionship(r),
	}, nil
}

func (s *championshipServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	err := s.championshipRepo.DeleteChampionship(req.UUID)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteResponse{}, nil
}
