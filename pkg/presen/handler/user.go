package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"go.uber.org/zap"

	customError "42tokyo-road-to-dojo-go/pkg/core/error"
	"42tokyo-road-to-dojo-go/pkg/presen/request"
	"42tokyo-road-to-dojo-go/pkg/presen/response"
	"42tokyo-road-to-dojo-go/pkg/usecase"
)

type UserHandler interface {
	Create(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		zap.Error(customError.ErrMethodNotFound)
		return
	}

	var req request.UserCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		zap.Error(err)
		log.Println("decode", err.Error())
		// errが発生している際の処理が必要
		return
	}

	createdUser, err := uh.userUsecase.Create(r.Context(), req.Name)
	if err != nil {
		zap.Error(err)
		log.Println("createdUser", err.Error())
		return
	}

	res := response.UserCreateResponse{
		Token: createdUser.Token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		zap.Error(err)
		return
	}
}
