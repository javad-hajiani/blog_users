package users

import (
	"fmt"
	"github.com/javad-hajiani/blog_users/datasources/mysql"
	"github.com/javad-hajiani/blog_users/utils/dateTime"
	"github.com/javad-hajiani/blog_users/utils/errors"
)

var (
	usersDb = make(map[int64]*User)
)

const (
	InsertQuery = "INSERT INTO users (first_name,last_name,password,email,date_created) VALUES (?,?,?,?,?)"
)

func (user *User) Save() *errors.RestError {
	if err := mysql.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := mysql.Client.Prepare(InsertQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dateTime.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName,user.LastName,user.Password,user.Email,user.DateCreated)
	if saveErr != nil {
		return errors.NewInternalServerError(saveErr.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error on getting last user id for %v",userId))
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestError  {
	return nil
}