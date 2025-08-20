package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/b1g-Dr4goN/hussatapi/configs/env"
	"github.com/b1g-Dr4goN/hussatapi/internal/user"
	"github.com/b1g-Dr4goN/hussatapi/utils"
	"github.com/go-playground/validator"
)

func (h *Handler) handleRegister(payload user.RegisterUserPayload) error {
	// Validate the Payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return fmt.Errorf("invalid payload %v", errors)
	}

	// Check if the User exists
	_, err := h.userStore.GetUserByEmail(payload.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", payload.Email)
	}

	return h.userStore.CreateUser(payload)
}

func (h *Handler) handleSignInWithGoogle(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AccessToken string `json:"accessToken"`
	}
	if err := utils.ParseJSON(r, &body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req, err := http.NewRequest(http.MethodGet, env.Envs.GoogleOAuthURL, nil)
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+body.AccessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var usr GoogleUser
	if err := json.NewDecoder(res.Body).Decode(&usr); err != nil {
		http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
		return
	}

	u, err := h.userStore.GetUserByEmail(usr.Email)
	if err != nil {
		userPayload := user.RegisterUserPayload{
			Name:       usr.Name,
			FamilyName: usr.FamilyName,
			GivenName:  usr.GivenName,
			Email:      usr.Email,
			Avatar:     usr.Picture,
			Role:       "user",
		}
		if err := h.handleRegister(userPayload); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		u, err = h.userStore.GetUserByEmail(usr.Email)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	secret := []byte(env.Envs.JWTSecret)
	token, err := CreateJWT(secret, u.UID, u.Role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		env.Envs.JWTTokenName: token,
		"expiresTime":         time.Now().Add(time.Second * time.Duration(env.Envs.JWTExpirationInSeconds)),
	})
}
