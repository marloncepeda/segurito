package main

import (
    "context"
    "log"
    "src/cfg"
    "src/repositories/quotationrepo"
    "src/services/quotationservice"
    "src/repositories/planrepo"
    "src/services/planservice"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "src/controllers"
)


func main() {
	router := gin.Default()
    if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
    cfgs := cfg.GetConfig()

    // Set client options
	clientOpt := options.Client().ApplyURI(cfgs.MongoDB.URI)
	// Connect to MongoDB
	mongoDB, err := mongo.Connect(context.Background(), clientOpt)
    if err != nil {
		panic(err)
	}

    //repositories
    planRepo := planrepo.NewPlanRepo(mongoDB)
    quotationRepo := quotationrepo.NewQuotationRepo(mongoDB)

    //services
    planService := planservice.NewPlanService(planRepo)
    quotationService := quotationservice.NewQuotationService(quotationRepo)

    //controllers
    planCtl := controllers.NewPlanController(planService)
    quotationCtl := controllers.NewQuotationController(quotationService, planService)

    //Routes
    v1 := router.Group("/api/v1")

    v1.POST("/plans", planCtl.PostPlan)
    v1.GET("/plans", planCtl.GetPlan)
    v1.POST("/quotations", quotationCtl.PostQuotation)
	router.Run()
}
