package healthz

import (
	"fmt"

	"github.com/tommyatchiron/togolist/internal/pkg/healthz/dto"
	"gorm.io/gorm"
)

type HealthzService struct {
	db *gorm.DB
}

func NewHealthzService(db *gorm.DB) *HealthzService {
	return &HealthzService{db}
}

func (hs *HealthzService) HealthCheck() (*dto.HealthzResult, error) {
	db, err := hs.db.DB()
	if err != nil {
		return nil, fmt.Errorf("fail to get underlying db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return &dto.HealthzResult{
			Status: "error",
			Details: dto.HealthzResultDetails{
				Db: dto.HealthzResultDetail{
					Status: "error",
				},
			},
		}, nil
	}
	return &dto.HealthzResult{
		Status: "ok",
		Details: dto.HealthzResultDetails{
			Db: dto.HealthzResultDetail{
				Status: "ok",
			},
		},
	}, nil
}
