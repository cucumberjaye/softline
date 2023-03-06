package postgresdb

import (
	"database/sql"
	"github.com/cucumberjaye/softline/internal/models"
)

var (
	selectStmt *sql.Stmt
	insertStmt *sql.Stmt
)

func newAuthStmts(db *sql.DB) error {
	var err error

	insertStmt, err = db.Prepare("INSERT INTO users (login, email, password_hash, phone_number) values ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return err
	}

	selectStmt, err = db.Prepare("SELECT id, login, email, phone_number FROM users WHERE email=$1 AND password_hash=$2")
	if err != nil {
		return err
	}

	return nil
}

func (r *Postgres) GetUser(loginUser models.LoginUser) (models.User, error) {
	var user models.User

	row := selectStmt.QueryRow(loginUser.Email, loginUser.Password)
	if err := row.Scan(&user.Id, &user.Login, &user.Email, &user.PhoneNumber); err != nil {
		return user, err
	}
	return user, nil
}

func (r *Postgres) CreateUser(user models.RegisterUser) (int, error) {
	var id int

	row := insertStmt.QueryRow(user.Login, user.Email, user.Password, user.PhoneNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
