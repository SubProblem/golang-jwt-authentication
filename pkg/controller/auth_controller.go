package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"subproblem/rest-api/pkg/dto"
	"subproblem/rest-api/pkg/service"
	"subproblem/rest-api/pkg/jwt"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(authervice *service.AuthService) *AuthController {
	return &AuthController{
		service: authervice,
	}
}

func(auth  *AuthController) Register(w http.ResponseWriter, r *http.Request) {

	var userReqeust dto.UserRequestDto

	err := json.NewDecoder(r.Body).Decode(&userReqeust)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err2 := auth.service.Register(userReqeust)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


func(auth *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest dto.LoginRequestDto

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err2 := auth.service.Login(loginRequest)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
	}

	response := map[string]string{"token":token}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}


func(auth *AuthController) Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			http.Error(w, "Incorrect token", http.StatusBadRequest)
            return
		}

		jwt, err := jwt.ValidateToken(token[1])
		if err != nil {
			http.Error(w, "Incorrect token", http.StatusUnauthorized)
            return
		}

		if !jwt.Valid {
			http.Error(w, "Incorrect token", http.StatusUnauthorized)
            return
		}

		next.ServeHTTP(w,r)
	})
}