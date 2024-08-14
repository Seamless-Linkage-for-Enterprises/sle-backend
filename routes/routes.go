package routes

import (
	api "sle/internal"
	"sle/internal/seller"
	"sle/service/email"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, seller *seller.Handler, email *email.Handler) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// seller routes
	r.POST("/sellers/signup", api.MakeHTTPHandleFunc(seller.SellerSignup))
	r.POST("/sellers/login", api.MakeHTTPHandleFunc(seller.SellerLogin))
	r.POST("/sellers/verify", api.MakeHTTPHandleFunc(email.VerifyOTP))
	r.GET("/sellers/:sid", api.MakeHTTPHandleFunc(seller.GetSellerByID))
	r.GET("/sellers", api.MakeHTTPHandleFunc(seller.GetAllSellers))
	r.PATCH("/sellers", api.MakeHTTPHandleFunc(seller.SellerForgetPassword))
}
