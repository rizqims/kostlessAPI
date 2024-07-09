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
	kS      service.KosService
	rS      service.RoomService
	tS      service.TransService
	sS      service.SeekerServ
	uS      service.UserServ
	jS      util.JwtToken
	aM      middleware.AuthMiddleware
	engine  *gin.Engine
	portApp string
}

func (s *Server) initiateRoute() {
	routerGroup := s.engine.Group("/api/v1")
	controller.NewKosController(s.kS, routerGroup).Route()
	controller.NewRoomController(s.rS, routerGroup).Route()
	controller.NewTransController(routerGroup, s.tS).Route()
	controller.NewSeekerContr(s.sS, routerGroup).Route()
	controller.NewUserContr(s.uS, routerGroup).Route()
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(s.portApp)
}

func NewServer() *Server {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Name)

	db, err := sql.Open("postgres", urlConnection)
	if err != nil {
		log.Fatal(err)
	}
	portApp := conf.AppPort
	kosRepo := repository.NewKosRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	transRepo := repository.NewTransRepo(db)
	userRepo := repository.NewUserRepo(db)
	seekerRepo := repository.NewUserSeeker(db)
	jwtUtil := util.NewJwtUtil(conf.JwtConfig)

	kosService := service.NewKosService(kosRepo)
	roomService := service.NewRoomService(roomRepo)
	transService := service.NewTransService(transRepo)
	userService := service.NewUserServ(userRepo, jwtUtil)
	seekerService := service.NewSeekerServ(seekerRepo, jwtUtil)

	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)

	return &Server{
		portApp: portApp,
		kS:      kosService,
		rS:      roomService,
		tS:      transService,
		sS:      seekerService,
		uS:      userService,
		jS:      jwtUtil,
		aM:      authMiddleware,
		engine:  gin.Default(),
	}
}