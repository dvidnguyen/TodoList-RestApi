package main

import (
	"log"
	"net/http"
	"os"
	"social-todo-list/common"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/transport"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

<<<<<<< HEAD
=======
func (TodoItem) TableName() string { return "todo_items" }

type ToDoItemCreatetion struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (ToDoItemCreatetion) TableName() string { return TodoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (ToDoItemUpdate) TableName() string { return TodoItem{}.TableName() }

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 20 {
		p.Limit = 10
	}
}
>>>>>>> a0d2b0f08439de8845809c52bb1225d72ff834a3
func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Ham marshall parse convert struct to json
	//jsonData, err := json.Marshal(item)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(jsonData))
	r := gin.Default()
	// CRUD Create, Read, Update, Delete
	// POST /v1/items (create new item)
	// GET /v1/items (get list items)
	// GET /v1/items/:id (get detail item by id )
	// PUT || PATCH /v1/items/:id (update item by id)
	// DELETE /v1/items/:id ( delete id by id )

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
<<<<<<< HEAD
			items.POST("", transport.CreateItem(db))
			items.GET("/:id", transport.GetItem(db))
			items.PATCH("/:id", transport.UpdateItem(db))
=======
			items.POST("", CreateItem(db))
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
>>>>>>> a0d2b0f08439de8845809c52bb1225d72ff834a3
			items.DELETE("/:id", DeleteItem(db))
			items.GET("", ListItem(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.SimpleSuccessRepsone("pongggggg"))
	})
	r.Run(":8080")
}

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//data.Id = id

		// Hard Delete
		//if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete().Error; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}

		//data.Id = id
		// Delete soft
<<<<<<< HEAD
		if err := db.Table(entity.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
=======
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
>>>>>>> a0d2b0f08439de8845809c52bb1225d72ff834a3
			"status": "Deleted",
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRepsone(true))
	}

}
func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// handle value default for paging
		paging.Process()

<<<<<<< HEAD
		var results []entity.TodoItem
		db = db.Where("status <> ?", "Deleted")
		if err := db.Table(entity.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
=======
		var results []TodoItem
		db = db.Where("status <> ?", "Deleted")
		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
>>>>>>> a0d2b0f08439de8845809c52bb1225d72ff834a3
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Order("id ").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&results).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SuccessRepsone(results, paging, nil))
	}
}
