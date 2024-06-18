package tests

import (
	"go-http2-server/models"
	"go-http2-server/store"
	"testing"
	"time"
)

func TestPetStore_AddPet(t *testing.T) {
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

	got := mockPetStore.AddPet(pet)
	expected := int64(0)

	if got != expected {
		t.Errorf("PetStore.AddPet() = %v, expected %v", got, expected)
	}

	petsMap := mockPetStore.GetPetStorePetsMap()
	_, exist := petsMap[got]
	if !exist {
		t.Error("PetStore.AddPet() did not store the pet correctly")
	}
}

func TestPetStore_GetPet(t *testing.T) {
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
	mockPetStore.AddPet(pet)

	tests := []struct {
		name    string
		petID   int64
		expectedPet *models.Pet
		expectedExist bool
	}{
		{
			name:    "Get existing pet",
			petID:   0,
			expectedPet: pet,
			expectedExist: true,
		},
		{
			name:    "Get non-existing pet",
			petID:   10,
			expectedPet: nil,
			expectedExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPet, gotExist := mockPetStore.GetPet(tt.petID)
			if gotExist != tt.expectedExist {
				t.Errorf("PetStore.GetPet() exist = %v, expected %v", gotExist, tt.expectedExist)
			}
			if gotExist && gotPet != tt.expectedPet {
				t.Errorf("PetStore.GetPet() = %v, expected %v", gotPet, tt.expectedPet)
			}
		})
	}
}

func TestPetStore_DeletePet(t *testing.T) {
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
	mockPetStore.AddPet(pet)

	tests := []struct {
		name    string
		petID   int64
		expectedError bool
	}{
		{
			name:    "Delete existing pet",
			petID:   0,
			expectedError: false,
		},
		{
			name:    "Delete non-existing pet",
			petID:   10,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockPetStore.DeletePet(tt.petID)
			if (err != nil) != tt.expectedError {
				t.Errorf("PetStore.DeletePet() error = %v, expectedErr %v", err, tt.expectedError)
			}

			petsMap := mockPetStore.GetPetStorePetsMap()
			if !tt.expectedError {
				_, exist := petsMap[tt.petID]
				if exist {
					t.Error("PetStore.DeletePet() did not delete the pet correctly")
				}
			}
		})
	}
}