package main

import (
	"abc/database"
	"abc/handler"
	"abc/repository"
	"abc/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()

	r := repository.NewRorepositoryAdapter(db)
	s := service.NewserviceAdapter(r)
	h := handler.NewhandlerAdapter(s)

	router := gin.Default()

	router.GET("/api/getBeers", h.GetHand)
	err := router.Run(":9000")
	if err != nil {
		panic(err.Error())

	}
	//* fmt.Println(r.Get())
}
