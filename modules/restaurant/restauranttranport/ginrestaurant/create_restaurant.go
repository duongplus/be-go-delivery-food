package ginrestaurant

import (
	"be-go-delivery-food/common"
	"be-go-delivery-food/component"
	"be-go-delivery-food/modules/restaurant/restaurantbiz"
	"be-go-delivery-food/modules/restaurant/restaurantmodel"
	"be-go-delivery-food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
