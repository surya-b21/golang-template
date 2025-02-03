package myclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/suryab-21/sigmatech-test/app/helper"
	"github.com/suryab-21/sigmatech-test/app/model"
	"github.com/suryab-21/sigmatech-test/app/service"
)

func PostMyClient(w http.ResponseWriter, r *http.Request) {
	var body model.MyClientApi
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db := service.DB
	cache := service.Cache

	newClient := model.MyClient{}
	newClient.MyClientApi = body
	db.Create(newClient)

	json, err := json.Marshal(newClient)
	if err != nil {
		fmt.Println(err)
	}

	cache.Set(context.Background(), *newClient.Slug, json, 0)

	helper.NewSuccessResponse(w, []byte("Successfully add new client"))
}
