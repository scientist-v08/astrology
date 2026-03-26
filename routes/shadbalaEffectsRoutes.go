package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/bphs/controller"
	"github.com/scientist-v08/bphs/initializers"
	"github.com/scientist-v08/bphs/middleware"
	"github.com/scientist-v08/bphs/token"
)

func RegisterShadbalaEffectsRoutes(r *gin.Engine) {
	symmetricKey := initializers.SK
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		log.Fatal("Failed to create token maker:", err)
	}
	r.POST("/v1/sthanaBalaOrKaalaBala", middleware.AuthMiddleware(tokenMaker),controller.SthanaBalaOrKaalaBalaCalculator)
	r.POST("/v1/sthanaBalaOrKaalaBalaGreater", middleware.AuthMiddleware(tokenMaker),controller.IsSthanaGreaterOrKaala)
	r.POST("/v1/shadbalaranks", middleware.AuthMiddleware(tokenMaker),controller.PureShadbalaRank)
	r.POST("/v1/digorchestaranks", middleware.AuthMiddleware(tokenMaker),controller.DigBalaOrChestaBalaCalculator)
	r.POST("/v1/ayanaranks", middleware.AuthMiddleware(tokenMaker),controller.AyanaBalaRanks)
}