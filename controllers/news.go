package controllers

import (
	"fmt"
	"net/http"

	"github.com/ConstantineGochev/go_api/models"
	mgo "gopkg.in/mgo.v2"
)

func Test(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blabla\n"))
}
func GetAllNews(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	news := []models.News{}
	fmt.Print(db)
	//db.Find(&news)
	respondJSON(w, http.StatusOK, news)
}

func CreateNewNews(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	fmt.Println("not implemented")
}
