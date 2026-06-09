package modules

import (
	"context"

	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/controller"
)

func ConfigureControllers(registry types.ServiceRegistry) error {
	ctx := context.Background()
	
	registry.Register(controller.NewUserController, types.LifetimeTransient)
	features.RegisterList[controller.Controller](ctx, registry)

	return nil
}
