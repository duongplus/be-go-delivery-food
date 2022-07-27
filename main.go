package main

import (
	"be-go-delivery-food/component"
	"be-go-delivery-food/component/uploadprovider"
	"be-go-delivery-food/middleware"
	"be-go-delivery-food/modules/restaurant/restauranttranport/ginrestaurant"
	"be-go-delivery-food/modules/upload/uploadtranport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//AKIA5JDCNTI4GZG2NYXG accessKEY
//63yM3qFghudaITjE4p6USH5V5ZToKXQIFqGWBCpp secretKEY
//delivery-food
//ap-southeast-1
//https://dakum0xg9bbry.cloudfront.net

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "pong",
		})
	})

	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run(":8080")
}
