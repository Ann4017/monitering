package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

type DB_info struct {
	s_user     string
	s_pwd      string
	s_host     string
	s_port     string
	s_database string
	s_engine   string
}

func (d *DB_info) DB_cfg() error {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		return err
	}
	d.s_user = file.Section("database").Key("user").String()
	d.s_pwd = file.Section("database").Key("pwd").String()
	d.s_host = file.Section("database").Key("host").String()
	d.s_port = file.Section("database").Key("port").String()
	d.s_database = file.Section("database").Key("database").String()
	d.s_engine = file.Section("database").Key("engine").String()

	return nil
}

func (d *DB_info) DB_con() error {
	db_source := d.s_user + ":" + d.s_pwd + "@tcp(" + d.s_host + ":" + d.s_port + ")/" + d.s_database
	db, err := sql.Open(d.s_engine, db_source)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func (d *DB_info) DB_moniter() error {
	db_source := d.s_user + ":" + d.s_pwd + "@tcp(" + d.s_host + ":" + d.s_port + ")/" + d.s_database
	db, err := sql.Open(d.s_engine, db_source)
	if err != nil {
		return err
	}
	defer db.Close()

	for {
		rows, err := db.Query("select * from info where id <= 2")
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			var age int

			if err := rows.Scan(&id, &name, &age); err != nil {
				return err
			}

			log.Printf("name: %s, age: %d\n", name, age)
			time.Sleep(time.Second)
		}
	}
}
