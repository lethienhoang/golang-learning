package databases

import (
	model "building-restapi-orm-with-gorm/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbContext struct {
	Db *gorm.DB
}

func (conn *DbContext) Initialize() {

	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	dbname := os.Getenv("APP_DB_NAME")

	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	conn.Db, err = gorm.Open("postgres", connString)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	defer conn.Db.Close()

	fmt.Println("Successfully connected!")

	conn.Db.AutoMigrate(&model.User{})
}
