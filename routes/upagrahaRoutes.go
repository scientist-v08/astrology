package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/bphs/controller"
	"github.com/scientist-v08/bphs/initializers"
	"github.com/scientist-v08/bphs/middleware"
	"github.com/scientist-v08/bphs/token"
)

func RegisterUpagrahaRoutes(r *gin.Engine) {
	symmetricKey := initializers.SK
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		log.Fatal("Failed to create token maker:", err)
	}
	r.POST("/v1/upagrahas", middleware.AuthMiddleware(tokenMaker), controller.UpagraHandler)
}