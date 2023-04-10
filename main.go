package main

import (
	"fmt"
	"log"
	"monitering/hw"
	"time"
)

type hw_info struct {
	c hw.CPU_info
	m hw.MEM_cap
	d hw.DISK_cap
}

func main() {
	hw := &hw_info{}
	go func() {
		for {
			cpu_m, err := hw.c.Get_cpu_model()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("cpu model: %s\n", cpu_m)
			fmt.Printf("cpu model: %s\n", hw.c.S_cpu_model)

			cpu_t_usage, err := hw.c.Get_cpu_total_usage()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("cpu total usage: %.2f%%\n", cpu_t_usage[0])
			fmt.Printf("cpu total usage: %.2f%%\n", hw.c.F_cpu_t_usage[0])

			cpu_c_usage, err := hw.c.Get_cpu_core_usage()
			if err != nil {
				log.Fatal(err)
			}
			for i, usage := range cpu_c_usage {
				fmt.Printf("core #%d usage: %.2f%%\n", i, usage)
			}
			for i, usage := range hw.c.F_cpu_c_usage {
				fmt.Printf("core #%d usage: %.2f%%\n", i, usage)
			}

			time.Sleep(time.Second * 5)
		}
	}()
	select {}
}
