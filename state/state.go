package main

type State interface {
	getState() string
	chargeBattery(battery int) error
	locking() error
	unlocking() error
	riding() error
}
