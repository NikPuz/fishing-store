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
		r.Put("/", routerMiddleware.RequestLogger(saleHandler.UpdateSale))
		r.Get("/", routerMiddleware.RequestLogger(saleHandler.ReadSales))
	})
}

func (h saleHandler) CreateSale(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var sales entity.SaleDTO
	err := json.NewDecoder(r.Body).Decode(&sales)

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
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

func (h saleHandler) UpdateSale(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var sales entity.SaleDTO
	err := json.NewDecoder(r.Body).Decode(&sales)

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.saleService.UpdateSale(r.Context(), &sales)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	w.WriteHeader(http.StatusOK)
	return nil, http.StatusOK, nil
}

func (h saleHandler) ReadSales(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	sales, err := h.saleService.ReadSales(r.Context())

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(sales)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}
