package config

import (
	"sle/internal/buyer"
	"sle/internal/product"
	"sle/internal/seller"
	"sle/routes"
	"sle/service/email"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var r = gin.Default()

func Configuration(db *pgxpool.Pool) email.Handler {

	emailRepo := email.NewEmailRespository(db)
	emailServ := email.NewEmailService(emailRepo)
	emailHand := email.NewEmailHandler(emailServ)

	// initialize seller
	sellerRepo := seller.NewSellerRepository(db)
	sellerServ := seller.NewSellerService(sellerRepo, emailRepo)
	sellerHand := seller.NewSellerHandler(sellerServ)

	// initialize buyer
	buyerRepo := buyer.NewBuyerRepository(db)
	buyerServ := buyer.NewBuyerService(buyerRepo)
	buyerHand := buyer.NewBuyerHandler(buyerServ)

	// initialize product
	productRepo := product.NewProductRepository(db)
	productServ := product.NewProductService(productRepo)
	productHand := product.NewProductHandler(productServ)

	routes.SetupRoutes(r, buyerHand, sellerHand, productHand, &emailHand)

	return emailHand
}

func RunServer(listenAddr string) {
	r.Run(listenAddr)
}
