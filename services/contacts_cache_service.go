package services

import (
	"contactbook/dao"
	"contactbook/models"
	"sync"
)

var contactsCacheServiceOnce sync.Once
var contactsCacheService ContactsCacheService

type ContactsCacheService interface {
	CacheContacts(userId int, contacts []models.Contact) error
	GetContacts(userId int) ([]models.Contact, error)
	AddContact(userId int, contact models.Contact) error
}

type ContactsCacheServiceImpl struct {
	contactsCacheDao dao.ContactsCacheDao
}

func NewContactsCacheService(
	contactsCacheDao dao.ContactsCacheDao,
) ContactsCacheService {
	contactsCacheServiceOnce.Do(func() {
		contactsCacheService = &ContactsCacheServiceImpl{
			contactsCacheDao: contactsCacheDao,
		}
	})
	return contactsCacheService
}

func (service *ContactsCacheServiceImpl) CacheContacts(userId int, contacts []models.Contact) error {
	err := service.contactsCacheDao.CacheContacts(userId, contacts)
	if err != nil {
		return err
	}

	return nil
}

func (service *ContactsCacheServiceImpl) GetContacts(userId int) ([]models.Contact, error) {
	contacts, err := service.contactsCacheDao.GetContacts(userId)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (service *ContactsCacheServiceImpl) AddContact(userId int, contact models.Contact) error {
	return service.contactsCacheDao.AddContact(userId, contact)
}
