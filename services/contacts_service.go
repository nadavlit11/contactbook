package services

import (
	"contactbook/dao"
	"contactbook/models"
	"github.com/gofiber/fiber/v2/log"
	"sync"
)

var contactsServiceOnce sync.Once
var contactsService ContactsService

type ContactsService interface {
	GetContacts(userId int) ([]models.Contact, error)
	GetContactsPage(userId int, page int, pageSize int) ([]models.Contact, error)
	InsertContact(userId int, contact models.Contact) error
	Search(userId int, search models.Contact) ([]models.Contact, error)
	Edit(userId int, contact models.Contact) error
	Delete(userId int, id int) error
}

type ContactsServiceImpl struct {
	contactsDao dao.ContactsDao
}

func NewContactsService(
	contactsDao dao.ContactsDao,
) ContactsService {
	contactsServiceOnce.Do(func() {
		contactsService = &ContactsServiceImpl{
			contactsDao: contactsDao,
		}
	})
	return contactsService
}

func (service *ContactsServiceImpl) GetContacts(userId int) ([]models.Contact, error) {
	users, err := service.contactsDao.GetContacts(userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (service *ContactsServiceImpl) GetContactsPage(userId int, page int, pageSize int) ([]models.Contact, error) {
	offset := (page - 1) * pageSize
	limit := pageSize - 1

	users, err := service.contactsDao.GetPage(userId, offset, limit)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (service *ContactsServiceImpl) InsertContact(userId int, contact models.Contact) error {
	err := service.contactsDao.Insert(userId, contact)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (service *ContactsServiceImpl) Search(userId int, search models.Contact) ([]models.Contact, error) {
	contacts, err := service.contactsDao.Search(userId, search)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return contacts, nil
}

func (service *ContactsServiceImpl) Edit(userId int, search models.Contact) error {
	err := service.contactsDao.Edit(userId, search)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (service *ContactsServiceImpl) Delete(userId int, id int) error {
	err := service.contactsDao.Delete(userId, id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
