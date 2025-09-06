package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"todolist/handler"
	"todolist/middleware"
	"todolist/repository"
)

func main() {
	s := NewServer()

	s.Setup()

	s.Run()
}

type Server struct {
	engine *gin.Engine
	repo   repository.Repository
}

func NewServer() Server {
	return Server{engine: gin.Default(),
		repo: repository.NewRepository()}
}

func (s *Server) Setup() {
	s.engine.Use(middleware.CORS())
	s.engine.GET("/api/v1/items", handler.GetItems(s.repo))
	s.engine.POST("/api/v1/item", handler.AddItem(s.repo))
	s.engine.PUT("/api/v1/item", handler.UpdateItem(s.repo))
	s.engine.DELETE("/api/v1/item/:id", handler.DeleteItem(s.repo))
}

func (s *Server) Run() {
	if err := s.engine.Run(":8088"); err != nil {
		log.Panicf("server run failed,err:%v", err)
	}
}
