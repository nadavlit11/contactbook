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
	db dao.Database,
) ContactsService {
	return &ContactsServiceImpl{
		db: db,
	}
}

func TestContactsServiceImpl_once(t *testing.T) {
	databaseMock := new(DatabaseMock)

	_ = NewContactsService(databaseMock)
	_ = NewContactsService(databaseMock)
}

func TestContactsServiceImpl_GetContactsPage(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	databaseMock.On("GetPage", 0, 9).Return([]models.Contact{}, nil)

	contacts, err := cs.GetContactsPage(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(contacts))
}

func TestContactsServiceImpl_GetContactsPageError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	databaseMock.On("GetPage", 0, 9).Return([]models.Contact{}, errors.New("error"))

	_, err := cs.GetContactsPage(1, 10)

	assert.Error(t, err)
}

func TestContactsServiceImpl_InsertContact(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Insert", contact).Return(nil)

	err := cs.InsertContact(contact)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_InsertContactError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Insert", contact).Return(errors.New("error"))

	err := cs.InsertContact(contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Search(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Search", contact).Return([]models.Contact{}, nil)

	contacts, err := cs.Search(contact)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(contacts))
}

func TestContactsServiceImpl_SearchError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Search", contact).Return([]models.Contact{}, errors.New("error"))

	_, err := cs.Search(contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Edit(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Edit", contact).Return(nil)

	err := cs.Edit(contact)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_EditError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	contact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	databaseMock.On("Edit", contact).Return(errors.New("error"))

	err := cs.Edit(contact)

	assert.Error(t, err)
}

func TestContactsServiceImpl_Delete(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	databaseMock.On("Delete", 1).Return(nil)

	err := cs.Delete(1)

	assert.NoError(t, err)
}

func TestContactsServiceImpl_DeleteError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	cs := NewContactsServiceTest(databaseMock)

	databaseMock.On("Delete", 1).Return(errors.New("error"))

	err := cs.Delete(1)

	assert.Error(t, err)
}

type DatabaseMock struct {
	mock.Mock
}

func (d *DatabaseMock) Connect() {
	_ = d.Called()
}

func (d *DatabaseMock) Insert(contact models.Contact) error {
	args := d.Called(contact)
	return args.Error(0)
}

func (d *DatabaseMock) GetPage(offset int, limit int) ([]models.Contact, error) {
	args := d.Called(offset, limit)
	return args.Get(0).([]models.Contact), args.Error(1)
}

func (d *DatabaseMock) Search(search models.Contact) ([]models.Contact, error) {
	args := d.Called(search)
	return args.Get(0).([]models.Contact), args.Error(1)
}

func (d *DatabaseMock) Edit(contact models.Contact) error {
	args := d.Called(contact)
	return args.Error(0)
}

func (d *DatabaseMock) Delete(id int) error {
	args := d.Called(id)
	return args.Error(0)
}
