package prefilter

import (
	"testing"
)

func TestPrefiltering(t *testing.T) {
	assertPrefilter(t, "WORD", false, "word")
	assertPrefilter(t, "word", false, "word")
	assertPrefilter(t, "Word", false, "word")
	assertPrefilter(t, "woRd", false, "word")

	assertPrefilter(t, "word", false, "word", "restricted")
	assertPrefilter(t, "restricted", false, "word", "restricted")
	assertPrefilter(t, "A sentence with a word in the middle", false, "word", "restricted")
	assertPrefilter(t, "A sentence with a word\n in the middle", false, "word", "restricted")
	assertPrefilter(t, "A sentence with a word	 in the middle", false, "word", "restricted")

	assertPrefilter(t, "Only good words", true, "word", "restricted")
	assertPrefilter(t, "A good sentence 123 with numbers", true, "word", "restricted")
	assertPrefilter(t, "Trailing space ", true, "word", "restricted")

	assertPrefilter(t, "A bad sentence with numbers word1", false, "word1", "restricted")
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
		"labore et dolore magna aliqua. Ut EMIM ad minim veniam, quis nostrud exercitation ULLAMCO laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure d RANDOM"
	reservedWords := newPrefilter(len(stringToCheck), "ANOTHER", "RANDOM", "WORD")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = reservedWords.IsAllowed(stringToCheck)
	}
}