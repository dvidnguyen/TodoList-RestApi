package transport

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/storage"
	"social-todo-list/modules/item/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(db) // đảm bảo  SqlStore implement GetItemStorage interface
		u := usecase.NewGetItemUC(store)

		data, err := u.GetItemByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRepsone(data))
	}
}
