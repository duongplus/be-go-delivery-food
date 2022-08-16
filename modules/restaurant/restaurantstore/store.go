package restaurantstore

import "gorm.io/gorm"

type restaurantSqlStore struct {
	db *gorm.DB
}

func NewRestaurantSqlStore(db *gorm.DB) *restaurantSqlStore {
	return &restaurantSqlStore{db: db}
}
