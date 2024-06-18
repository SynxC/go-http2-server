package store

import (
	"errors"
	"go-http2-server/models"
	"sync"
)

type PetStore struct {
	pets map[int64] * models.Pet
	mu sync.RWMutex
	id int64
}

func ConstructPetStore() *PetStore {
	return &PetStore{
		pets: make(map[int64] * models.Pet),
		id: 0,
	}
}

// used to access internal map of store object (for testing purposes)
func (store *PetStore) GetPetStorePetsMap() map[int64]*models.Pet {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.pets
}

func (store *PetStore) AddPet(pet *models.Pet) int64 {
	store.mu.Lock()
	defer store.mu.Unlock()
	pet.ID = store.id
	// assumption made that DateAdded is provided
	// pet.DateAdded = time.Now().UTC()
	store.pets[store.id] = pet
	store.id ++
	return pet.ID
}

func (store *PetStore) GetPet(petID int64) (*models.Pet, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	pet, exist := store.pets[petID]
	return pet, exist
}

func (store *PetStore) DeletePet(petID int64) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	_, exist := store.pets[petID]
	if !exist {
		return errors.New("pet not found")
	}

	delete(store.pets, petID)
	return nil
}
