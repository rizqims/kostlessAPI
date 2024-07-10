package main

import (
	"database/sql"
	"fmt"
	"kostless/config"
	"kostless/controller"
	"kostless/middleware"
	"kostless/repository"
	"kostless/service"
	"kostless/util"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	sS service.SeekerServ
	uS service.UserServ
	jS util.JwtToken
	aM middleware.AuthMiddleware
	engine *gin.Engine
	PortApp string
}

func (s *Server) InitiateRoute(){
	routerGroup := s.engine.Group("/api/v1")
	controller.NewSeekerContr(s.sS, routerGroup, s.aM).Route()
	controller.NewUserContr(s.uS, routerGroup, s.aM).Route()
}

func (s *Server) Start(){
	s.InitiateRoute()
	s.engine.Run(s.PortApp)
}

func NewServer() *Server{
	cn, _ := config.NewConfig()

	urlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cn.Host, cn.Port, cn.User, cn.Password, cn.Name)

	db, err := sql.Open("postgres", urlConn)
	if err != nil {
		log.Fatal(err)
	}
	
	portApp := cn.AppPort
	userRepo := repository.NewUserRepo(db)
	seekerRepo := repository.NewUserSeeker(db)

	jwtUtil :=  util.NewJwtUtil(cn.JwtConfig)
	userService := service.NewUserServ(userRepo, jwtUtil)
	seekerService := service.NewSeekerServ(seekerRepo, jwtUtil)

	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)

	return &Server{
		sS:      seekerService,
		uS:      userService,
		jS: jwtUtil,
		aM: authMiddleware,
		engine:  gin.Default(),
		PortApp: portApp,
	}
}