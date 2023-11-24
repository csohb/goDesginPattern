package scooterStatus

import (
	"fmt"
)

type Processor struct {
	Name string
}

func (p *Processor) Update(battery int, coordinates Coordinates) {
	fmt.Printf("%sProcessor got information :: scooter Status has updated!!\nbattery:%d, Coordinates:{lng:%f,lat:%f}\n", p.Name, battery, coordinates.Longitude, coordinates.Latitude)
}

func (p *Processor) GetProcessorName() string {
	return p.Name
}
