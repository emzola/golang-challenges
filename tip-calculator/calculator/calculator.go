package calculator

func Calculate(bill, tipRate uint64) float64 {
	return float64(bill) * (float64(tipRate) / 100)
}