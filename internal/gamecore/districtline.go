package gamecore

type DistrictCard int8

func (d DistrictCard) Card() Card {
	return Card(d & 7)
}

func (d DistrictCard) Closed() bool {
	return (d >> 3 & 1) == 1
}

type DistrictLine []DistrictCard
