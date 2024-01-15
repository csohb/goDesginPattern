package main

type Scooter struct {
	BatteryOff State
	BatteryOn  State
	Lock       State
	Unlock     State

	CurrentState State

	batteryPercent int
}

func newScooter(battery int) *Scooter {
	scooter := &Scooter{
		batteryPercent: battery,
	}
	return scooter
}

func (s *Scooter) chargeBattery(battery int) error {
	return s.CurrentState.chargeBattery(battery)
}
func (s *Scooter) locking() error {
	return s.CurrentState.locking()
}
func (s *Scooter) unlocking() error {
	return s.CurrentState.unlocking()
}
func (s *Scooter) riding() error {
	return s.CurrentState.riding()
}
