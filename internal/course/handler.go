package course

import (
	"awesomeProject/internal/apperror"
	"awesomeProject/internal/handlers"
	"awesomeProject/pkg/logging"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	coursesURL = "/courses"
	courseURL  = "/courses/:uuid"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, coursesURL, apperror.Middleware(h.GetList))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}
	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
	return nil
}
