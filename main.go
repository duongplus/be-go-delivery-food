package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	dns := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	fmt.Println(db, "err =>", err)

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restaurants := r.Group("restaurants")
	{
		restaurants.POST("", func(c *gin.Context) {
			var data Restaurant
			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Create(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    data,
			})
		})

		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			var data Restaurant
			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    data,
			})
		})

		restaurants.GET("", func(c *gin.Context) {
			var data []Restaurant

			type Filter struct {
				CityId int `json:"city_id" form:"city_id"`
			}

			var filter Filter

			if err := c.ShouldBind(&filter); err != nil {

			}

			newDb := db

			if filter.CityId > 0 {
				newDb = db.Where("city_id = ?", filter.CityId)
			}

			if err := newDb.Find(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"data":    data,
			})
		})

		restaurants.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data RestaurantUpdate

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"data":    data,
			})
		})
	}
	return r.Run()
}

type Restaurant struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name,omitempty" gorm:"column:name;"`
	Addr string `json:"address,omitempty" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name,omitempty" gorm:"column:name;"`
	Addr *string `json:"address,omitempty" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
