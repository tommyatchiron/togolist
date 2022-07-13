package list

import (
	"github.com/tommyatchiron/togolist/internal/pkg/list/entities"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(NewListService),
	fx.Provide(NewListController),
	fx.Provide(NewListGrpcController),
	fx.Provide(NewListResolver),
	fx.Invoke(registerListRoutes),
	fx.Invoke(registerListGrpcServer),
	fx.Invoke(registerListResolver),
	fx.Invoke(func(db *gorm.DB) {
		db.AutoMigrate(&entities.List{})
	}),
)
