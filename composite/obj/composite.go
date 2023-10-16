package obj

import "fmt"

type IoT struct {
	PMs     []PersonalMobility
	IoTName string
}

func (g *IoT) Produce() {
	fmt.Printf("%s can produce below models.\n", g.IoTName)
	for _, composite := range g.PMs {
		composite.Produce()
	}
}

func (g *IoT) AddIot(iot PersonalMobility) {
	g.PMs = append(g.PMs, iot)
}
