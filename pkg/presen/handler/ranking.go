package handler

import (
	"encoding/json"
	"net/http"

	customError "42tokyo-road-to-dojo-go/pkg/core/error"
	"42tokyo-road-to-dojo-go/pkg/core/logger"
	"42tokyo-road-to-dojo-go/pkg/presen/request"
	"42tokyo-road-to-dojo-go/pkg/presen/response"
	"42tokyo-road-to-dojo-go/pkg/usecase"
)

type RankingHandler interface {
	List(http.ResponseWriter, *http.Request)
}

type rankingHandler struct {
	rankingUsecase usecase.RankingUsecase
}

func NewRankingHandler(rankingUsecase usecase.RankingUsecase) RankingHandler {
	return &rankingHandler{rankingUsecase: rankingUsecase}
}

func (rh *rankingHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.ErrorLogging("GET ranking/list: method not ", customError.ErrMethodNotFound, r)
		return
	}

	var req request.RankingListRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		// TODO: パラメータの間違いなどのuserによるエラーだけが原因とは限らないので、
		// できればerrの内容でハンドリングする
		w.WriteHeader(http.StatusBadRequest)
		logger.ErrorLogging("GET ranking/list: decode error", err, r)
		return
	}

	userEntities, err := rh.rankingUsecase.List(r.Context(), req.Start)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("GET ranking/list: exec error", err, r)
		return
	}

	var ranks []response.User
	for _, ue := range userEntities {
		user := response.User{
			Id:        ue.Id,
			Name:      ue.Name,
			HighScore: ue.HighScore,
			Coin:      ue.Coin,
		}
		ranks = append(ranks, user)
	}

	res := response.RankingListResponse{
		Ranks: ranks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		logger.ErrorLogging("GET ranking/list: encode error", err, r)
		return
	}
}
