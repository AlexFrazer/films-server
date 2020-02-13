package movies

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/AlexFrazer/films-server/app/util"
)

func GetAll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	movies := []Movie{}
	db.Find(&movies)
	util.RespondJSON(w, http.StatusOK, movies)
}

func Create(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := util.DecodeJSONBody(w, r, &movie)
	if err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Save(&movie).Error; err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSON(w, http.StatusCreated, movie)
}
