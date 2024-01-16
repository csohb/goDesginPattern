package main

import "fmt"

func main() {
	battery := 19
	isLock := true
	scooter := newScooter(battery, isLock)

	fmt.Println("currentState:", scooter.CurrentState.getState())

	// 초기화 시 기본 상태 == lock
	err := scooter.CurrentState.locking()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = scooter.CurrentState.riding()
	if err != nil {
		fmt.Println(err.Error())
	}

	// unlocking Scooter Current State : Unlock
	err = scooter.CurrentState.unlocking()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("currentState:", scooter.CurrentState.getState())

	// riding 시도 -> 배터리 이슈로 불가
	err = scooter.CurrentState.riding()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = scooter.CurrentState.chargeBattery(30)
	if err != nil {
		fmt.Println(err.Error())
	}

	// riding 완료
	err = scooter.CurrentState.riding()
	if err != nil {
		fmt.Println(err.Error())
	}

	// lock 완료
	err = scooter.CurrentState.locking()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("currentState:", scooter.CurrentState.getState())

}
