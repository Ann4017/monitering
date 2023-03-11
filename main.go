package main

import (
	"log"
	con "monitering/config"
	"monitering/db"
)

func main() {
	//ini 존재 유무 확인 및 생성
	err := con.Check_ini("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	//database 구성
	db_con := &db.DB_info{}
	err = db_con.DB_cfg()
	if err != nil {
		log.Fatal(err)
	}

	//mysql 연동
	err = db_con.DB_con()
	if err != nil {
		log.Fatal(err)
	}

	err = db_con.DB_moniter()
	if err != nil {
		log.Fatal(err)
	}
}
