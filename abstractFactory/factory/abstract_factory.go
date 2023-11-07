package factory

import (
	"errors"
	"goDesginPattern/abstractFactory/product"
)

type IFishCookingFactory interface {
	cookSeaWaterFish() product.ISeaWaterFish
	cookFreshWaterFish() product.IFreshWaterFish
}

func SelectCookFactory(waterType string) (IFishCookingFactory, error) {
	switch waterType {
	case "seaWater":

	case "freshWater":

	default:
		return nil, errors.New("waterType is wrong")
	}
}
