package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func Get_disk_info() (*disk.UsageStat, error) {
	disk, err := disk.Usage("C:")
	if err != nil {
		return nil, err
	}

	total, used, free := float64(disk.Total)/(1024*1024*1024), float64(disk.Used)/(1024*1024*1024), float64(disk.Free)/(1024*1024*1024)

	fmt.Printf("disk total cap: %.2fGB\n", total)
	fmt.Printf("disk used cap: %.2fGB\n", used)
	fmt.Printf("disk free cap: %.2fGB\n", free)

	return disk, nil
}
