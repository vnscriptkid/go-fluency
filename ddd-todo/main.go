package main

import (
	"ddd/domain"
	repo "ddd/repo"
	server "ddd/server"
	service "ddd/service"
	"log"
)

func main() {

	userRepo := &repo.UserRepoInMem{}
	taskRepo := &repo.TaskRepoInMem{}

	firstUser := domain.User{
		ID:       1,
		Username: "thanh",
	}
	userRepo.Save(&firstUser)

	userSvc := service.NewUserService(userRepo)
	taskSvc := service.NewTaskService(userRepo, taskRepo)

	httpServer := server.NewHttpServer(userSvc, taskSvc)

	err := httpServer.Start()

	if err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}
