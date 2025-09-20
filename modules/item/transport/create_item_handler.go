package transport

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/storage"
	"social-todo-list/modules/item/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entity.ToDoItemCreatetion

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)

		usecase := usecase.NewCreateItemUsecase(store)

		if err := usecase.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRepsone(data.Id))
	}
}
