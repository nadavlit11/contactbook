package dao

import (
	"boilerplate/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"strings"
	"sync"
)

var (
	contactsDb []models.Contact
	mu         sync.Mutex
	autoIncId  int
)

var databaseOnce sync.Once
var database Database

type Database interface {
	Connect()
	Insert(user models.Contact) error
	GetPage(offset int, limit int) ([]models.Contact, error)
	Search(search models.Contact) ([]models.Contact, error)
	Edit(contact models.Contact) error
	Delete(id int) error
}

type DatabaseImpl struct {
}

func NewDatabase() Database {
	databaseOnce.Do(func() {
		database = &DatabaseImpl{}
	})
	return database
}

func (d *DatabaseImpl) Connect() {
	contactsDb = make([]models.Contact, 0)
	fmt.Println("Connected with Database")
}

func (d *DatabaseImpl) Insert(contact models.Contact) error {
	mu.Lock()
	autoIncId++
	contact.ID = autoIncId
	contactsDb = append(contactsDb, contact)
	mu.Unlock()
	return nil
}

func (d *DatabaseImpl) GetPage(offset int, limit int) ([]models.Contact, error) {
	if len(contactsDb) == 0 {
		return nil, nil
	}

	if offset > len(contactsDb) {
		return nil, fmt.Errorf("offset + limit is greater than the length of the database")
	}

	offsetLimit := offset + limit
	if offsetLimit > len(contactsDb) {
		offsetLimit = len(contactsDb)
	}
	return contactsDb[offset:offsetLimit], nil
}

func (d *DatabaseImpl) Search(search models.Contact) ([]models.Contact, error) {
	filteredContacts := contactsDb

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

func (d *DatabaseImpl) Edit(contact models.Contact) error {
	_, i, found := lo.FindIndexOf(contactsDb, func(item models.Contact) bool {
		return item.ID == contact.ID
	})

	if !found {
		log.Error("contact not found")
		return errors.New("contact not found")
	}

	contactsDb[i] = contact
	return nil
}

func (d *DatabaseImpl) Delete(id int) error {
	_, i, found := lo.FindIndexOf(contactsDb, func(item models.Contact) bool {
		return item.ID == id
	})

	if !found {
		log.Error("contact not found")
		return errors.New("contact not found")
	}

	contactsDb = append(contactsDb[:i], contactsDb[i+1:]...)

	return nil
}
