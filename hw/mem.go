package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

type MEM_cap struct {
	F_mem_total   float64
	F_mem_used    float64
	F_mem_free    float64
	F_mem_usedper float64
}

func (m *MEM_cap) Get_mem_info() (*mem.VirtualMemoryStat, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	total, used, free := float64(vm.Total)/(1024*1024*1024), float64(vm.Used)/(1024*1024*1024), float64(vm.Free)/(1024*1024*1024)

	fmt.Printf("total memory: %.1fGB\n", total)
	fmt.Printf("memory used: %.1fGB\n", used)
	fmt.Printf("memory free: %.1fGB\n", free)
	fmt.Printf("memory usedpercent: %.1f%%\n", vm.UsedPercent)

	m.F_mem_total = total
	m.F_mem_used = used
	m.F_mem_free = free
	m.F_mem_usedper = vm.UsedPercent

	return vm, nil
}
