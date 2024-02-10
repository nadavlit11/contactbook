package main

import (
	"contactbook/dao"
	"contactbook/database"
	"contactbook/services"
)

var contactsService services.ContactsService
var usersService services.UsersService
var postgresClient database.PostgresClient

var pageSize = 10

func InitServices() {
	postgresClient = database.NewPostgresClient()
	contactsDao := dao.NewContactsDao(postgresClient)
	contactsDao.Connect()
	contactsService = services.NewContactsService(contactsDao)

	redisClient := database.NewRedisClient()
	usersDao := dao.NewUsersDao(postgresClient, redisClient)
	usersService = services.NewUsersService(usersDao, contactsService)

	postgresClient = database.NewPostgresClient()
}
