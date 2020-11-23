package main

import(

	"fmt"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"encoding/json"
	"github.com/gorilla/mux"
)

var db *gorm.DB
var err error

type User struct{
	gorm.Model
	Name string
	Email string
}

func InitialMigration(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("Greska u spajanju sa bazom")
	}
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Greska u spajanju sa bazom")
	}

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func NewUser(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Greska u spajanju sa bazom")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "Novi korisnik stvoren")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Greska u spajanju sa bazom")
	}

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Korisnik uspjesno obrisan")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Greska u spajanju sa bazom")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user User

	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Korisnik uspjesno izmjenjen")
}