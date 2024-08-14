package main

import (
	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/router"
	"rest-api/usecase"
	"rest-api/validator"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
