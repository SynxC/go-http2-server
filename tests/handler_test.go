package tests

import (
	"bytes"
	"encoding/json"
	"go-http2-server/config"
	"go-http2-server/handlers"
	"go-http2-server/models"
	"go-http2-server/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPostPets(t *testing.T) {
	mockPetStore := store.ConstructPetStore()

	pet := &models.Pet{
		Name:     "Doggie",
		Status:   "available",
		Category: "dog",
		Tags:     []string{"friendly", "energetic"},
		Age:      3,
		Color:    "Golden",
		Gender:   "male",
		Weight:   25.5,
		DateAdded: time.Now().UTC(),
	}

	t.Run("Post a valid pet", func(t *testing.T) {
		body, _ := json.Marshal(pet)
		req, err := http.NewRequest("POST", config.ApiURL, bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		res := httptest.NewRecorder()
		handler := handlers.PostPets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusCreated {
			t.Errorf("PostPets returned wrong status code: got %v expected %v", status, http.StatusCreated)
		}

		expected := `{"id":0}`
		if strings.TrimSpace(res.Body.String()) != strings.TrimSpace(expected) {
			t.Errorf("PostPets returned unexpected body: got %v expected %v", res.Body.String(), expected)
		}

		location := res.Header().Get("Location")
		expectedLocation := "http://localhost:8080/api/v1/pets/0"
		if location != expectedLocation {
			t.Errorf("PostPets returned unexpected Location header: got %v expected %v", location, expectedLocation)
		}
	})

	t.Run("Post an invalid pet", func(t *testing.T) {
		pet = &models.Pet{}
		body, _ := json.Marshal(pet)
		req, err := http.NewRequest("POST", config.ApiURL, bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}

		res := httptest.NewRecorder()
		handler := handlers.PostPets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusBadRequest{
			t.Errorf("PostPets returned wrong status code: got %v expected %v", status, http.StatusBadRequest)
		}
	})
}

func TestGetOrDeletePets(t *testing.T) {
	mockPetStore := store.ConstructPetStore()

	pet := &models.Pet{
		Name:     "Doggie",
		Status:   "available",
		Category: "dog",
		Tags:     []string{"friendly", "energetic"},
		Age:      3,
		Color:    "Golden",
		Gender:   "male",
		Weight:   25.5,
	}

	petID := mockPetStore.AddPet(pet)
	petsMap := mockPetStore.GetPetStorePetsMap()
	dateAdded, _ := time.Parse(time.RFC3339, "2023-01-20T08:30:00Z")
	petsMap[petID].DateAdded = dateAdded

	t.Run("Get pet with valid ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.ApiURL+"/"+strconv.FormatInt(petID, 10), nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		res := httptest.NewRecorder()
		handler := handlers.GetOrDeletePets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusOK {
			t.Errorf("GetOrDeletePets GET valid pet returned wrong status code: got %v expected %v", status, http.StatusOK)
		}

		expected := `{"id":0,"name":"Doggie","status":"available","category":"dog","tags":["friendly","energetic"],"age":3,"color":"Golden","gender":"male","weight":25.5,"dateAdded":"2023-01-20T08:30:00Z"}`
		if strings.TrimSpace(res.Body.String()) != expected {
			t.Errorf("GetOrDeletePets GET valid pet returned unexpected body: got %v expected %v", res.Body.String(), expected)
		}
	})

	t.Run("Get pet with invalid petID", func(t *testing.T) {
		req, err := http.NewRequest("GET", config.ApiURL+"/99", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		res := httptest.NewRecorder()
		handler := handlers.GetOrDeletePets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusNotFound{
			t.Errorf("GetOrDeletePets GET invalid pet returned wrong status code: got %v expected %v", status, http.StatusNotFound)
		}
	})

	t.Run("Delete pet with valid petID", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", config.ApiURL+"/"+strconv.FormatInt(petID, 10), nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		res := httptest.NewRecorder()
		handler := handlers.GetOrDeletePets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusNoContent{
			t.Errorf("GetOrDeletePets DELETE valid pet returned wrong status code: got %v expected %v", status, http.StatusNoContent)
		}
	})

	t.Run("Delete pet with invalid petID", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", config.ApiURL+"/99", nil)
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		res := httptest.NewRecorder()
		handler := handlers.GetOrDeletePets(mockPetStore)
		handler.ServeHTTP(res, req)

		status := res.Code
		if status != http.StatusNotFound{
			t.Errorf("GetOrDeletePets DELETE invalid pet returned wrong status code: got %v expected %v", status, http.StatusNotFound)
		}
	})
}