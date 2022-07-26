package main

import (
	"be-go-delivery-food/component"
	"be-go-delivery-food/modules/restaurant/restauranttranport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("restaurant")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	}

	return r.Run(":8080")
}
