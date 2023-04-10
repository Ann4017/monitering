package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

type DISK_cap struct {
	F_disk_total float64
	F_disk_used  float64
	F_disk_free  float64
}

func (d *DISK_cap) Get_disk_info() (disk_cap *disk.UsageStat, err error) {
	disk, err := disk.Usage("C:")
	if err != nil {
		return nil, err
	}

	total, used, free := float64(disk.Total)/(1024*1024*1024), float64(disk.Used)/(1024*1024*1024), float64(disk.Free)/(1024*1024*1024)

	fmt.Printf("disk total cap: %.2fGB\n", total)
	fmt.Printf("disk used cap: %.2fGB\n", used)
	fmt.Printf("disk free cap: %.2fGB\n", free)

	d.F_disk_total = total
	d.F_disk_used = used
	d.F_disk_free = free

	return disk, nil
}
