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

type saleHandler struct {
	saleService entity.ISaleService
}

func RegisterSaleHandlers(r *chi.Mux, service entity.ISaleService, routerMiddleware httpMiddleware.IMiddleware) {
	saleHandler := new(saleHandler)
	saleHandler.saleService = service
	r.Route("/sales", func(r chi.Router) {
		r.Use(routerMiddleware.PanicRecovery)
		r.Use(middleware.Timeout(time.Second * 10))
		r.Use(middleware.RealIP)
		r.Use(middleware.RequestID)
		r.Use(routerMiddleware.ContentTypeJSON)

		r.Post("/", routerMiddleware.RequestLogger(saleHandler.CreateSale))
	})
}

func (h saleHandler) CreateSale(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var sales entity.SaleDTO
	err := json.NewDecoder(r.Body).Decode(&sales)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.saleService.CreateSale(r.Context(), &sales)

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
