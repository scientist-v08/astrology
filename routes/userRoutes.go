package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/bphs/controller"
	"github.com/scientist-v08/bphs/initializers"
	"github.com/scientist-v08/bphs/token"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 30 * time.Minute
)

func UserRoutes(r *gin.Engine) {
	symmetricKey := initializers.SK
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		log.Fatal("Failed to create token maker:", err)
	}
	authCtrl := controller.NewAuthController(initializers.DB, tokenMaker, accessTokenDuration, refreshTokenDuration)

	r.POST("/register", authCtrl.Register)
	r.POST("/login",    authCtrl.Login)
	r.POST("/refresh", authCtrl.Refresh)
}