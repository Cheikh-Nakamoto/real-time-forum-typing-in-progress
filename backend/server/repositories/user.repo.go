package repositories

import (
	"database/sql"
	"fmt"
	db "forum/backend/database"
	opt "forum/backend/database/operators"
	q "forum/backend/database/query"
	"forum/backend/models"
)

type UserRepository struct {
	BaseRepo
}

func (r *UserRepository) init() {
	r.DB = db.DB
	r.TableName = db.USERS_TABLE
}

func (r *UserRepository) GetUserById(userId string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"user_id": opt.Equals(userId)})
	if err == sql.ErrNoRows {
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this Id")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"username": opt.Equals(username)})
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this username")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"email": opt.Equals(email)})
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this username")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	err := r.DB.Update(r.TableName, user, q.WhereOption{"user_id": opt.Equals(user.UserId)})
	return err
}

func (r *UserRepository) DeleteUser(user models.User) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"user_id": opt.Equals(user.UserId)})
	return err
}

func (r *UserRepository) SaveUser(user models.User) error {
	err := r.DB.Insert(r.TableName, user)
	return err
}

func (r *UserRepository) GetAllUsers() (users []models.User, err error) {
	var user models.User
	rows, err := r.DB.GetAllFrom(r.TableName,nil, "username",nil)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		rows.Scan(&user.UserId,&user.Firstname,&user.Lastname,&user.Username,&user.AvatarUrl,&user.Email,&user.Password,&user.Status,&user.Blocked,&user.EmailConfirmed,&user.Role,&user.CreatedAt,&user.UpdatedAt)
		users = append(users, user)
	}

	return users, nil
}