package restaurantbiz

import (
	"be-go-delivery-food/modules/restaurant/model"
	"context"
	"errors"
)

type CreateRestaurantStorage interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
}

type createRestaurantBiz struct {
	storage CreateRestaurantStorage
}

func NewCreateRestaurantBiz(storage CreateRestaurantStorage) *createRestaurantBiz {
	return &createRestaurantBiz{storage: storage}
}

func (biz createRestaurantBiz) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("restaurant name can not blank")
	}

	err := biz.storage.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
