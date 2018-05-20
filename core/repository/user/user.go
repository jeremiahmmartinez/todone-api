package user

import (
	"todone-api/model"
	"todone-api/core/database"
	"fmt"
	"strings"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"todone-api/settings"
	utility "todone-api/core/utility"
)

func UsernameExists(username string) (bool, error) {
	return rowExists("SELECT id FROM users WHERE username=?", username)
}

func EmailExists(email string) (bool, error) {
	return rowExists("SELECT id FROM users WHERE email=?", email)
}

func rowExists(query string, args ...interface{}) (bool, error) {
	adapter := database.GetDatabaseAdapter()

	adapter.Open()
	defer adapter.Close()

	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)

	err := adapter.DB.QueryRow(query, args...).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}

func CreateUser(user *model.User) (bool, error) {
	adapter := database.GetDatabaseAdapter()

	adapter.Open()
	defer adapter.Close()


	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hashedPassword)

	stm, err := adapter.DB.Prepare("insert into users (username, password, email, created) values (?, ?, ?, NOW())")

	if err != nil {
		return false, err
	}
	defer stm.Close()

	if err = writeUser(stm, user); err != nil {
		return false, err
	}

	return true, nil
}

func writeUser(stm *sql.Stmt, user *model.User) error {
	params := []interface{}{
		user.Username,
		user.Password,
		user.Email,
	}

	if _, err := stm.Exec(params...); err != nil {
		return err
	}
	return nil
}

func GetUser(args map[string]interface{}) (model.User, error) {
	adapter := database.GetDatabaseAdapter()

	adapter.Open()
	defer adapter.Close()

	user := model.User{}

	querySql := "select * from users"

	var values []interface{}
	var where []string

	for _, keys := range []string{"id", "username", "firstname", "lastname", "email"} {

		if value, ok := args[keys]; ok {
			values = append(values, value)
			where = append(where, fmt.Sprintf("%v = ?", keys))
		}
	}

	if len(values) == 0 {
		return user, errors.New("not enough parameters supplied")
	}

	querySql = querySql  + " WHERE " + strings.Join(where, " AND ")

	rows, err := adapter.DB.Query(querySql, values...)

	if err != nil {
		return user, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.Created)

		if err != nil {
			return user, err
		}
	}
	err = rows.Err()

	return user, err
}

var OrderByFields = []string{
	"id",
	"username",
	"firstname",
	"lastname",
	"email",
	"created",
}

func GetUsers(offset int, limit int, orderBy string, orderDirection string) ([]model.User, error) {
	if limit == 0 {
		limit = settings.Get().DefaultPageLimit
	}

	var (
		orderByValue        string
		orderDirectionValue string
		)

	if !utility.Contains(OrderByFields, orderBy) {
		orderByValue = settings.Get().DefaultUserSortKey
	} else {
		orderByValue = orderBy
	}

	if orderDirection == "ASC" || orderDirection == "DESC" {
		orderDirectionValue = orderDirection
	} else {
		orderDirectionValue = settings.Get().DefaultUserSortDirection
	}

	adapter := database.GetDatabaseAdapter()

	adapter.Open()
	defer adapter.Close()

	users := make(model.Users, 0)

	querySql := fmt.Sprintf("SELECT * FROM users ORDER BY %s %s LIMIT ?, ?", orderByValue, orderDirectionValue)

	rows, err := adapter.DB.Query(querySql, offset, limit)

	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {

		user := model.User{}

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.Created)

		if err != nil {
			continue
		}

		users = append(users, user)
	}
	err = rows.Err()

	return users, err
}
