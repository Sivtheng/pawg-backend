package models

import (
	"database/sql"
	"time"
)

type AdoptionApplication struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	PhoneNumber        string    `json:"phone_number"`
	Address            string    `json:"address"`
	InterestInAdopting string    `json:"interest_in_adopting"`
	TypeOfAnimal       string    `json:"type_of_animal"`
	SpecialNeedsAnimal string    `json:"special_needs_animal"`
	OwnPetBefore       string    `json:"own_pet_before"`
	WorkingTime        string    `json:"working_time"`
	LivingSituation    string    `json:"living_situation"`
	OtherAnimals       string    `json:"other_animals"`
	AnimalAccess       string    `json:"animal_access"`
	Travel             string    `json:"travel"`
	LeaveCambodia      string    `json:"leave_cambodia"`
	Feed               string    `json:"feed"`
	AnythingElse       string    `json:"anything_else"`
	CreatedAt          time.Time `json:"created_at"`
}

// ListAdoptionApplications retrieves all adoption applications from the database
func ListAdoptionApplications(db *sql.DB) ([]AdoptionApplication, error) {
	rows, err := db.Query(`SELECT id, name, email, phone_number, address, interest_in_adopting, type_of_animal, special_needs_animal, own_pet_before, working_time, living_situation, other_animals, animal_access, travel, leave_cambodia, feed, anything_else, created_at FROM adoption_applications`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applicationList []AdoptionApplication
	for rows.Next() {
		var application AdoptionApplication
		if err := rows.Scan(&application.ID, &application.Name, &application.Email, &application.PhoneNumber, &application.Address, &application.InterestInAdopting, &application.TypeOfAnimal, &application.SpecialNeedsAnimal, &application.OwnPetBefore, &application.WorkingTime, &application.LivingSituation, &application.OtherAnimals, &application.AnimalAccess, &application.Travel, &application.LeaveCambodia, &application.Feed, &application.AnythingElse, &application.CreatedAt); err != nil {
			return nil, err
		}
		applicationList = append(applicationList, application)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return applicationList, nil
}

// Insert new application into database and return the created application with its ID and creation timestamo
func CreateAdoptionApplication(db *sql.DB, name, email, phoneNumber, address, interestInAdopting, typeOfAnimal, specialNeedsAnimal, ownPetBefore, workingTime, livingSituation, otherAnimals, animalAccess, travel, leaveCambodia, feed, anythingElse string) (*AdoptionApplication, error) {
	var application AdoptionApplication
	err := db.QueryRow(
		`INSERT INTO adoption_applications (name, email, phone_number, address, interest_in_adopting, type_of_animal, special_needs_animal, own_pet_before, working_time, living_situation, other_animals, animal_access, travel, leave_cambodia, feed, anything_else) 
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id, created_at`,
		name, email, phoneNumber, address, interestInAdopting, typeOfAnimal, specialNeedsAnimal, ownPetBefore, workingTime, livingSituation, otherAnimals, animalAccess, travel, leaveCambodia, feed, anythingElse,
	).Scan(&application.ID, &application.CreatedAt)
	if err != nil {
		return nil, err
	}
	application.Name = name
	application.Email = email
	application.PhoneNumber = phoneNumber
	application.Address = address
	application.InterestInAdopting = interestInAdopting
	application.TypeOfAnimal = typeOfAnimal
	application.SpecialNeedsAnimal = specialNeedsAnimal
	application.OwnPetBefore = ownPetBefore
	application.WorkingTime = workingTime
	application.LivingSituation = livingSituation
	application.OtherAnimals = otherAnimals
	application.AnimalAccess = animalAccess
	application.Travel = travel
	application.LeaveCambodia = leaveCambodia
	application.Feed = feed
	application.AnythingElse = anythingElse
	return &application, nil
}

// Retrieves application from database by ID and return the details
func GetAdoptionApplicationByID(db *sql.DB, id int) (*AdoptionApplication, error) {
	var application AdoptionApplication
	err := db.QueryRow(
		`SELECT id, name, email, phone_number, address, interest_in_adopting, type_of_animal, special_needs_animal, own_pet_before, working_time, living_situation, other_animals, animal_access, travel, leave_cambodia, feed, anything_else, created_at 
         FROM adoption_applications WHERE id = $1`,
		id,
	).Scan(&application.ID, &application.Name, &application.Email, &application.PhoneNumber, &application.Address, &application.InterestInAdopting, &application.TypeOfAnimal, &application.SpecialNeedsAnimal, &application.OwnPetBefore, &application.WorkingTime, &application.LivingSituation, &application.OtherAnimals, &application.AnimalAccess, &application.Travel, &application.LeaveCambodia, &application.Feed, &application.AnythingElse, &application.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &application, nil
}

// Update existing application in database with new info and return the updated application
func UpdateAdoptionApplication(db *sql.DB, id int, name, email, phoneNumber, address, interestInAdopting, typeOfAnimal, specialNeedsAnimal, ownPetBefore, workingTime, livingSituation, otherAnimals, animalAccess, travel, leaveCambodia, feed, anythingElse string) (*AdoptionApplication, error) {
	var application AdoptionApplication
	err := db.QueryRow(
		`UPDATE adoption_applications SET name = $1, email = $2, phone_number = $3, address = $4, interest_in_adopting = $5, type_of_animal = $6, special_needs_animal = $7, own_pet_before = $8, working_time = $9, living_situation = $10, other_animals = $11, animal_access = $12, travel = $13, leave_cambodia = $14, feed = $15, anything_else = $16 WHERE id = $17 RETURNING id, created_at`,
		name, email, phoneNumber, address, interestInAdopting, typeOfAnimal, specialNeedsAnimal, ownPetBefore, workingTime, livingSituation, otherAnimals, animalAccess, travel, leaveCambodia, feed, anythingElse, id,
	).Scan(&application.ID, &application.CreatedAt)
	if err != nil {
		return nil, err
	}
	application.Name = name
	application.Email = email
	application.PhoneNumber = phoneNumber
	application.Address = address
	application.InterestInAdopting = interestInAdopting
	application.TypeOfAnimal = typeOfAnimal
	application.SpecialNeedsAnimal = specialNeedsAnimal
	application.OwnPetBefore = ownPetBefore
	application.WorkingTime = workingTime
	application.LivingSituation = livingSituation
	application.OtherAnimals = otherAnimals
	application.AnimalAccess = animalAccess
	application.Travel = travel
	application.LeaveCambodia = leaveCambodia
	application.Feed = feed
	application.AnythingElse = anythingElse
	return &application, nil
}

// Remove application from database based on ID
func DeleteAdoptionApplication(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM adoption_applications WHERE id = $1`, id)
	return err
}
