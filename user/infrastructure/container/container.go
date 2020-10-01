package container

import (
	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/translator"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/validator"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/go-micro/handler"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/iris/controller"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/iris/router"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/iris/server"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/repository"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/domain/service"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/adapter"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/interactor"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//--Config
	container.Provide(config.NewConfig)
	container.Provide(config.NewJwtConfig)

	//--Domain
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	//--Usecase
	container.Provide(adapter.NewUserAdapter)
	container.Provide(interactor.NewUserInteractor)

	//--Utility
	container.Provide(database.NewGorm)
	container.Provide(db.NewDbContext)
	container.Provide(translator.NewTranslator)
	container.Provide(validator.NewValidate)

	//--Middleware
	container.Provide(middleware.NewTransactionMiddleware)
	container.Provide(middleware.NewTranslatorMiddleware)
	container.Provide(middleware.NewErrorMiddleware)

	//--Controller
	container.Provide(controller.NewUserController)
	//--Router
	container.Provide(router.NewRouter)
	//--Server
	container.Provide(server.NewServer)

	//--Grpc
	container.Provide(handler.NewUserHandler)

	return container
}
