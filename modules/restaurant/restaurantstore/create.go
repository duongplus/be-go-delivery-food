package restaurantstore

import (
	"be-go-delivery-food/modules/restaurant/model"
	"context"
)

func (s *restaurantSqlStore) Create(ctx context.Context, data *model.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
