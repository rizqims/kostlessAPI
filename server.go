package main

import (
	"database/sql"
	"fmt"
	"kostless/config"
	"kostless/controller"
	"kostless/repository"
	"kostless/service"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	tS service.TransService
	engine *gin.Engine
	PortApp string
}

func (s *Server) InitiateRoute(){
	routerGroup := s.engine.Group("/api/v1")
	controller.NewTransController(routerGroup, s.tS).Route()
}

func (s *Server) Start(){
	s.InitiateRoute()
	s.engine.Run(s.PortApp)
}

func NewServer() *Server{
	cn, err := config.LoadConfig()

	urlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cn.Database.Host, cn.Database.Port, cn.Database.User, cn.Database.Password, cn.Database.Name)

	db, err := sql.Open("postgres", urlConn)
	if err != nil {
		log.Fatal(err)
	}

	transRepo := repository.NewTransRepo(db)

	transService := service.NewTransService(transRepo)


	return &Server{
		tS: transService,
		engine:  gin.Default(),
		PortApp: cn.Server.Port,
	}
}