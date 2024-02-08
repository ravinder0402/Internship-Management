package main

//importing packages
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// decalring database
type Data struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	ENAME        string `json:"ename" binding:"required"`
	EAGE         int    `json:"eage" binding:"required"`
	EADDRESS     string `json:"eaddress" binding:"required"`
	EPOST        string `json:"epost" binding:"required"`
	EUNINVERSITY string `json:"euniversity" binding:"required"`
	ESTIPEND     string `json:"estipend" binding:"required"`
	EDATE        string `json:"edate" binding:"required"`
	ESTATUS      string `json:"estatus" binding:"required"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("Database Connection Failure" + err.Error())
	}
	db.AutoMigrate(&Data{})

}

func createData(c *gin.Context) {
	var data Data
	if err := c.BindJSON((&data)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	db.Create(&data)
	c.JSON(http.StatusCreated, data)
}

func updateData(c *gin.Context) {
	var pdata Data
	var id = c.Param("id")
	if err := db.First(&pdata, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": "ID does not exist"})
		return

	}
	var ndata Data
	if err := c.BindJSON(&ndata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error()})
		return
	}
	db.Model(&pdata).Updates(ndata)
	c.JSON(http.StatusOK, pdata)
}
func getData(c *gin.Context) {
	var data Data
	var id = c.Param("id")
	if err := db.First(&data, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": "ID Does not Exist"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func deleteData(c *gin.Context) {
	var data Data
	var id = c.Param("id")
	if err := db.First(&data, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": "Id does not exist"})
		return
	}
	db.Delete(&data)
	c.JSON(http.StatusOK, gin.H{
		"message:": "Deletion is successful...."})
}
func wholeData(c *gin.Context) {
	var data []Data
	db.Find(&data)
	c.JSON(http.StatusOK, data)
}

func main() {
	initDB()
	r := gin.Default()
	r.POST("/data", createData)
	r.GET("/data/:id", getData)
	r.PUT("/data/:id", updateData)
	r.DELETE("/data/:id", deleteData)
	r.GET("/data", wholeData)
	r.Run(":6666")
}
