package handlers

import (
	"database/sql"
	"dreadfiles/curso-golang/13-mysql/model"
	"fmt"
	"log"
)

func ListContacts(db *sql.DB) {
	//query contacts
	query := "SELECT * FROM contact"

	//Execute query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Show rows
	fmt.Println("***********************************************")
	fmt.Println("************* CONTACT LIST ********************")

	for rows.Next() {
		//New contact
		contact := model.Contact{}

		//To allow null values
		var valueEmail sql.NullString
		//Read values from row
		err := rows.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		//Email validation
		contact.Email = "No email address provided"
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		}

		//Show each row
		fmt.Printf("[ID:%d][Name:%s][Email:%s][Phone:%s]\n",
			contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("***********************************************")
}

func GetContactByID(db *sql.DB, contactID int) {
	fmt.Println("************** GET CONTACT BY ID **************")

	//Get contact by ID query
	query := "SELECT * FROM contact WHERE id = ?"

	//Execute query row
	row := db.QueryRow(query, contactID)

	//New contact
	contact := model.Contact{}
	var valueEmail sql.NullString

	//Get contact information
	err := row.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("ID %d not found", contactID)
			return
		}
	}

	//Validate email
	contact.Email = "No email address provided"
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	}

	//Show contact row
	fmt.Printf("[ID:%d][Name:%s][Email:%s][Phone:%s]\n",
		contact.ID, contact.Name, contact.Email, contact.Phone)

	fmt.Println("***********************************************")
}

func CreateContact(db *sql.DB, contact model.Contact) {
	fmt.Println("*************** CREATE CONTACT ****************")

	//Create contact query
	query := "INSERT INTO contact(name, email, phone) VALUES(?, ?, ?)"

	//Execute insert query
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Println("unable to insert row: ", err.Error())
		return
	}

	fmt.Println("contact added successfully")
	fmt.Println("***********************************************")
}

func UpdateContact(db *sql.DB, contact model.Contact) {
	fmt.Println("*************** UPDATE CONTACT ****************")

	//Update contact query
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	//Execute update query
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.ID)
	if err != nil {
		log.Println("unable to update row: ", err.Error())
		return
	}

	fmt.Println("contact updated successfully")
	fmt.Println("***********************************************")
}

func DeleteContact(db *sql.DB, contactID int) {
	fmt.Println("*************** DELETE CONTACT ****************")

	//Delete contact query
	query := "DELETE FROM contact WHERE id = ?"

	//Execute delete query
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Println("unable to delete row: ", err.Error())
		return
	}

	fmt.Println("contact deleted successfully")
	fmt.Println("***********************************************")
}
