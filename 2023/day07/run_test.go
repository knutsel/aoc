package day07

import "testing"

func TestReal(t *testing.T) {
	// Run("example.txt")
	Run("input.txt")
}

func TestRanking(t *testing.T) {

	for _, v := range []string{"JAAJA", "AAAJA", "JJJJJ"} {
		s := scoreHandP2(v, 0)
		t.Logf("%s -> 0x%x (%d)", v, s, s)
	}
}
