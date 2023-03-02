package mapper

import (
	"github.com/google/uuid"
	"github.com/ivan-sabo/golang-playground/api/grpc/model"
	"github.com/ivan-sabo/golang-playground/internal/championship/domain"
)

func NewChampionship(dc domain.Championship) *model.Championship {
	return &model.Championship{
		Id:   dc.ID.String(),
		Name: dc.Name,
	}
}

func NewChampionships(dc domain.Championships) []*model.Championship {
	cs := make([]*model.Championship, 0, len(dc))
	for _, c := range dc {
		cp := NewChampionship(c)
		cs = append(cs, cp)
	}
	return cs
}

func ChampionshipToEntity(c *model.Championship) (domain.Championship, error) {
	id, err := uuid.Parse(c.GetId())
	if err != nil {
		return domain.Championship{}, err
	}

	return domain.Championship{
		ID:   id,
		Name: c.GetName(),
	}, nil
}
