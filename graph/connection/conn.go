package connection

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Video struct {
	gorm.Model
	Name   string
	Url    string
	Author string
}

// const (
// 	Host     = "localhost"
// 	Port     = 5432
// 	UserName = "postgres"
// 	Password = "ghosh"
// 	Dbname   = "firstdb"
// )

func Conn() {
	DB, err = gorm.Open(postgres.Open("user=postgres password=ghosh dbname=firstdb sslmode=disable"), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Cannot connect to DB")
	}

	err := DB.AutoMigrate(&Video{})
	log.Println(err, ";;;;;;;;;;;;;;;;;;;;;;;;;;;")
	fmt.Println("successful conn")
	//sqlstatement:=`I`

}
