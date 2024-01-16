package main

import (
	"errors"
	"fmt"
)

type BatteryOff struct {
	state   string
	Scooter *Scooter
}

func (b *BatteryOff) chargeBattery(battery int) error {
	fmt.Printf("battery charged to - %d\n", battery)
	b.Scooter.batteryPercent = battery
	return nil
}

func (b *BatteryOff) locking() error {
	return errors.New("cannot lock while battery off")
}

func (b *BatteryOff) unlocking() error {
	return errors.New("cannot unlock while battery off")
}

func (b *BatteryOff) riding() error {
	return errors.New("cannot ride while battery off")
}

func (b *BatteryOff) getState() string {
	return b.state
}
