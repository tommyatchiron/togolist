package healthz

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewHealthzService),
	fx.Provide(NewHealthzController),
	fx.Invoke(registerHealthzRoutes),
)
