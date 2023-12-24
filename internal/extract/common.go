package extract

func MuitoDiferente(a, b uint64) bool {
	if a > b {
		return a-b > 10
	}
	return b-a > 10
}
