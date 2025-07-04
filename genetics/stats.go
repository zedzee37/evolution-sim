package genetics

const (
	HealthStat = iota
	WalkSpeedStat
	SwimSpeedStat
	AirRespirationStat
	WaterRespirationStat
	ScaleStat
	HungerStat
	ThirstStat
	StaminaStat
	GestationPeriod
	AgeStat
	LifeSpanStat
	AdultAgeStat
	VisionStat
	AttackStat
	AggressionStat	
	StatCount
)

type Stats [StatCount]float64
