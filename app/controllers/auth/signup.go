package auth

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suryab-21/golang-template/app/helper"
	"github.com/suryab-21/golang-template/app/model"
	"github.com/suryab-21/golang-template/app/service"
)

type SignUpBody struct {
	SignInBody
	Name string `json:"name"`
}

// @Summary      Sign Up
// @Description  Sign up new account
// @Tags         User Authentication & Role Management
// @Accept       application/json
// @Produce		 application/json
// @Param        data   body  auth.SignUpBody  true  "Sign Up Payload"
// @Router       /api/sign-up [post]
func SignUp(w http.ResponseWriter, r *http.Request) {
	var body SignUpBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Print(body)
	if body.Name == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}

	if body.Username == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "Username is required")
		return
	}

	if body.Password == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "Password is required")
		return
	}

	hash := sha256.New()
	hash.Write([]byte(body.Password))
	hashPassword := fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))

	db := service.DB

	newUser := model.User{}
	newUser.Name = &body.Name
	newUser.Username = &body.Username
	newUser.Password = &hashPassword
	db.Create(&newUser)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": *newUser.ID,
		"exp":     time.Now().Add(time.Hour * 6).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	response, _ := json.Marshal(map[string]interface{}{
		"token":   signedToken,
		"expired": 6 * 3600,
	})

	helper.NewSuccessResponse(w, response)
}
