package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func Check_ini(inipath string) error {
	if _, err := os.Stat(inipath); os.IsNotExist(err) {
		file := ini.Empty()

		d_sec := file.Section("database")
		d_sec.NewKey("user", "root")
		d_sec.NewKey("pwd", "RLGH3qjs!!")
		d_sec.NewKey("host", "localhost")
		d_sec.NewKey("port", "3306")
		d_sec.NewKey("database", "addrbook")
		d_sec.NewKey("engine", "mysql")

		err = file.SaveTo(inipath)
		if err != nil {
			return err
		}

		fmt.Println("Create a new ini file")
	} else {
		fmt.Println("ini file already exists")
	}
	return nil
}
