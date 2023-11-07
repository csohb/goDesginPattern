package factory

type KingCrab struct {
}

func (k *KingCrab) cookKingCrab() ISeaWaterFactory {
	return &KingCrab{}
}
