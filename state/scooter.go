package main

type Scooter struct {
	BatteryOff State
	Lock       State
	Unlock     State

	CurrentState State

	currentState   string
	batteryPercent int
	isLock         bool
}

func newScooter(battery int, isLock bool) *Scooter {
	scooter := &Scooter{
		batteryPercent: battery,
		isLock:         isLock,
	}

	batteryOff := &BatteryOff{
		state:   "battery-off",
		Scooter: scooter,
	}
	lock := &Lock{
		state:   "lock",
		Scooter: scooter,
	}
	unlock := &Unlock{
		state:   "unlock",
		Scooter: scooter,
	}

	if battery <= 0 {
		scooter.setState(batteryOff)
	} else {
		scooter.setState(lock)
	}

	scooter.Lock = lock
	scooter.Unlock = unlock

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

func (s *Scooter) getState() string {
	return s.CurrentState.getState()
}

func (s *Scooter) setState(state State) {
	s.CurrentState = state
}

func (s *Scooter) getBattery() int {
	return s.batteryPercent
}

func (s *Scooter) getLockStatus() bool {
	return s.isLock
}
