package main

import (
	"database/sql"
	"fmt"
	"miniproject/controllers"
	"miniproject/database"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	} else {
		fmt.Println("success read file environtment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connected Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//Router GIN

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("persons", controllers.InsertPerson)
	router.PUT("persons/:id", controllers.UpdatePerson)
	router.DELETE("persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))

}
