package grpc

import (
	"context"

	"github.com/ivan-sabo/golang-playground/api/grpc/mapper"
	"github.com/ivan-sabo/golang-playground/api/grpc/model"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
	"github.com/ivan-sabo/golang-playground/internal/championship/infrastructure/database/mysql/repository"
	"gorm.io/gorm"
)

type Config struct {
	conn *gorm.DB
}

var _ model.ChampionshipServiceServer = (*grpcServer)(nil)

type grpcServer struct {
	model.UnimplementedChampionshipServiceServer
	*Config
	championshipRepo domain.ChampionshipRepo
}

func NewGrpcServer(config *Config) (srv *grpcServer, err error) {
	srv = &grpcServer{
		Config:           config,
		championshipRepo: repository.NewChampionshipMySQLRepo(config.conn),
	}

	return srv, nil
}

func (s *grpcServer) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	newChampionship, err := domain.NewChampionship("", req.Championship.GetName())
	if err != nil {
		return nil, err
	}

	c, err := s.championshipRepo.CreateChampionship(newChampionship)
	if err != nil {
		return nil, err
	}

	return &model.CreateResponse{
		Championship: &model.Championship{
			Id:   c.ID.String(),
			Name: c.Name,
		},
	}, nil
}

func (s *grpcServer) GetSingle(ctx context.Context, req *model.GetSingleRequest) (*model.GetSingleResponse, error) {
	c, err := s.championshipRepo.GetChampionship(req.GetUUID())
	if err != nil {
		return nil, err
	}

	return &model.GetSingleResponse{
		Championship: &model.Championship{
			Id:   c.ID.String(),
			Name: c.Name,
		},
	}, nil
}

func (s *grpcServer) GetList(ctx context.Context, req *model.GetListRequest) (*model.GetListResponse, error) {
	// @todo: implement filter
	dcs, err := s.championshipRepo.GetChampionships(domain.ChampionshipFilter{})
	if err != nil {
		return nil, err
	}

	cs := mapper.NewChampionships(dcs)

	return &model.GetListResponse{
		Championships: cs,
	}, nil
}

func (s *grpcServer) Update(ctx context.Context, req *model.UpdateRequest) (*model.UpdateResponse, error) {
	dc, err := mapper.ChampionshipToEntity(req.Championship)
	if err != nil {
		return nil, err
	}

	r, err := s.championshipRepo.UpdateChampionship(dc.ID.String(), dc)
	if err != nil {
		return nil, err
	}

	return &model.UpdateResponse{
		Championship: mapper.NewChampionship(r),
	}, nil
}

func (s *grpcServer) Delete(ctx context.Context, req *model.DeleteRequest) (*model.DeleteResponse, error) {
	err := s.championshipRepo.DeleteChampionship(req.UUID)
	if err != nil {
		return nil, err
	}

	return &model.DeleteResponse{}, nil
}
