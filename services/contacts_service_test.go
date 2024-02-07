package services

import (
	"contactbook/dao"
	"contactbook/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func NewContactsServiceTest(
	contactsDao dao.ContactsDao,
) ContactsService {
	return &ContactsServiceImpl{
		contactsDao: contactsDao,
	}
}

func TestContactsServiceImpl_once(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)

	_ = NewContactsService(contactsDaoMock)
	_ = NewContactsService(contactsDaoMock)
}

func TestContactsServiceImpl_GetContactsPage(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contactsDaoMock.On("GetPage", 1, 0, 9).Return([]models.Contact{}, nil)

	contacts, err := cs.GetContactsPage(1, 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(contacts))
}

func TestContactsServiceImpl_GetContactsPageError(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contactsDaoMock.On("GetPage", 1, 0, 9).Return([]models.Contact{}, errors.New("error"))

	_, err := cs.GetContactsPage(1, 1, 10)

	assert.Error(t, err)
}

func TestContactsServiceImpl_InsertContact(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Insert", 1, contact).Return(nil)

	err := cs.InsertContact(1, contact)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_InsertContactError(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Insert", 1, contact).Return(errors.New("error"))

	err := cs.InsertContact(1, contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Search(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Search", 1, contact).Return([]models.Contact{}, nil)

	contacts, err := cs.Search(1, contact)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(contacts))
}

func TestContactsServiceImpl_SearchError(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Search", 1, contact).Return([]models.Contact{}, errors.New("error"))

	_, err := cs.Search(1, contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Edit(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Edit", 1, contact).Return(nil)

	err := cs.Edit(1, contact)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_EditError(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	contactsDaoMock.On("Edit", 1, contact).Return(errors.New("error"))

	err := cs.Edit(1, contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Delete(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contactsDaoMock.On("Delete", 1, 1).Return(nil)

	err := cs.Delete(1, 1)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_DeleteError(t *testing.T) {
	contactsDaoMock := new(ContactsDaoMock)
	cs := NewContactsServiceTest(contactsDaoMock)

	contactsDaoMock.On("Delete", 1, 1).Return(errors.New("error"))

	err := cs.Delete(1, 1)

	assert.Error(t, err)
}

type ContactsDaoMock struct {
	mock.Mock
}

func (d *ContactsDaoMock) Connect() {
	_ = d.Called()
}

func (d *ContactsDaoMock) Insert(userId int, contact models.Contact) error {
	args := d.Called(userId, contact)
	return args.Error(0)
}

func (d *ContactsDaoMock) GetPage(userId int, offset int, limit int) ([]models.Contact, error) {
	args := d.Called(userId, offset, limit)
	return args.Get(0).([]models.Contact), args.Error(1)
}

func (d *ContactsDaoMock) Search(userId int, search models.Contact) ([]models.Contact, error) {
	args := d.Called(userId, search)
	return args.Get(0).([]models.Contact), args.Error(1)
}

func (d *ContactsDaoMock) Edit(userId int, contact models.Contact) error {
	args := d.Called(userId, contact)
	return args.Error(0)
}

func (d *ContactsDaoMock) Delete(userId int, id int) error {
	args := d.Called(userId, id)
	return args.Error(0)
}
