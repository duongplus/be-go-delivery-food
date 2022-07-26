package ginrestaurant

import (
	"be-go-delivery-food/common"
	"be-go-delivery-food/component"
	"be-go-delivery-food/modules/restaurant/restaurantbiz"
	"be-go-delivery-food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		//id, err := strconv.Atoi(context.Param("id"))
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurant(store)
		if err := biz.DeleteRestaurant(context.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
