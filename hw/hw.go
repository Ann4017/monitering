package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

func Get_cpu_total_usage() ([]float64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	fmt.Printf("CPU total usage: %.2f%%\n", percent[0])
	return percent, nil
}

func Get_cpu_core_usage() ([]float64, error) {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}

	for i, usage := range percent {
		fmt.Printf("cpu core #%d usage: %.2f%%\n", i+1, usage)
	}
	return percent, nil
}

func Get_cpu_temps() ([]host.TemperatureStat, error) {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	for i, temp := range temps {
		if temp.SensorKey == "coretemp" {
			fmt.Printf("core #%d: %.1f°C\n", i+1, temp.Temperature)
		}
	}
	return temps, nil
}
