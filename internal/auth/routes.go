package auth

import (
	"net/http"

	"github.com/b1g-Dr4goN/hussatapi/internal/user"
	"github.com/gorilla/mux"
)

type Handler struct {
	userStore *user.Store
}

func NewHandler(userStore *user.Store) *Handler {
	return &Handler{userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/signin-with-google", h.handleSignInWithGoogle).Methods(http.MethodPost)
}
