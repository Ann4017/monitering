package main

// "log"
// con "monitering/config"
// "monitering/db"

func main() {
	// //ini 존재 유무 확인 및 생성
	// err := con.Check_ini("config/config.ini")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //database 구성
	// db_con := &db.DB_info{}
	// err = db_con.Cfg()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //mysql 연동
	// // err = db_con.DB_con()
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// err = db_con.Moniter()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	go func() {
	// 		_, err := hw.Get_cpu_model()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	go func() {
	// 		_, err := hw.Get_cpu_total_usage()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	go func() {
	// 		_, err := hw.Get_cpu_core_usage()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// go func() {
	// 	_, err := hw.Get_cpu_temps()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// 	go func() {
	// 		_, err := hw.Get_mem_info()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	go func() {
	// 		_, err := hw.Get_disk_info()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	time.Sleep(time.Second * 5)
	// }

	// for {
	// 	status, err := http.Get_http_status("https://www.naver.com/")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(status)

	// 	ping, err := http.Get_http_ping("https://www.naver.com/")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(ping)

	// 	certs, err := http.Get_https_cert("naver.com")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	for _, cert := range certs {
	// 		fmt.Printf("Expiration date: %s\n", cert.NotAfter.Format(time.RFC3339))
	// 	}

	// 	time.Sleep(time.Second * 5)
	// }

}
