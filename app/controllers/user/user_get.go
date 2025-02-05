package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suryab-21/golang-template/app/helper"
	"github.com/suryab-21/golang-template/app/model"
	"github.com/suryab-21/golang-template/app/service"
)

// @Summary      Users Info
// @Description  Get user info
// @Tags         User Authentication & Role Management
// @Accept       application/json
// @Produce		 application/json
// @Router       /api/users/me [get]
// @Security BearerAuth
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		helper.NewErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(headerParts[1], claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("KEY")), nil
	})

	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	db := service.DB
	user := model.User{}
	db.First(&user, "id = ?", claims["user_id"])

	response, _ := json.Marshal(map[string]interface{}{
		"status": "success",
		"data":   user,
	})

	helper.NewSuccessResponse(w, response)
}
