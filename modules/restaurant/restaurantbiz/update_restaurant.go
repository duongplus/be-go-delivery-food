package restaurantbiz

import (
	"be-go-delivery-food/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}
type updateRestaurant struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantStore(store UpdateRestaurantStore) *updateRestaurant {
	return &updateRestaurant{store: store}
}

func (biz *updateRestaurant) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
