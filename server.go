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
	kS      service.KosService
	rS      service.RoomService
	engine  *gin.Engine
	portApp string
}

func (s *Server) initiateRoute() {
	routerGroup := s.engine.Group("/api/v1")
	controller.NewKosController(s.kS, routerGroup).Route()
	controller.NewRoomController(s.rS, routerGroup).Route()
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(s.portApp)
}

func NewServer() *Server {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name)

	db, err := sql.Open("postgres", urlConnection)
	if err != nil {
		log.Fatal(err)
	}
	portApp := conf.Server.Port
	kosRepo := repository.NewKosRepository(db)
	roomRepo := repository.NewRoomRepository(db)

	kosService := service.NewKosService(kosRepo)
	roomService := service.NewRoomService(roomRepo)

	return &Server{
		portApp: portApp,
		kS:      kosService,
		rS:      roomService,
		engine:  gin.Default(),
	}
}
