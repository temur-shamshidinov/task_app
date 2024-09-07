package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/task_app/models"
	"github.com/temur-shamshidinov/task_app/storage/repoI"
)

type userRepo struct {
	con *pgx.Conn
}

func NewUserRepo(con *pgx.Conn) repoI.UserRepoI {
	return &userRepo{
		con: con,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user models.User) error {

	query :=
		`
	 	INSERT INTO 
			users (
				username, 
				email, 
				password_hash
			) VALUES ($1, $2, $3)`

	_, err := u.con.Exec(ctx, query, user.Username, user.Email, user.PasswordHash)
	if err != nil {

		log.Println("Error inserting user:", err)
		return err
	}

	return nil

}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) ([]models.User, error) {

	var users []models.User

	query := `
		SELECT 
			id, 
			username, 
			email, 
			password_hash 
		FROM 
			users 
		WHERE email = $1`

	rows, err := u.con.Query(ctx, query, email)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user models.User

		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error retrieving rows:", err)
		return nil, err
	}
	return users, nil
}
