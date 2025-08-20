package user

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// )

// func TestUserServiceHandler(t *testing.T) {
// 	userStore := &mockUserStore{}
// 	handler := NewHandler(userStore)

// 	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
// 		payload := RegisterUserPayload{
// 			Name:       "Tran Minh",
// 			GivenName:  "Tran",
// 			FamilyName: "Minh",
// 			Email:      "kuanmin.bigdragon.56",
// 			Role:       "root",
// 		}
// 		marshalled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusBadRequest {
// 			t.Errorf("expect status code %d, got %d", http.StatusBadRequest, rr.Code)
// 		}
// 	})

// 	t.Run("should correctly register the user", func(t *testing.T) {
// 		payload := RegisterUserPayload{
// 			Name:       "Tran Minh",
// 			GivenName:  "Tran",
// 			FamilyName: "Minh",
// 			Email:      "kuanmin.bigdragon.56@gmail.com",
// 			Role:       "root",
// 		}
// 		marshalled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusCreated {
// 			t.Errorf("expect status code %d, got %d", http.StatusCreated, rr.Code)
// 		}
// 	})
// }

// type mockUserStore struct {
// }

// func (m *mockUserStore) GetUserByEmail(email string) (*User, error) {
// 	return nil, fmt.Errorf("user not found")
// }

// func (m *mockUserStore) GetUserById(id string) (*User, error) {
// 	return nil, nil
// }

// func (m *mockUserStore) CreateUser(RegisterUserPayload) error {
// 	return nil
// }
