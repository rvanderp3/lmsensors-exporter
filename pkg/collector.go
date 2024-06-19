package pkg

import (
	"fmt"
	"github.com/mdlayher/lmsensors"
	"log"
)

var (
	scanner *lmsensors.Scanner
	devices []*lmsensors.Device
)

func init() {
	scanner = lmsensors.New()

	lmdevices, err := scanner.Scan()
	if err != nil {
		log.Fatal(fmt.Errorf("error scanning devices: %w", err))
	}

	for _, device := range lmdevices {
		devices = append(devices, device)
	}
}

func collect() {
	for _, device := range devices {
		for _, sensor := range device.Sensors {
			tempSensor, ok := sensor.(*lmsensors.TemperatureSensor)
			if ok {
				reportGuageMetric(fmt.Sprintf("%s-%s", device.Name, tempSensor.Name), int64(tempSensor.Input))
				continue
			}
			fanSensor, ok := sensor.(*lmsensors.FanSensor)
			if ok {
				reportGuageMetric(fmt.Sprintf("%s-%s", device.Name, fanSensor.Name), int64(fanSensor.Input))
				continue
			}
		}
	}
}
