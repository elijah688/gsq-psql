package db

import (
	"database/sql"
	"fmt"
	"gql/graph/model"
	"gql/util"
	"log"
	"os"
	"strconv"
)

type Psql struct {
	Db *sql.DB
}

func (psql *Psql) Connect() {
	port, err := strconv.Atoi(os.Getenv("DBPORT"))
	util.CheckError(err)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), port, os.Getenv("DBUSER"), os.Getenv("PASS"), os.Getenv("NAME"))
	db, err := sql.Open("postgres", psqlconn)
	psql.Db = db

	util.CheckError(err)
	err = db.Ping()
	util.CheckError(err)
	log.Println("Connected!")
}

func (psql *Psql) Write(link *model.Link) {
	insertDynStmt := `insert into "links"("url", "name") values($1, $2)`
	_, err := psql.Db.Exec(insertDynStmt, (*link).URL, (*link).Name)
	util.CheckError(err)
}

func (psql *Psql) ReadAll() *[]*model.Link {
	rows, err := psql.Db.Query(`SELECT "url", "name" FROM "links"`)
	util.CheckError(err)
	defer rows.Close()

	var links []*model.Link
	log.Println(rows)
	for rows.Next() {
		var url string
		var name string

		err = rows.Scan(&name, &url)
		util.CheckError(err)

		links = append(links, &model.Link{Name: name, URL: url})
		log.Println(links)
	}
	util.CheckError(err)
	return &links
}
