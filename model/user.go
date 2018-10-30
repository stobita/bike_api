package model

import (
	"time"
)

// column name `ID` will be the default primary field

// User database
type User struct {
	ID           int64     `xorm:"id pk" json:"id"`
	Name         string    `xorm:"name"`
	Email        string    `xorm:"email"`
	Password     string    `xorm:"password"`
	Image        string    `xorm:"image"`
	Introduction string    `xorm:"introduction"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

// NewUser create User instance
func NewUser() *User {
	return new(User)
}

// Create create new user
func (u *User) Create(name, email, password string) error {
	u.Name = name
	u.Email = email
	u.Password = password

	_, err := engine.Insert(u)
	if err != nil {
		return err
	}

	return nil
}

// FindOneByEmail find user by email
func (u *User) FindOneByEmail(email string) bool {
	u.Email = email
	has, _ := engine.Get(u)
	return has
}

// Update User update
func (u *User) Update(userID int64) error {
	_, err := engine.Id(userID).Cols("email").Update(&u)
	if err != nil {
		return err
	}
	return nil
}
