package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func Get_mem_info() (*mem.VirtualMemoryStat, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	total, used, free := float64(vm.Total)/(1024*1024*1024), float64(vm.Used)/(1024*1024*1024), float64(vm.Free)/(1024*1024*1024)

	fmt.Printf("total memory: %.1fGB\n", total)
	fmt.Printf("memory used: %.1fGB\n", used)
	fmt.Printf("memory free: %.1fGB\n", free)
	fmt.Printf("memory usedpercent: %.1f%%\n", vm.UsedPercent)
	return vm, nil
}
