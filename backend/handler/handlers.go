package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"todolist/model"
	"todolist/repository"
)

func GetItems(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := repo.GetAll()
		if err != nil {
			log.Printf("get items failed,err:%v", err)
			ResponseError(c)
			return
		}

		Response(c, items)
	}
}

func AddItem(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		item, err := BindItem(c)
		if err != nil {
			log.Printf("add item bind item failed,err:%v", err)
			ResponseError(c)
			return
		}

		item.ID = uuid.New().String()

		err = repo.Add(item)
		if err != nil {
			log.Printf("add item repository failed,err:%v", err)
			ResponseError(c)
			return
		}

		Response(c, item)
	}
}

func UpdateItem(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		item, err := BindItem(c)
		if err != nil {
			log.Printf("update item bind item failed,err:%v", err)
			ResponseError(c)
			return
		}

		err = repo.Update(item)
		if err != nil {
			log.Printf("update item repository failed,err:%v", err)
			ResponseError(c)
			return
		}

		ResponseSuccess(c)
	}
}

func DeleteItem(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			log.Printf("delete item: id parameter is required")
			ResponseError(c)
			return
		}

		err := repo.Delete(id)
		if err != nil {
			log.Printf("delete item repository failed,err:%v", err)
			ResponseError(c)
			return
		}

		ResponseSuccess(c)
	}
}

func BindItem(c *gin.Context) (*model.TodoItem, error) {
	item := &model.TodoItem{}
	if err := c.ShouldBind(item); err != nil {
		return nil, err
	}
	return item, nil
}

func Response(c *gin.Context, msg any) {
	c.JSON(http.StatusOK, msg)
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}

func ResponseError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "internal server error")
}
