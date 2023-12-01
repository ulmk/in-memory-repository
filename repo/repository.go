package repo

import (
	"errors"
	"fmt"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

type DataRepository interface {
	GetByID(id int) (*User, error)
	FetAll() ([]*User, error)
	Create(data *User) error
	Update(data *User) error
	Delete(id int) error
}

type InMemoryRepo struct {
	data []*User
}

func (dr *InMemoryRepo) GetByID(id int) (*User, error) {
	for _, data := range dr.data {
		if data.ID == id {
			return data, nil
		}
	}
	return nil, errors.New("data not found in GetByID ")
}

func (dr *InMemoryRepo) GetAll() ([]*User, error) {
	return dr.data, nil
}

func (dr *InMemoryRepo) Create(newdata *User) error {
	newdata.ID = len(dr.data) + 1
	dr.data = append(dr.data, newdata)
	return nil
}

func (dr *InMemoryRepo) Update(newdata *User) error {
	for i, d := range dr.data {
		if d.ID == newdata.ID {
			dr.data[i] = newdata
			return nil
		}
	}
	return errors.New("data not found in Update")
}

func (dr *InMemoryRepo) Delete(id int) error {
	for i, d := range dr.data {
		if d.ID == id {
			dr.data = append(dr.data[:i], dr.data[i+1:]...)
			return nil
		}
	}
	return errors.New("data not found in Delete")
}
