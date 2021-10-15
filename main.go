package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

// func Home(ctx *gin.Context) {
// 	ctx.JSON(200, gin.H{ "message": "HomePage",})
// }

func main() {

	httpRouter := gin.Default()

	webSocketRouter := melody.New()

	webSocketRouter.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	log.Println("ok")

	httpRouter.GET("/echo", func (ctx *gin.Context) {

		webSocketRouter.HandleRequest(ctx.Writer, ctx.Request)

	})

	webSocketRouter.HandleMessage(func (s *melody.Session, m []byte) {

		webSocketRouter.Broadcast(m)
	})

	httpRouter.Use(static.Serve("/", static.LocalFile("./public", true)))


	httpRouter.Run(":8080")
}
