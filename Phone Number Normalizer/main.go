package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type contact struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	name   string
	number string `gorm:"unique"`
}

type datastore interface {
	create(name, number string)
	find(name string) []contact
	updateName(key uint, name string)
	updateNumber(key uint, number string)
	delete(key uint)
}

// Mydb provides methods on my sqlite database
type Mydb struct {
	db gorm.DB
}

func (d Mydb) create(name, number string) {

}

func (d Mydb) find(name string) []contact {
	var contacts []contact

	return contacts
}

func (d Mydb) updateName(key uint, name string) {

}

func (d Mydb) updateNumber(key uint, number string) {

}

func (d Mydb) delete(key uint) {

}

// NewDb returns an instance of sqlite database
func NewDb() (Mydb, error) {
	var newdb Mydb
	db, err := gorm.Open(sqlite.Open("./db/contacts.db"), &gorm.Config{})
	if err != nil {
		return newdb, err
	}
	newdb.db = *db
	return newdb, nil
}

func main() {
	var choice int
	db, err := NewDb()
	if err != nil {
		panic(err)
	}

	for choice != 6 {
		fmt.Println("Enter your choice:")
		fmt.Println("1. Insert Contact")
		fmt.Println("2. Find Contact(s)")
		fmt.Println("3. Update Name")
		fmt.Println("4. Update Number")
		fmt.Println("5. Delete Contact")
		fmt.Println("6. Exit")

		fmt.Printf("\n=============================\n")
		fmt.Scanln(&choice)

		switch {
		case choice == 1:
			var name, number string
			fmt.Printf("=============================\n")
			fmt.Printf("Name: ")
			fmt.Scanln(&name)
			fmt.Printf("Number: ")
			fmt.Scanln(&number)
			db.create(name, number)

		case choice == 2:
			var name string
			fmt.Printf("=============================\n")
			fmt.Printf("Name: ")
			fmt.Scanln(&name)
			contacts := db.find(name)

			for _, contact := range contacts {
				fmt.Printf("Name: %s; Number: %s", contact.name, contact.number)
			}

		case choice == 3:
			var name string
			var id uint
			fmt.Printf("=============================\n")
			fmt.Printf("ID: ")
			fmt.Scanln(&id)
			fmt.Printf("New Name: ")
			fmt.Scanln(&name)
			db.updateName(id, name)

		case choice == 4:
			var number string
			var id uint
			fmt.Printf("=============================\n")
			fmt.Printf("ID: ")
			fmt.Scanln(&id)
			fmt.Printf("New Number: ")
			fmt.Scanln(&number)
			db.updateName(id, number)

		case choice == 5:
			var id uint
			fmt.Printf("=============================\n")
			fmt.Printf("ID: ")
			fmt.Scanln(&id)
			db.delete(id)

		case choice == 6:
			break

		default:
			fmt.Println("Invalid Choice!!")
		}

		fmt.Printf("=============================\n\n")

	}
}
