package hw

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

type CPU_info struct {
	S_cpu_model   string
	F_cpu_t_usage []float64
	F_cpu_c_usage []float64
	F_cpu_temps   []host.TemperatureStat
}

func (c *CPU_info) Get_cpu_model() (cpu_model string, err error) {
	info, err := cpu.Info()
	if err != nil {
		return "", err
	}
	model := info[0].ModelName
	fmt.Printf("cpu model: %s\n", model)
	c.S_cpu_model = model
	return model, nil
}

func (c *CPU_info) Get_cpu_total_usage() (cpu_t_usage []float64, err error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	fmt.Printf("CPU total usage: %.2f%%\n", percent[0])
	c.F_cpu_t_usage = percent
	return percent, nil
}

func (c *CPU_info) Get_cpu_core_usage() (cpu_c_usage []float64, err error) {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}

	for i, usage := range percent {
		fmt.Printf("cpu core #%d usage: %.2f%%\n", i, usage)
	}
	c.F_cpu_c_usage = percent
	return percent, nil
}

func (c *CPU_info) Get_cpu_temps() (temps []host.TemperatureStat, err error) { //mac,linux 에서 만 가능
	temps, err = host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	for i, temp := range temps {
		if temp.SensorKey == "coretemp" {
			fmt.Printf("core #%d: %.1f°C\n", i, temp.Temperature)
		}
	}
	c.F_cpu_temps = temps
	return temps, nil
}
