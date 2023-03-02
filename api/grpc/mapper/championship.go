package mapper

import (
	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/api/grpc/proto"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
)

func NewChampionship(dc domain.Championship) *proto.Championship {
	return &proto.Championship{
		Id:   dc.ID.String(),
		Name: dc.Name,
	}
}

func NewChampionships(dc domain.Championships) []*proto.Championship {
	cs := make([]*proto.Championship, 0, len(dc))
	for _, c := range dc {
		cp := NewChampionship(c)
		cs = append(cs, cp)
	}
	return cs
}

func ChampionshipToEntity(c *proto.Championship) (domain.Championship, error) {
	id, err := uuid.Parse(c.GetId())
	if err != nil {
		return domain.Championship{}, err
	}

	return domain.Championship{
		ID:   id,
		Name: c.GetName(),
	}, nil
}
