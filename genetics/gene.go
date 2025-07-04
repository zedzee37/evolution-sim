package genetics

const (
	RespirationTypeGenoType = iota
	LegCountGenoType
	LifeSpanGenoType
	GestationLengthTGenoType
	EyeCountGenoType		
	FoodPreferenceGenoType
	ScaleGenoType
	GenoTypeCount
)

type Gene struct {
	GenoType int
	Name string
	Modifier Stats
}
