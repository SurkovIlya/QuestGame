package globalF

import "math/rand"

func CheckAgil(agil int) int {
	youLuck := rand.Intn(24)

	return youLuck + agil
}

func CheckInt(intel int) int {
	youLuck := rand.Intn(24)

	return youLuck + intel
}
