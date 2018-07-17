package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // only run init
	"os"
	"zl/alis/src/utils"
	"log"
)

type Data struct {
	ID int64
	Context string
	Cmd string
	CmdType string
	Next int64
	ExtendInfo string
}

const createSQL = `
CREATE TABLE cmd(
   id INTEGER PRIMARY KEY        AUTOINCREMENT,
   context            CHAR(50)   NOT NULL,
   cmd                TEXT       NOT NULL,
   type               CHAR(50)   NOT NULL,
   next               INT        NOT NULL,
	 extendInfo         TEXT,
	 
	 unique(context, cmd, type)
);
`
var dbName = os.Getenv("HOME") + "/.alis.cmd.db"

func init() {
	log.Println("start to init db!")
	isExist, _ := utils.IsFileExist(dbName)

	if isExist { return }

	db, err := sql.Open("sqlite3", dbName)
	utils.HandleErr(err)
	defer db.Close()

	_, err = db.Exec(createSQL)
	utils.HandleErr(err)
}

func Create(d *Data) sql.Result {
	db, err := sql.Open("sqlite3", dbName)
	utils.HandleErr(err)
	defer db.Close()

	insertSQL := "INSERT INTO cmd(context, cmd, type, next, extendInfo) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Exec(insertSQL, d.Context, d.Cmd, d.CmdType, d.Next, d.ExtendInfo)
	utils.HandleErr(err)

	return result
}

func Find(d *Data) []*Data {
	db, err := sql.Open("sqlite3", dbName)
	utils.HandleErr(err)
	defer db.Close()

	findSQL := "select * from cmd where context=? and cmd=? and type=?;"
	rows, err := db.Query(findSQL, d.Context, d.Cmd, d.CmdType)
	defer rows.Close()

	utils.HandleErr(err)
	return convert2Data(rows)
}

func convert2Data (rows *sql.Rows) []*Data {
	var datas []*Data
	for rows.Next() {
		d := &Data{}
		err := rows.Scan(&d.ID, &d.Context, &d.Cmd, &d.CmdType, &d.Next, &d.ExtendInfo)
		utils.HandleErr(err)
		datas = append(datas, d)
	}
	return datas
}

func FindOne(d *Data) *Data {
	db, err := sql.Open("sqlite3", dbName)
	utils.HandleErr(err)
	defer db.Close()

	findSQL := "select * from cmd where id=?"
	rows, err := db.Query(findSQL, d.ID)
	defer rows.Close()

	utils.HandleErr(err)
	datas := convert2Data(rows)
	return datas[0]
}

func FindOriginOne(d *Data) *Data {
	result := FindOne(d)
	for {
		if result.Next == 0 {
			return result
		}
		result = FindOne(&Data{ ID: result.Next })
	}
}