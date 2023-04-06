package todos

import (
	"fmt"
	"gonic-trial/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddTodoDto struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoDto struct {
	Title string `json: "title"`
}

func CreateTodo(c *gin.Context) {
   var input AddTodoDto
   if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H {
		"status": http.StatusBadRequest,
		"message": err.Error(),
	  });

	  return;
   }

   todo := models.Todo {
	Title: input.Title,
   }

   models.DB.Create(&todo)

   c.JSON(http.StatusCreated, gin.H {
	"status": http.StatusCreated,
	"data": todo,
   })
}

func GetTodos(c *gin.Context) {
	var todos [] models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"data": todos,
	})
}

func FindTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"status": http.StatusNotFound,
			"message": "Record not found",
		})

		return;
	}

	c.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"data": todo,
	})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"status": http.StatusNotFound,
			"message": "Record not found",
		})

		return;
	}

	// Validate input

	var input UpdateTodoDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"message": err.Error(),
		})

		return;
	}

	models.DB.Model(&todo).Updates(models.Todo { Title: input.Title })
	message := fmt.Sprintf("Todo Item with ID: '%v' was updated", todo.ID)

	c.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"data": todo,
		"message": message,
	})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"status": http.StatusNotFound,
			"message": "Record not found",
		})

		return;
	}

	models.DB.Delete(&todo)

	message := fmt.Sprintf("Todo Item with ID: '%v' was deleted", todo.ID)

	c.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"message": message,
	})

}