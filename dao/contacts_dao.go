package dao

import (
	"contactbook/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"strings"
	"sync"
)

var (
	contactsDb map[int][]models.Contact
	mu         sync.Mutex
	autoIncId  int
)

var contactsDaoOnce sync.Once
var contactsDao ContactsDao

type ContactsDao interface {
	Connect()
	Insert(userId int, contact models.Contact) error
	GetPage(userId int, offset int, limit int) ([]models.Contact, error)
	Search(userId int, search models.Contact) ([]models.Contact, error)
	Edit(userId int, contact models.Contact) error
	Delete(userId int, id int) error
}

type ContactsDaoImpl struct {
}

func NewContactsDao() ContactsDao {
	contactsDaoOnce.Do(func() {
		contactsDao = &ContactsDaoImpl{}
	})
	return contactsDao
}

func (d *ContactsDaoImpl) Connect() {
	contactsDb = map[int][]models.Contact{}
	fmt.Println("Connected with ContactsDao")
}

func (d *ContactsDaoImpl) Insert(userId int, contact models.Contact) error {
	mu.Lock()

	userBook, ok := contactsDb[userId]
	if !ok {
		log.Error("user not found")
		return errors.New("user not found")
	}

	autoIncId++
	contact.ID = autoIncId
	userBook = append(userBook, contact)
	contactsDb[userId] = userBook

	mu.Unlock()
	return nil
}

func (d *ContactsDaoImpl) GetPage(userId int, offset int, limit int) ([]models.Contact, error) {
	userBook, ok := contactsDb[userId]
	if !ok {
		log.Error("user not found")
		return nil, errors.New("user not found")
	}

	if len(userBook) == 0 {
		return nil, nil
	}

	if offset > len(userBook) {
		return nil, fmt.Errorf("offset + limit is greater than the length of the contactsDao")
	}

	offsetLimit := offset + limit
	if offsetLimit > len(userBook) {
		offsetLimit = len(userBook)
	}
	return userBook[offset:offsetLimit], nil
}

func (d *ContactsDaoImpl) Search(userId int, search models.Contact) ([]models.Contact, error) {
	userBook, ok := contactsDb[userId]
	if !ok {
		log.Error("user not found")
		return nil, errors.New("user not found")
	}

	filteredContacts := userBook

	if len(search.FirstName) > 0 {
		filteredContacts = lo.Filter(filteredContacts, func(item models.Contact, _ int) bool {
			return strings.EqualFold(item.FirstName, search.FirstName)
		})
	}

	if len(search.LastName) > 0 {
		filteredContacts = lo.Filter(filteredContacts, func(item models.Contact, _ int) bool {
			return strings.EqualFold(item.LastName, search.LastName)
		})
	}

	if len(search.Phone) > 0 {
		filteredContacts = lo.Filter(filteredContacts, func(item models.Contact, _ int) bool {
			return strings.EqualFold(item.Phone, search.Phone)
		})
	}

	if len(search.Address) > 0 {
		filteredContacts = lo.Filter(filteredContacts, func(item models.Contact, _ int) bool {
			return strings.EqualFold(item.Address, search.Address)
		})
	}

	return filteredContacts, nil
}

func (d *ContactsDaoImpl) Edit(userId int, contact models.Contact) error {
	userBook, ok := contactsDb[userId]
	if !ok {
		log.Error("user not found")
		return errors.New("user not found")
	}

	_, i, found := lo.FindIndexOf(userBook, func(item models.Contact) bool {
		return item.ID == contact.ID
	})

	if !found {
		log.Error("contact not found")
		return errors.New("contact not found")
	}

	userBook[i] = contact
	contactsDb[userId] = userBook

	return nil
}

func (d *ContactsDaoImpl) Delete(userId int, id int) error {
	userBook, ok := contactsDb[userId]
	if !ok {
		log.Error("user not found")
		return errors.New("user not found")
	}

	_, i, found := lo.FindIndexOf(userBook, func(item models.Contact) bool {
		return item.ID == id
	})

	if !found {
		log.Error("contact not found")
		return errors.New("contact not found")
	}

	userBook = append(userBook[:i], userBook[i+1:]...)
	contactsDb[userId] = userBook

	return nil
}
