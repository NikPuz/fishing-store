package handler

import (
	"encoding/json"
	"errors"
	"fishing-store/internal/api/httpMiddleware"
	"fishing-store/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type supplyHandler struct {
	supplyService entity.ISupplyService
}

func RegisterSupplyHandlers(r *chi.Mux, service entity.ISupplyService, routerMiddleware httpMiddleware.IMiddleware) {
	supplyHandler := new(supplyHandler)
	supplyHandler.supplyService = service
	r.Route("/supplies", func(r chi.Router) {
		r.Use(routerMiddleware.PanicRecovery)
		r.Use(middleware.Timeout(time.Second * 10))
		r.Use(middleware.RealIP)
		r.Use(middleware.RequestID)
		r.Use(routerMiddleware.ContentTypeJSON)

		r.Post("/", routerMiddleware.RequestLogger(supplyHandler.CreateSupply))
	})
}

func (h supplyHandler) CreateSupply(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var supplies []entity.Supply
	err := json.NewDecoder(r.Body).Decode(&supplies)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.supplyService.CreateSupplies(r.Context(), supplies)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	w.WriteHeader(http.StatusCreated)
	return nil, http.StatusCreated, nil
}
