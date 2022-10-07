package controllers

import (
	"HttpServerPureGolang/main/service"
	u "HttpServerPureGolang/main/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateContact = func(writer http.ResponseWriter, request *http.Request) {
	contact := &service.Contact{}

	err := json.NewDecoder(request.Body).Decode(contact)
	if err != nil {
		u.Respond(writer, u.Message(false, "Error while decoding request body"))
		return
	}

	response := service.CreateContact(contact)
	u.Respond(writer, response)
}

var GetContact = func(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	data := service.GetContact(id)

	var response map[string]interface{}
	if data != nil {
		response = u.Message(true, "success")
		response["data"] = data
	} else {
		response = u.Message(false, "not found")
	}
	u.Respond(writer, response)
}

var GetContacts = func(writer http.ResponseWriter, request *http.Request) {
	data := service.GetContacts()
	response := u.Message(true, "success")
	response["data"] = data
	u.Respond(writer, response)
}
