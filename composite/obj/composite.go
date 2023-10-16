package obj

import "fmt"

type Company struct {
	PMs     []PersonalMobility
	IoTName string
}

func (g *Company) Produce() {
	fmt.Printf("%s can produce below models.\n", g.IoTName)
	for _, composite := range g.PMs {
		composite.Produce()
	}
}

func (g *Company) AddPM(pm PersonalMobility) {
	g.PMs = append(g.PMs, pm)
}
