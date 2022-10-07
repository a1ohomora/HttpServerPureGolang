package service

import (
	"HttpServerPureGolang/main/configs"
	u "HttpServerPureGolang/main/utils"
	"fmt"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Contact name should not be empty"), false
	}
	if len(contact.Name) > 40 {
		return u.Message(false, "Contact name length should be less than 41 chars"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should not be empty"), false
	}
	if len(contact.Phone) > 15 {
		return u.Message(false, "Phone number length should be less than 16 numbers"), false
	}

	return u.Message(true, "success"), true
}

func CreateContact(contact *Contact) map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	
	response := u.Message(true, "success")

	_, insertError := configs.GetDB().Exec("insert into contacts (name, phone) values ($1, $2)", contact.Name, contact.Phone)
	if insertError != nil {
		fmt.Println(insertError)
		response = u.Message(false, "insertion error")
		return response
	}

	response["contact"] = contact
	return response
}

func GetContact(id int) *Contact {
	contact := &Contact{}
	row := configs.GetDB().QueryRow("select * from contacts where id=$1", id)
	err := row.Scan(&contact.Id, &contact.Name, &contact.Phone)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contact
}

func GetContacts() []*Contact {

	contacts := make([]*Contact, 0)
	rows, err := configs.GetDB().Query("select * from contacts")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		contact := Contact{}
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, &contact)
	}
	return contacts
}
