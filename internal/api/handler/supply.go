package handler

import (
	"encoding/json"
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
		r.Get("/", routerMiddleware.RequestLogger(supplyHandler.ReadSupplies))
	})
}

func (h supplyHandler) CreateSupply(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var supply entity.Supply
	err := json.NewDecoder(r.Body).Decode(&supply)

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	respSupply, err := h.supplyService.CreateSupply(r.Context(), &supply)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(respSupply)

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
	return nil, http.StatusCreated, nil
}

func (h supplyHandler) ReadSupplies(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	supplies, err := h.supplyService.ReadSupplies(r.Context())

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(supplies)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}
