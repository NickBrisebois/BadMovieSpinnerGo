package data

type DrawProperties struct {
	Step       int
	StartAngle float32
	EndAngle   float32
}

func NewDrawProperties() *DrawProperties {
	return &DrawProperties{}
}

func GetNextStateSliceDrawProperties(currentDrawProperties *DrawProperties) *DrawProperties {
	return nil
}

type SpinnerState struct {
}
