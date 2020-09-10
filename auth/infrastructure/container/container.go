package container

import (
	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/go-micro/handler"
	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/iris/controller"
	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/iris/router"
	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/iris/server"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/repository"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/domain/service"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/adapter"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/interactor"
	gateway "github.com/VulpesFerrilata/boardgame-server/grpc/service"
	"github.com/VulpesFerrilata/boardgame-server/library/config"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/database"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
	transaction_middleware "github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/transaction"
	translator_middleware "github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/translator"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/translator"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/validator"

	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//--Config
	container.Provide(config.NewConfig)
	container.Provide(config.NewJwtConfig)

	//--Domain
	container.Provide(repository.NewAuthRepository)
	container.Provide(service.NewAuthService)
	//--Usecase
	container.Provide(adapter.NewAuthAdapter)
	container.Provide(interactor.NewAuthInteractor)
	//--Gateways
	container.Provide(gateway.NewUserService)

	//--Utility
	container.Provide(database.NewGorm)
	container.Provide(db.NewDbContext)
	container.Provide(translator.NewTranslator)
	container.Provide(validator.NewValidate)

	//--Middleware
	container.Provide(transaction_middleware.NewTransactionMiddleware)
	container.Provide(translator_middleware.NewTranslatorMiddleware)

	//--Controller
	container.Provide(controller.NewAuthController)
	//--Router
	container.Provide(router.NewRouter)
	//--Server
	container.Provide(server.NewServer)

	//--Grpc
	container.Provide(handler.NewAuthHandler)

	return container
}
