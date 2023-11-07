package product

type ISeaWaterFish interface {
	wash(times int)
	setHowToCook(how string)
	setCustomer(customer string)
}

type SeaWaterFish struct {
	washTimes int
	howToCook string
	customer  string
}

type IFreshWaterFish interface {
	wash(times int)
	setHowToCook(how string)
	setCustomer(customer string)
}

type FreshWaterFish struct {
	washTimes int
	howToCook string
	customer  string
}
