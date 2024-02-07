package main

import (
	"contactbook/dao"
	"contactbook/services"
)

var contactsService services.ContactsService

var pageSize = 10

func InitServices() {
	db := dao.NewDatabase()
	db.Connect()
	contactsService = services.NewContactsService(db)
}
