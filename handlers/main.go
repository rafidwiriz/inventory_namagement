package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Container struct {
	ID          uint   `gorm:"primaryKey;not null"`
	Name        string `json:"name" gorm:"not null"`
	MaxCapacity int    `json:"max_capacity" gorm:"not null"`
}

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	ContainerID uint   `json:"container_id" gorm:"not null"`
	Container   Container
}

var (
	db  *gorm.DB
	err error
)

func DBConnection() error {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Container{}, &Item{})
	if err != nil {
		return err
	}
	return nil
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	_ = db.Find(&items)
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		log.Println(err)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
	}
	db.Create(&item)
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
	}
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item Item
	db.First(&item, id)
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item Item
	db.First(&item, id)
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
	}
	db.Save(&item)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
	}
}

func GetAllContainers(w http.ResponseWriter, r *http.Request) {
	var containers []Container
	_ = db.Find(&containers)
	err := json.NewEncoder(w).Encode(containers)
	if err != nil {
		log.Println(err)
	}
}

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	var container Container
	err = json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		log.Println(err)
	}
	db.Create(&container)
	err := json.NewEncoder(w).Encode(container)
	if err != nil {
		log.Println(err)
	}
}

func GetContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r.Response.Request)
	id := vars["id"]
	var container Container
	db.First(&container, id)
	err := json.NewEncoder(w).Encode(container)
	if err != nil {
		log.Println(err)
	}
}

func isContainerFull(containerID uint) bool {
	var container Container
	var items []Item
	db.First(&container, containerID)
	result := db.Find(&items, "container_id = ?", containerID)

	isFull := result.RowsAffected >= int64(container.MaxCapacity)
	return isFull
}
