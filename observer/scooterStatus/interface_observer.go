package scooterStatus

type Observer interface {
	Update(battery int, coordinates Coordinates)
	GetProcessorName() string
}
