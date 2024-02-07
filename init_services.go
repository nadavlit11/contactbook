package main

import (
	"contactbook/dao"
	"contactbook/services"
)

var contactsService services.ContactsService

var pageSize = 10

func InitServices() {
	contactsDao := dao.NewContactsDao()
	contactsDao.Connect()
	contactsService = services.NewContactsService(contactsDao)
}
