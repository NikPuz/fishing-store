package handler

import (
	"encoding/json"
	"errors"
	"fishing-store/internal/api/httpMiddleware"
	"fishing-store/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"time"
)

type manufacturerHandler struct {
	manufacturerService entity.IManufacturerService
}

func RegisterManufacturerHandlers(r *chi.Mux, service entity.IManufacturerService, routerMiddleware httpMiddleware.IMiddleware) {
	manufacturerHandler := new(manufacturerHandler)
	manufacturerHandler.manufacturerService = service
	r.Route("/manufacturers", func(r chi.Router) {
		r.Use(routerMiddleware.PanicRecovery)
		r.Use(middleware.Timeout(time.Second * 10))
		r.Use(middleware.RealIP)
		r.Use(middleware.RequestID)
		r.Use(routerMiddleware.ContentTypeJSON)

		r.Post("/", routerMiddleware.RequestLogger(manufacturerHandler.CreateManufacturer))
		r.Get("/{id}", routerMiddleware.RequestLogger(manufacturerHandler.ReadManufacturer))
		r.Put("/", routerMiddleware.RequestLogger(manufacturerHandler.UpdateManufacturer))
		r.Delete("/{id}", routerMiddleware.RequestLogger(manufacturerHandler.DeleteManufacturer))
		r.Get("/", routerMiddleware.RequestLogger(manufacturerHandler.ReadManufacturers))
	})

}

func (h manufacturerHandler) CreateManufacturer(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var manufacturer entity.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&manufacturer)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.manufacturerService.CreateManufacturer(r.Context(), &manufacturer)

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

func (h manufacturerHandler) ReadManufacturer(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	manufacturer, err := h.manufacturerService.ReadManufacturer(r.Context(), id)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(manufacturer)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}

func (h manufacturerHandler) UpdateManufacturer(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var manufacturer entity.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&manufacturer)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.manufacturerService.UpdateManufacturer(r.Context(), &manufacturer)

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

func (h manufacturerHandler) DeleteManufacturer(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.manufacturerService.DeleteManufacturer(r.Context(), id)

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

func (h manufacturerHandler) ReadManufacturers(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	manufacturer, err := h.manufacturerService.ReadManufacturers(r.Context())

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(manufacturer)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}
