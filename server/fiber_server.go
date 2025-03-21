package server

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"

	"go.db.restapi/config"
	ctrl "go.db.restapi/controller"
)

type FiberServer struct {
	app     *fiber.App
	user    *ctrl.UserController
	started bool
}

var _ bootstrap.Application = &FiberServer{}

func NewFiberServer() Server {
	return &FiberServer{}
}

func NewApp() bootstrap.Application {
	return &FiberServer{}
}

// Init method boots the end-points for the server
func (fs *FiberServer) Init() error {
	if fs.started {
		err := config.ReadTOML()
		if err != nil {
			log.Fatal(err)
			return err
		}
		if strings.ToLower(config.TOMLConfig.App.JsonProcessor) == "sonic" {
			fs.app = fiber.New(fiber.Config{
				JSONEncoder: sonic.Marshal,
				JSONDecoder: sonic.Unmarshal,
			})
		} else {
			fs.app = fiber.New()
		}
		fs.user = &ctrl.UserController{}

		fs.user.Init(fs.app)
		port := strconv.Itoa(config.TOMLConfig.App.Port)
		err = fs.app.Listen(":" + port)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fs.started = true
	}
	return nil
}

// Run The entrypoint for the Parsley application.
func (fs *FiberServer) Run(_ context.Context) error {
	return fs.Init()
}
