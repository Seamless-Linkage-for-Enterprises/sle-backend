package routes

import (
	api "sle/internal"
	"sle/internal/buyer"
	"sle/internal/seller"
	"sle/service/email"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, buyer *buyer.Handler, seller *seller.Handler, email *email.Handler) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// seller routes
	r.POST("/sellers/signup", api.MakeHTTPHandleFunc(seller.SellerSignup))
	r.POST("/sellers/login", api.MakeHTTPHandleFunc(seller.SellerLogin))
	r.POST("/sellers/verify", api.MakeHTTPHandleFunc(email.VerifyOTP))
	r.POST("/sellers/resend", api.MakeHTTPHandleFunc(email.ResendOTP))
	r.GET("/sellers/:sid", api.MakeHTTPHandleFunc(seller.GetSellerByID))
	r.GET("/sellers", api.MakeHTTPHandleFunc(seller.GetAllSellers))
	r.PATCH("/sellers", api.MakeHTTPHandleFunc(seller.SellerForgetPassword))
	r.DELETE("/sellers/:sid", api.MakeHTTPHandleFunc(seller.DeleteSeller))

	// buyer routes
	// signup
	r.POST("/buyers/signup", api.MakeHTTPHandleFunc(buyer.BuyerSignup))
	// login
	r.POST("/buyers/login", api.MakeHTTPHandleFunc(buyer.BuyerLogin))
	// get all buyers
	r.GET("/buyers", api.MakeHTTPHandleFunc(buyer.GetAllBuyers))
	// get a buyer by phone number
	r.GET("/buyers/:bid", api.MakeHTTPHandleFunc(buyer.GetBuyerByID))
	// update buyer details
	r.PATCH("/buyers", api.MakeHTTPHandleFunc(buyer.BuyerForgetPassword))
	// delete buyer
	r.DELETE("/buyers/:bid", api.MakeHTTPHandleFunc(buyer.DeleteBuyer))
	// it check the buyer is verified or not at time of login
	r.GET("/buyers/verify/check/:bid", api.MakeHTTPHandleFunc(buyer.IsBuyerVerified))
	// it verifies the buyer when the otp verification is completed (perform only once)
	r.GET("/buyers/verify/:bid", api.MakeHTTPHandleFunc(buyer.VerifyBuyer))

}
