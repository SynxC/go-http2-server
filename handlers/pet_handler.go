package handlers

import (
	"encoding/json"
	"go-http2-server/config"
	"go-http2-server/models"
	"go-http2-server/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func init() {
	validate.RegisterValidation("status", validateStatus)
	validate.RegisterValidation("category", validateCategory)
}

/* REST API controllers */
func PostPets(petStore *store.PetStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var pet models.Pet
		err := json.NewDecoder(r.Body).Decode(&pet)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		err = validate.Struct(&pet)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		petID := petStore.AddPet(&pet)
		petObj := petStore.GetPetStorePetsMap()[petID]
		w.Header().Set("Location", config.BaseURL+config.ApiURL+"/"+strconv.FormatInt(petID, 10))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(petObj)
	}
}

func GetOrDeletePets(petStore *store.PetStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodDelete {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

		strPetID := r.PathValue("id")
		if strPetID == "" {
			split := strings.Split(r.URL.Path, "/")
			strPetID = split[len(split)-1]
		}
		petID, err := strconv.ParseInt(strPetID, 10, 64)
		if err != nil{
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			pet, exist := petStore.GetPet(petID)
			if !exist {
				http.Error(w, "Pet not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pet)

		case http.MethodDelete:
			err := petStore.DeletePet(petID)
			if err != nil {
				http.Error(w, "Pet not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

/* Validation functions */
func validateStatus(fl validator.FieldLevel) bool {
	status := models.Status(fl.Field().String())
	switch status {
	case models.Available, models.Pending, models.Sold:
		return true
	}

	return false
}

func validateCategory(fl validator.FieldLevel) bool {
	category := models.Category(fl.Field().String())
	switch category {
	case models.Dog, models.Cat, models.Bird, models.Fish, models.Reptile, models.Other:
		return true
	}
	return false
}