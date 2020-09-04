package container

import (
	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database/gorm"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
	transaction_middleware "github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/transaction"
	translator_middleware "github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/translator"
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
	container.Provide(config.NewDatabaseConfig)
	container.Provide(config.NewJwtConfig)

	//--Domain
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	//--Usecase
	container.Provide(adapter.NewUserAdapter)
	container.Provide(interactor.NewUserInteractor)

	//--Utility
	container.Provide(gorm.NewGorm)
	container.Provide(db.NewDbContext)
	container.Provide(translator.NewTranslator)
	container.Provide(validator.NewValidate)

	//--Middleware
	container.Provide(transaction_middleware.NewTransactionMiddleware)
	container.Provide(translator_middleware.NewTranslatorMiddleware)

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
