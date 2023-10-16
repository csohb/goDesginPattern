package obj

import "fmt"

type Scooter struct {
	ModelName string
}

func (s *Scooter) Produce() {
	fmt.Printf("Produced Scooter ModelName:%s\n", s.ModelName)
}

func (s *Scooter) GetScooterName() string {
	return s.ModelName
}
