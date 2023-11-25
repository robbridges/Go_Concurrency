package channels

import ("testing")

func TestCountConcurrently(t *testing.T) {
	results := CountConcurrently(1)
	

	counts := make(map[string]int)

	for _, result := range results {
		counts[result]++
	}
	if counts["fish"] != 5 {
        t.Errorf("Expected 5 'fish', but got %d", counts["fish"])
    }

    if counts["sheep"] != 5 {
        t.Errorf("Expected 5 'sheep', but got %d", counts["sheep"])
    }
}