package main

import (
	"dreadfiles/curso-golang/13-mysql/database"
	"dreadfiles/curso-golang/13-mysql/handlers"
	"dreadfiles/curso-golang/13-mysql/model"
	"fmt"
	"log"
)

func main() {
	fmt.Println("***********************************************")
	fmt.Println("*********** Starting 13-mysql *****************")

	//Database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Show contact list
	handlers.ListContacts(db)

	//Show contact by ID
	handlers.GetContactByID(db, 2)
	handlers.GetContactByID(db, 6)
	handlers.GetContactByID(db, 1)

	//Add new contact
	createContact := model.Contact{
		Name:  "New Contact",
		Email: "newcontact@mail.com",
		Phone: "1234567890",
	}
	handlers.CreateContact(db, createContact)

	//Update contact
	updateContact := model.Contact{
		ID:    2,
		Name:  "Modified Contact",
		Email: "modifiedcontact@mail.com",
		Phone: "1234567890",
	}
	handlers.UpdateContact(db, updateContact)

	//Delete contact
	handlers.DeleteContact(db, 8)

	//Show contact list
	handlers.ListContacts(db)

	fmt.Println("************ Ending 13-mysql ******************")
	fmt.Println("***********************************************")
}
