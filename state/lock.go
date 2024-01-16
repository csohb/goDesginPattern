package main

import (
	"errors"
	"fmt"
)

type Lock struct {
	state   string
	Scooter *Scooter
}

func (l *Lock) chargeBattery(battery int) error {
	fmt.Printf("battery charged to - %d\n", battery)
	l.Scooter.batteryPercent = battery
	return nil
}

func (l *Lock) locking() error {
	return errors.New("scooter is already locked")
}

func (l *Lock) unlocking() error {
	fmt.Println("scooter is unlocking")
	l.Scooter.setState(l.Scooter.Unlock)
	l.Scooter.isLock = false
	return nil
}

func (l *Lock) riding() error {
	return errors.New("cannot ride while scooter is locked")
}

func (l *Lock) getState() string {
	return l.state
}
