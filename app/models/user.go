package models

import (
	"fmt"
	"log"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role      string `json:"role"`
}

func NewUser(id, username, password, firstname, lastname, role string) *User {
	return &User{
		ID:        id,
		UserName:  username,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
		Role:      role,
	}
}

func (u *User) TableName() string {
	return GetUserTableName("user")
}

func (u *User) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, username, password, firstname, lastname, role) VALUES (?, ?, ?, ?, ?, ?)", u.TableName())
	_, err := DbConnection.Exec(cmd, u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role)
	if err != nil {
		return err
	}
	return err
}

func (u *User) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET username = ?, password = ?, firstname = ?, lastname = ?, role = ? WHERE id = ?", u.TableName())
	_, err := DbConnection.Exec(cmd, u.UserName, u.Password, u.FirstName, u.LastName, u.Role, u.ID)
	if err != nil {
		return err
	}
	return err
}

func GetUser(id string) (*User, error) {
	tableName := GetUserTableName("user")
	cmd := fmt.Sprintf(`SELECT id, username, password, firstname, lastname, role FROM %s WHERE id = '%s'`,
		tableName, id)
	row := DbConnection.QueryRow(cmd)
	var u User
	err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.FirstName, &u.LastName, &u.Role)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}

func GetUserByUserName(un string) (*User, error) {
	tableName := GetUserTableName("user")
	cmd := fmt.Sprintf(`SELECT id, username, password, firstname, lastname, role FROM %s WHERE username = '%s'`,
		tableName, un)
	row := DbConnection.QueryRow(cmd)
	var u User
	err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.FirstName, &u.LastName, &u.Role)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}
