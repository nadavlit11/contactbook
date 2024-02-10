package main

import (
	"contactbook/dao"
	"contactbook/database"
	"contactbook/services"
	"github.com/go-co-op/gocron"
	"time"
)

var contactsService services.ContactsService
var usersService services.UsersService
var postgresClient database.PostgresClient

var pageSize = 10

func InitServices() {
	postgresClient = database.NewPostgresClient()
	contactsDao := dao.NewContactsDao(postgresClient)
	redisClient := database.NewRedisClient()
	contactsCacheDao := dao.NewContactsCacheDao(redisClient)
	contactsCacheService := services.NewContactsCacheService(contactsCacheDao)
	contactsService = services.NewContactsService(contactsDao, contactsCacheService)

	usersDao := dao.NewUsersDao(postgresClient, redisClient)
	usersService = services.NewUsersService(usersDao, contactsService, contactsCacheService)

	postgresClient = database.NewPostgresClient()

	syncDataService := services.NewSyncDataService()
	s := gocron.NewScheduler(time.UTC)
	s = s.Cron("0 * * * *").SingletonMode()
	_, _ = s.Do(func() {
		syncDataService.Sync()
	})
}
