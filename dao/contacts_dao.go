package dao

import (
	"contactbook/database"
	"contactbook/models"
	"github.com/gofiber/fiber/v2/log"
	"sync"
)

var ()

var contactsDaoOnce sync.Once
var contactsDao ContactsDao

type ContactsDao interface {
	Insert(userId int, contact models.Contact) error
	GetContacts(userId int) ([]models.Contact, error)
	GetPage(userId int, offset int, limit int) ([]models.Contact, error)
	Search(userId int, search models.Contact) ([]models.Contact, error)
	Edit(userId int, contact models.Contact) error
	Delete(userId int, id int) error
}

type ContactsDaoImpl struct {
	contactsDb     map[int][]models.Contact
	mu             sync.Mutex
	autoIncId      int
	postgresClient database.PostgresClient
}

func NewContactsDao(
	postgresClient database.PostgresClient,
) ContactsDao {
	contactsDaoOnce.Do(func() {
		contactsDao = &ContactsDaoImpl{
			postgresClient: postgresClient,
		}
	})
	return contactsDao
}

func (d *ContactsDaoImpl) Insert(userId int, contact models.Contact) error {
	pConn := d.postgresClient.GetConn()
	query := `
		INSERT INTO contacts
		(user_id, first_name, last_name, phone, address)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := pConn.Exec(query, userId, contact.FirstName, contact.LastName, contact.Phone, contact.Address)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (d *ContactsDaoImpl) GetContacts(userId int) ([]models.Contact, error) {
	pConn := d.postgresClient.GetConn()
	query := `
		SELECT * FROM contacts
		WHERE user_id = $1`

	rows, err := pConn.Query(query, userId)
	defer rows.Close()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err = rows.Scan(&contact.ID, &contact.UserId, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Address)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (d *ContactsDaoImpl) GetPage(userId int, offset int, limit int) ([]models.Contact, error) {
	pConn := d.postgresClient.GetConn()
	query := `
		SELECT * FROM contacts
		WHERE user_id = $1
		OFFSET $2 LIMIT $3`

	rows, err := pConn.Query(query, userId, offset, limit)
	defer rows.Close()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var contacts []models.Contact
	err = rows.Scan(&contacts)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return contacts, nil
}

func (d *ContactsDaoImpl) Search(userId int, search models.Contact) ([]models.Contact, error) {
	pConn := d.postgresClient.GetConn()
	query := `
		SELECT * FROM contacts
		WHERE user_id = $1`
	if search.FirstName != "" {
		query += ` AND first_name = $2`
	}
	if search.LastName != "" {
		query += ` AND last_name = $3`
	}
	if search.Phone != "" {
		query += ` AND phone = $4`
	}
	if search.Address != "" {
		query += ` AND address = $5`
	}

	rows, err := pConn.Query(query, userId)
	defer rows.Close()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var contacts []models.Contact
	err = rows.Scan(&contacts)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return contacts, nil
}

func (d *ContactsDaoImpl) Edit(userId int, contact models.Contact) error {
	pConn := d.postgresClient.GetConn()
	query := `
		UPDATE contacts
		SET first_name = $3, last_name = $4, phone = $5, address = $6
		WHERE  user_id = $1 AND id = $2`

	_, err := pConn.Exec(query, userId, contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Address)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (d *ContactsDaoImpl) Delete(userId int, id int) error {
	pConn := d.postgresClient.GetConn()
	query := `
		DELETE FROM contacts
		WHERE user_id = $1 AND id = $2`

	_, err := pConn.Exec(query, userId, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
