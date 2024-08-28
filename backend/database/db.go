package database

import (
	"database/sql"
	"fmt"
	"forum/backend/config"
	q "forum/backend/database/query"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
	mu sync.Mutex
	Test string
}

const (
	USERS_TABLE      = "users"
	POSTS_TABLE      = "posts"
	CATEGORIES_TABLE = "categories"
	CHATS_TABLE      = "chats"
	COMMENTS_TABLE   = "comments"
	FOLLOWS_TABLE    = "follows"
	REACTIONS_TABLE  = "reactions"
	SESSIONS_TABLE   = "sessions"
	MESSAGES_TABLE   = "messages"
	USERCHATS_TABLE  = "users_chats"
	CAT_POST_TABLE   = "cats_posts"
)

var (
	DB *Database
)

func init() {
	filepath := config.Get("DATABASE_FILEPATH").ToString()
	sqlPath := config.Get("SQL_CREATE_PATH").ToString()
	insertCatSqlPath := config.Get("SQL_INSERT_CATS_PATH").ToString()
	insertCatFile,err :=  os.ReadFile(insertCatSqlPath)
	if err != nil {
		fmt.Println("cannot open the sql file")
		os.Exit(1)
	}
	sqlFile, err := os.ReadFile(sqlPath)
	if err != nil {
		fmt.Println("cannot open the sql file")
		os.Exit(1)
	}
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		fmt.Println("cannot open the database file")
		os.Exit(1)
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.Exec(string(insertCatFile))
	DB = &Database{DB: db}
}

func (d *Database) Insert(table string, data any) error {

	query,err := q.InsertQuery(table, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	prep, err := d.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = prep.Exec()

	return err
}

func (d *Database) Delete(table string, where q.WhereOption) error {

	query := q.DeleteQuery(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) Update(table string, object any, where q.WhereOption) error {
	var err error
	query,err := q.UpdateQuery(table, object, where)
	if err != nil {
		fmt.Println(err)
		return err
	}
	stmt, err := d.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) GetOneFrom(table string, where q.WhereOption,) (*sql.Row, error) {

	query := q.SelectOneFrom(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()

	return row, nil
}

func (d *Database) GetAllFrom(table string, where q.WhereOption,orderby string,limit []int) (*sql.Rows, error) {
	var query string
	if where == nil {
		query = q.SelectAllFrom(table,orderby,limit)
	} else {
		query = q.SelectAllWhere(table, where,orderby,limit)
	}

	stmt, err := d.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetAllAndJoin(table string, j []q.JoinCondition, where q.WhereOption,orderby string,limit []int) (*sql.Rows, error) {
	var query = q.SelectWithJoinQuery(table, j, where,orderby,limit)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetCount(table string, where q.WhereOption) (*sql.Row, error) {
	var query = q.GetCountQuery(table, where)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()

	if err != nil {
		return nil, err
	}
	return row, nil
}
