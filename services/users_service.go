package services

import (
	"contactbook/dao"
	"contactbook/models"
	"github.com/gofiber/fiber/v2/log"
	"sync"
)

var usersServiceOnce sync.Once
var usersService UsersService

type UsersService interface {
	CreateUser(name string) (int, error)
}

type UsersServiceImpl struct {
	usersDao        dao.UsersDao
	contactsService ContactsService
}

func NewUsersService(
	usersDao dao.UsersDao,
	contactsService ContactsService,
) UsersService {
	usersServiceOnce.Do(func() {
		usersService = &UsersServiceImpl{
			usersDao:        usersDao,
			contactsService: contactsService,
		}
	})
	return usersService
}

func (service *UsersServiceImpl) CreateUser(name string) (int, error) {
	user := models.User{
		Name: name,
	}
	userId, err := service.usersDao.Insert(user)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	err = service.contactsService.InitUserContactBook(userId)

	return userId, nil
}
