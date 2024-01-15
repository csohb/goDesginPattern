package main

type State interface {
	chargeBattery(battery int) error
	locking() error
	unlocking() error
	riding() error
}
