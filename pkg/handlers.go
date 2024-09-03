package pkg

import (
	"backend/db"
	"encoding/json"
	"net/http"
)

// Pet represents a pet structure
type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetPets handles the "GET /pets" route
func GetPets(w http.ResponseWriter, r *http.Request) {
	pets := []Pet{}

	rows, err := db.DB.Query(r.Context(), "SELECT id, name, age FROM pets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pet Pet
		err = rows.Scan(&pet.ID, &pet.Name, &pet.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pets = append(pets, pet)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

// Other CRUD functions (CreatePet, UpdatePet, DeletePet) can be added here...
