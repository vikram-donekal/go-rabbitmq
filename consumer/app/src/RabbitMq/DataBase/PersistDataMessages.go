package DataBase


import (
	"../../../bin/gorm"
	"fmt"
	"errors"
	"log"
	"../DTO"
	_ "../../../bin/pq"
)

var DatabaseConnection *gorm.DB //database



const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)



func  init ()   {
	log.Print("connecting to Data Base Postgresql")
	err:=errors.New("")
	DatabaseInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	DatabaseConnection,err =gorm.Open("postgres",DatabaseInfo)
	if err!= nil {
		panic(err.Error())
	}
	//defer  	DatabaseConnection.Close()
	DataBase:=DatabaseConnection.DB()
	//defer  DataBase.Close()
	err =DataBase.Ping()

	if err!= nil {
		panic(err.Error())
	}
	log.Println("creating tables using GORM ")
	DatabaseConnection.DropTableIfExists(&DTO.MessageDTO{})
	DatabaseConnection.CreateTable(&DTO.MessageDTO{})
}

func GetDB() *gorm.DB {
	return DatabaseConnection
}

func SaveMessages( dto *DTO.MessageDTO)  {

	GetDB().Debug().Save(&dto)



}