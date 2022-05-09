package handler

import (
	"encoding/json"
	"net/http"

	customError "42tokyo-road-to-dojo-go/pkg/core/error"
	"42tokyo-road-to-dojo-go/pkg/core/logger"
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"42tokyo-road-to-dojo-go/pkg/presen/request"
	"42tokyo-road-to-dojo-go/pkg/presen/response"
	"42tokyo-road-to-dojo-go/pkg/usecase"
)

type UserHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
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
		logger.ErrorLogging("POST user/create: decode error", customError.ErrMethodNotFound, r)
		return
	}

	var req request.UserCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		// TODO: パラメータの間違いなどのuserによるエラーだけが原因とは限らないので、
		// できればerrの内容でハンドリングする
		w.WriteHeader(http.StatusBadRequest)
		logger.ErrorLogging("POST user/create: decode error", err, r)
		return
	}

	createdUser, err := uh.userUsecase.Create(r.Context(), req.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("POST user/create: exec error", err, r)
		return
	}

	res := response.UserCreateResponse{
		Token: createdUser.Token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("POST user/create: encode error", err, r)
		return
	}
}

func (uh *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.ErrorLogging("GET user/get: method not allowed", customError.ErrMethodNotFound, r)
		return
	}

	token := r.Context().Value(middleware.Token).(string)
	targetUser, err := uh.userUsecase.Get(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("GET user/get: exec error", err, r)
		return
	}

	res := response.UserGetResponse{
		Name: targetUser.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("GET user/get: encode error", err, r)
		return
	}
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.ErrorLogging("POST user/update: method not allowed", customError.ErrMethodNotFound, r)
		return
	}

	var req request.UserUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.ErrorLogging("POST user/update: decode error", err, r)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	token := r.Context().Value(middleware.Token).(string)
	_, err = uh.userUsecase.Update(r.Context(), req.Name, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("POST user/update: exec error", err, r)
	}
}
