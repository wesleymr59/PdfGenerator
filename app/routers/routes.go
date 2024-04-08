package routers

import (
	"fmt"
	"pdfGenerator/app/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	fmt.Println("ta aqui")
	r.GET("/healthy_check", controllers.HealthCheck)
	r.Run(":8080")
}
