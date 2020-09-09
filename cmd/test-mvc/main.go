package main

import (
	"os"

	"github.com/morihara-y/test-mvc/infrastracuture"

	"github.com/morihara-y/test-mvc/domain/dao"
	"github.com/morihara-y/test-mvc/domain/service"
	"github.com/morihara-y/test-mvc/interfaces/api/server/handler"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

const title = "Testing Server for MVC"
const discription = "Testing Server for MVC"
const version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = title
	app.Usage = discription
	app.Version = version

	app.Action = func(context *cli.Context) {
		execute(context.String("port"))
	}
	app.Flags = getAppFlags()

	app.Run(os.Args)
}

func execute(port string) {
	dbConnetion := infrastracuture.NewDBConnectionImpl()
	messageTrnDao := dao.NewMessageTrnDaoImpl(dbConnetion)
	messageService := service.NewMessageServiceImpl(messageTrnDao)
	messageHandler := handler.NewMessageHandlerImpl(messageService)

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		message := v1.Group("/message")
		message.GET("/all", messageHandler.GetAll)
		message.DELETE("/all", messageHandler.Delete)
		message.POST("/add", messageHandler.Post)
	}
	router.Run(":" + port)
}

func getAppFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "input port number",
		},
	}
}
