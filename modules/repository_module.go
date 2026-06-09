package modules

import (
	"context"

	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/model"
	"go.db.restapi/repository"
)

func ConfigureRepositories(registry types.ServiceRegistry) error {
	ctx := context.Background()

	registry.Register(repository.NewUserMongoRepository, types.LifetimeTransient)

	registry.Register(repository.NewUserRedisRepository, types.LifetimeTransient)
	registry.Register(repository.NewUserValkeyRepository, types.LifetimeTransient)
	features.RegisterList[repository.UserCacheRepository[model.User]](ctx, registry)

	return nil
}
