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
	GetContactsPage(page int, pageSize int) ([]models.Contact, error)
	InsertContact(models.Contact) error
	Search(search models.Contact) ([]models.Contact, error)
	Edit(contact models.Contact) error
	Delete(id int) error
}

type ContactsServiceImpl struct {
	db dao.Database
}

func NewContactsService(
	db dao.Database,
) ContactsService {
	contactsServiceOnce.Do(func() {
		contactsService = &ContactsServiceImpl{
			db: db,
		}
	})
	return contactsService
}

func (service *ContactsServiceImpl) GetContactsPage(page int, pageSize int) ([]models.Contact, error) {
	offset := (page - 1) * pageSize
	limit := pageSize - 1

	users, err := service.db.GetPage(offset, limit)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (service *ContactsServiceImpl) InsertContact(contact models.Contact) error {
	err := service.db.Insert(contact)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (service *ContactsServiceImpl) Search(search models.Contact) ([]models.Contact, error) {
	contacts, err := service.db.Search(search)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return contacts, nil
}

func (service *ContactsServiceImpl) Edit(search models.Contact) error {
	err := service.db.Edit(search)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (service *ContactsServiceImpl) Delete(id int) error {
	err := service.db.Delete(id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
