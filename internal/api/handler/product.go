package handler

import (
	"encoding/json"
	"fishing-store/internal/api/httpMiddleware"
	"fishing-store/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"time"
)

type productHandler struct {
	productService entity.IProductService
}

func RegisterProductHandlers(r *chi.Mux, service entity.IProductService, routerMiddleware httpMiddleware.IMiddleware) {
	productHandler := new(productHandler)
	productHandler.productService = service
	r.Route("/products", func(r chi.Router) {
		r.Use(routerMiddleware.PanicRecovery)
		r.Use(middleware.Timeout(time.Second * 10))
		r.Use(middleware.RealIP)
		r.Use(middleware.RequestID)
		r.Use(routerMiddleware.ContentTypeJSON)

		r.Post("/", routerMiddleware.RequestLogger(productHandler.CreateProduct))
		r.Get("/{id}", routerMiddleware.RequestLogger(productHandler.ReadProduct))
		r.Put("/", routerMiddleware.RequestLogger(productHandler.UpdateProduct))
		r.Delete("/{id}", routerMiddleware.RequestLogger(productHandler.DeleteProduct))
		r.Get("/", routerMiddleware.RequestLogger(productHandler.ReadProducts))
	})
}

func (h productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	respProduct, err := h.productService.CreateProduct(r.Context(), &product)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(respProduct)

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
	return nil, http.StatusCreated, nil
}

func (h productHandler) ReadProduct(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	product, err := h.productService.ReadProduct(r.Context(), id)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(product)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}

func (h productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.productService.UpdateProduct(r.Context(), &product)

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

func (h productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(err, http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.productService.DeleteProduct(r.Context(), id)

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

func (h productHandler) ReadProducts(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	product, err := h.productService.ReadProducts(r.Context())

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(product)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}
