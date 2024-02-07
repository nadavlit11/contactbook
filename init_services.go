package main

import (
	"contactbook/dao"
	"contactbook/services"
)

var contactsService services.ContactsService
var usersService services.UsersService

var pageSize = 10

func InitServices() {
	contactsDao := dao.NewContactsDao()
	contactsDao.Connect()
	contactsService = services.NewContactsService(contactsDao)

	usersDao := dao.NewUsersDao()
	usersDao.Connect()
	usersService = services.NewUsersService(usersDao, contactsService)
}
