package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	customError "42tokyo-road-to-dojo-go/pkg/core/error"
	"42tokyo-road-to-dojo-go/pkg/core/logger"
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"42tokyo-road-to-dojo-go/pkg/presen/response"
	"42tokyo-road-to-dojo-go/pkg/usecase"
)

type CollectionHandler interface {
	List(http.ResponseWriter, *http.Request)
}

type collectionHandler struct {
	collectionUsecase usecase.CollectionUsecase
}

func NewCollectionHandler(collectionUsecase usecase.CollectionUsecase) CollectionHandler {
	return &collectionHandler{collectionUsecase: collectionUsecase}
}

func (ch *collectionHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.ErrorLogging("GET collection/list: method not ", customError.ErrMethodNotFound, r)
		return
	}

	token := r.Context().Value(middleware.Token).(string)
	ucEntities, err := ch.collectionUsecase.List(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.ErrorLogging("GET collection/list: exec error", err, r)
		return
	}

	var collections []response.CollectionItem
	for _, uce := range ucEntities {
		collection := response.CollectionItem{
			CollectionID: strconv.Itoa(uce.Id),
			Name:         uce.Chara.Name,
			Rarity:       uce.Chara.Rarity,
			HasItem:      true,
		}
		collections = append(collections, collection)
	}

	res := response.CollectionListResponse{
		Collections: collections,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		logger.ErrorLogging("GET collection/list: encode error", err, r)
		return
	}
}
