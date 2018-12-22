package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)


func Database() *gorm.DB{

	//open a db connection
	db_user   := os.Getenv("DB_USER")
	db_pass   := os.Getenv("DB_PASS")
	db_host   := os.Getenv("DB_HOST")
	db_port   := os.Getenv("DB_PORT")
	db, err := gorm.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/kubernetes")
	if err != nil{
		fmt.Println(err.Error())
	}
	return db
}



type customers struct {
	Id string `json:"id" binding:"required"`
	Nama string `json:"username" binding:"required"`
	Ktp int  `json:"ktp" binding:"required"`
	Status int  `json:"status" binding:"required"`
	Reg_date string  `json:"reg_date" binding:"required"`
	Alamat string  `json:"alamat" binding:"required"` 
}

func GetCustomer(c *gin.Context) {
	var customer []customers
	db := Database()
	db.Find(&customer)

	if len(customer) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Customer not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "count": len(customer), "result" : customer})
}


func main() {

	//router := gin.Default()
        router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/customer", GetCustomer)
	router.Run(":8082")
}
