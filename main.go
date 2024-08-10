package main

import (
	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/router"
	"rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)

	e.Logger.Fatal(e.Start(":8080"))
}
