package ginrestaurant

import (
	"be-go-delivery-food/common"
	"be-go-delivery-food/component"
	"be-go-delivery-food/modules/restaurant/restaurantbiz"
	"be-go-delivery-food/modules/restaurant/restaurantmodel"
	"be-go-delivery-food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := context.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantStore(store)

		if err := biz.UpdateRestaurant(context.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
