package myclient

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/suryab-21/sigmatech-test/app/helper"
	"github.com/suryab-21/sigmatech-test/app/service"
)

func DeleteMyClient(w http.ResponseWriter, r *http.Request) {
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

	db.Delete(myClient)
	cache.Del(context.Background(), *myClient.Slug)

	helper.NewSuccessResponse(w, []byte("Successfully delete client"))
}
