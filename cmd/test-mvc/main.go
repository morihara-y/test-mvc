package main

import (
	"os"

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
		// execute(context.String("port"))
	}
	app.Flags = getAppFlags()

	app.Run(os.Args)
}

func execute(port string) {
	// userCommentRepo := dao.NewUserCommentRepository()
	// cmntService := service.NewCommentService(userCommentRepo)
	// apiUsecase := usecase.NewAPIUsecase(cmntService)

	router := gin.Default()
	// var cmntHandler handler.CommentHandler
	// v1 := router.Group("/v1")
	// {
	// 	comments := v1.Group("/comments")
	// 	{
	// 		cmntHandler = handler.NewCmntHandler(apiUsecase)
	// 		comments.GET("/:id", cmntHandler.FetchComment)
	// 	}
	// }
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
