package main

import (
	"database/sql"
	"fmt"
	"kostless-api/config"
	"kostless-api/controller"
	"kostless-api/repository"
	"kostless-api/service"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	sS service.SeekerServ
	uS service.UserServ
	engine *gin.Engine
	PortApp string
}

func (s *Server) InitiateRoute(){
	routerGroup := s.engine.Group("/api/v1")
	controller.NewSeekerContr(s.sS, routerGroup).Route()
	controller.NewUserContr(s.uS, routerGroup).Route()
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

	userRepo := repository.NewUserRepo(db)
	seekerRepo := repository.NewUserSeeker(db)

	userService := service.NewUserServ(userRepo)
	seekerService := service.NewSeekerServ(seekerRepo)

	return &Server{
		sS:      seekerService,
		uS:      userService,
		engine:  gin.Default(),
		PortApp: cn.Server.Port,
	}
}