package myclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/suryab-21/sigmatech-test/app/helper"
	"github.com/suryab-21/sigmatech-test/app/model"
	"github.com/suryab-21/sigmatech-test/app/service"
	"gorm.io/gorm"
)

func GetMyClient(w http.ResponseWriter, r *http.Request) {
	db := service.DB

	myClient := []model.MyClient{}

	query := db.Model(&model.MyClient{}).Find(&myClient)

	if query.RowsAffected < 1 {
		helper.NewErrorResponse(w, 404, "data not found")
		return
	}

	json, err := json.Marshal(myClient)
	if err != nil {
		fmt.Println(err)
	}

	helper.NewSuccessResponse(w, json)
}

func GetByIdMyClient(w http.ResponseWriter, r *http.Request) {
	db := service.DB

	id := r.PathValue("id")

	if err := uuid.Validate(id); err != nil {
		helper.NewErrorResponse(w, 400, "ID is not valid")
		return
	}

	myClient, err := FindMyClientById(id, db)
	if err != nil {
		helper.NewErrorResponse(w, 404, err.Error())
		return
	}

	json, err := json.Marshal(myClient)
	if err != nil {
		fmt.Println(err)
	}

	helper.NewSuccessResponse(w, json)
}

func FindMyClientById(id string, db *gorm.DB) (*model.MyClient, error) {
	myClient := model.MyClient{}

	query := db.Model(model.MyClient{}).First(&myClient, `id = ?`, id)
	if query.RowsAffected < 1 {
		return nil, errors.New("data not found")
	}

	return &myClient, nil
}
