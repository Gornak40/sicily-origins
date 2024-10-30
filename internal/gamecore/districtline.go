package gamecore

type DistrictCard int

func (d DistrictCard) Card() Card {
	return Card(d & 0x1ff)
}

func (d DistrictCard) Closed() bool {
	return (d >> 3 & 1) == 1
}

type DistrictLine struct {
}
