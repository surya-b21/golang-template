package myclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/suryab-21/sigmatech-test/app/helper"
	"github.com/suryab-21/sigmatech-test/app/model"
	"github.com/suryab-21/sigmatech-test/app/service"
)

func PutMyClient(w http.ResponseWriter, r *http.Request) {
	var body model.MyClientApi
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db := service.DB
	cache := service.Cache

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

	myClient.MyClientApi = body
	db.Updates(&myClient)

	json, err := json.Marshal(myClient)
	if err != nil {
		fmt.Println(err)
	}

	cache.Del(context.Background(), *myClient.Slug)
	cache.Set(context.Background(), *myClient.Slug, json, 0)

	helper.NewSuccessResponse(w, []byte("Successfully update client"))
}
