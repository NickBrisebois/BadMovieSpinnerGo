package swidgetutils

func CalculatePercentOf(p int, total int) float64 {
	fPercent := float64(p)
	fTotal := float64(total)
	return fPercent / 100.0 * fTotal
}
