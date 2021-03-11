package model

import (
	"database/sql"
	"fmt"

	"github.com/ribeirohugo/go_users/internal/fault"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Connection interface {
	New() *sql.DB
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

type PostgresCon struct {
	Config Config
	sql    *sql.DB
}

type MysqlCon struct {
	Config Config
	sql    *sql.DB
}

func connect(driver string, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	fault.HandleFatalError("", err)

	return db
}

func (con *PostgresCon) New() *sql.DB {
	plsql := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		con.Config.Host, con.Config.Port, con.Config.User, con.Config.Password, con.Config.Db)

	return connect("postgres", plsql)
}

func (con *MysqlCon) New() *sql.DB {
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		con.Config.User, con.Config.Password, con.Config.Host, con.Config.Port, con.Config.Db)

	return connect("mysql", mysql)
}

func (con *PostgresCon) getAll() (*sql.Rows, error) {
	return con.sql.Query("SELECT * FROM users")
}

func (con *MysqlCon) addUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("INSERT INTO users (name, password, email, phone, timestamp)"+
		"VALUES (%s, %s, %s, %s, %d)",
		user.Name, user.Password, user.Email, user.Phone, user.Timestamp)

	return con.sql.Query(query)
}

func (con *PostgresCon) addUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("INSERT INTO users (name, password, email, phone, timestamp)"+
		"VALUES (%s, %s, %s, %s, %d)",
		user.Name, user.Password, user.Email, user.Phone, user.Timestamp)

	return con.sql.Query(query)
}

func (con *PostgresCon) removeUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("DELETE FROM users "+
		"WHERE email = \"%s\" OR phone = \"%s\" ",
		user.Email, user.Phone)

	return con.sql.Query(query)
}

func (con *MysqlCon) removeUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("DELETE FROM users "+
		"WHERE email = \"%s\" OR phone = \"%s\" ",
		user.Email, user.Phone)

	return con.sql.Query(query)
}

func (con *PostgresCon) updateUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("UPDATE users "+
		"SET name = \"%s\", password = \"%s\", email=\"%s\", phone=\"%s\" , timestamp=%d "+
		"WHERE email = \"%s\" OR phone = \"%s\" ",
		user.Name, user.Password, user.Email, user.Phone, user.Timestamp, user.Email, user.Phone)

	return con.sql.Query(query)
}

func (con *MysqlCon) updateUser(user User) (*sql.Rows, error) {
	query := fmt.Sprintf("UPDATE users "+
		"SET name = \"%s\", password = \"%s\", email=\"%s\", phone=\"%s\" , timestamp=%d "+
		"WHERE email = \"%s\" OR phone = \"%s\" ",
		user.Name, user.Password, user.Email, user.Phone, user.Timestamp, user.Email, user.Phone)

	return con.sql.Query(query)
}
