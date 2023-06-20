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

type categoryHandler struct {
	categoryService entity.ICategoryService
}

func RegisterCategoryHandlers(r *chi.Mux, service entity.ICategoryService, routerMiddleware httpMiddleware.IMiddleware) {
	categoryHandler := new(categoryHandler)
	categoryHandler.categoryService = service

	r.Route("/categories", func(r chi.Router) {
		r.Use(routerMiddleware.PanicRecovery)
		r.Use(middleware.Timeout(time.Second * 10))
		r.Use(middleware.RealIP)
		r.Use(middleware.RequestID)
		r.Use(routerMiddleware.ContentTypeJSON)

		r.Post("/", routerMiddleware.RequestLogger(categoryHandler.CreateCategory))
		r.Get("/{id}", routerMiddleware.RequestLogger(categoryHandler.ReadCategory))
		r.Put("/", routerMiddleware.RequestLogger(categoryHandler.UpdateCategory))
		r.Delete("/{id}", routerMiddleware.RequestLogger(categoryHandler.DeleteCategory))
		r.Get("/", routerMiddleware.RequestLogger(categoryHandler.ReadCategories))
	})
}

func (h categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.categoryService.CreateCategory(r.Context(), &category)

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

func (h categoryHandler) ReadCategory(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	category, err := h.categoryService.ReadCategory(r.Context(), id)

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(category)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}

func (h categoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.categoryService.UpdateCategory(r.Context(), &category)

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

func (h categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		logicError := entity.NewLogicError(errors.New("входные данные не распознаны"), http.StatusBadRequest)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	err = h.categoryService.DeleteCategory(r.Context(), id)

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

func (h categoryHandler) ReadCategories(w http.ResponseWriter, r *http.Request) ([]byte, int, *entity.LogicError) {

	category, err := h.categoryService.ReadCategories(r.Context())

	// Обработка ошибки
	if err != nil {
		logicError := entity.ResponseLogicError(err)
		resp := logicError.JsonMarshal()
		w.WriteHeader(logicError.Code)
		w.Write(resp)
		return resp, logicError.Code, logicError
	}

	resp, _ := json.Marshal(category)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return resp, http.StatusOK, nil
}
