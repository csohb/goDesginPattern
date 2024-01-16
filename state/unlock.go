package main

import (
	"errors"
	"fmt"
)

type Unlock struct {
	state   string
	Scooter *Scooter
}

func (u *Unlock) chargeBattery(battery int) error {
	fmt.Printf("battery charged to - %d\n", battery)
	u.Scooter.batteryPercent = battery
	return nil
}

func (u *Unlock) locking() error {
	fmt.Println("scooter is locking")
	u.Scooter.setState(u.Scooter.Lock)
	u.Scooter.isLock = true
	return nil
}

func (u *Unlock) unlocking() error {
	return errors.New("scooter is already unlocked")
}

func (u *Unlock) riding() error {
	if u.Scooter.batteryPercent >= 20 && u.Scooter.isLock == false {
		fmt.Println("scooter can ride!")
		return nil
	} else {
		return errors.New("scooter cannot ride while battery is under 20%")
	}
}

func (u *Unlock) getState() string {
	return u.state
}
