package data

type DrawProperties struct {
	Step  int
	Start float32
	End   float32
}

func NewDrawProperties() *DrawProperties {
	return &DrawProperties{}
}

func GetNextStateSliceDrawProperties(currentDrawProperties *DrawProperties) *DrawProperties {
	return nil
}

type SpinnerState struct {
}
