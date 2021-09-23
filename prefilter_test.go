package prefilter

import (
	"testing"
)

func TestPrefiltering(t *testing.T) {
	assertPrefilter(t, "WORD", false, "WORD")
	assertPrefilter(t, "word", false, "WORD")
	assertPrefilter(t, "Word", false, "WORD")
	assertPrefilter(t, "woRd", false, "WORD")

	assertPrefilter(t, "word", false, "WORD", "RESTRICTED")
	assertPrefilter(t, "restricted", false, "WORD", "RESTRICTED")
	assertPrefilter(t, "A sentence with a word in the middle", false, "WORD", "RESTRICTED")
	assertPrefilter(t, "A sentence with a word\n in the middle", false, "WORD", "RESTRICTED")
	assertPrefilter(t, "A sentence with a word	 in the middle", false, "WORD", "RESTRICTED")

	assertPrefilter(t, "Only good words", true, "WORD", "RESTRICTED")
	assertPrefilter(t, "A good sentence with numbers 1234", true, "WORD", "RESTRICTED")
	assertPrefilter(t, "A good sentence 123 with numbers", true, "WORD", "RESTRICTED")
	assertPrefilter(t, "Trailing space ", true, "WORD", "RESTRICTED")

	assertPrefilter(t, "A bad sentence with numbers word1", false, "WORD1", "RESTRICTED")
}

func assertPrefilter(t *testing.T, input string, expected bool, reserved ...string) {
	t.Helper()

	reservedWords := newPrefilter(len(input), reserved...)
	actual := reservedWords.IsAllowed(input)

	if actual == expected {
		return
	}

	t.Errorf("\n"+
		"Expected: %t\n"+
		"Actual:   %t\n",
		expected,
		actual,
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BenchmarkTest(b *testing.B) {
	b.ReportAllocs()
	// 256 chars
	stringToCheck := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
		"labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure d RANDOM"
	reservedWords := newPrefilter(len(stringToCheck), "ANOTHER", "RANDOM", "WORD")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = reservedWords.IsAllowed(stringToCheck)
	}
}