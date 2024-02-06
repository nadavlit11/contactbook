package main

import (
	"boilerplate/dao"
	"boilerplate/services"
)

var contactsService services.ContactsService

var pageSize = 10

func InitServices() {
	db := dao.NewDatabase()
	db.Connect()
	contactsService = services.NewContactsService(db)
}
