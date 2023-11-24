package main

import "fmt"

func main() {
	careTaker := &CareTaker{mementoArray: make([]*Memento, 0)}

	originator := &Originator{speed: 25}
	fmt.Printf("Trip Started With Speed : %d\n", originator.getSpeed())
	careTaker.addMemento(originator.createMemento())

	fmt.Println("Entered Geofence Zone (15km/h)")
	originator.setSpeed(15)
	fmt.Printf("Scooter has entered geofence area with speed : %d\n", originator.getSpeed())
	careTaker.addMemento(originator.createMemento())

	fmt.Println("Outside Geofence Zone")
	originator.restoreMemento(careTaker.getMemento(0))
	fmt.Printf("Scooter is outside of geofence area with speed : %d\n", originator.getSpeed())

}
