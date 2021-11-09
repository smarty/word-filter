package wordfilter

import "testing"

func TestTreePrefiltering(t *testing.T) {
	assertTreePrefilter(t, "WORD", false, "word")
	assertTreePrefilter(t, "word", false, "word")
	assertTreePrefilter(t, "woRd", false, "word")

	assertTreePrefilter(t, "A sentence with a word in the middle", false, "word", "restricted")
	assertTreePrefilter(t, "A sentence with a word\n in the middle", false, "word", "restricted")
	assertTreePrefilter(t, "A sentence with a word	 in the middle", false, "word", "restricted")

	assertTreePrefilter(t, "Only good words", true, "word", "restricted")
	assertTreePrefilter(t, "A good sentence 123 with numbers", true, "word", "restricted")
	assertTreePrefilter(t, "Trailing space ", true, "word", "restricted")

	assertTreePrefilter(t, "A bad sentence with numbers word1", false, "word1", "restricted")
}

func assertTreePrefilter(t *testing.T, input string, expected bool, reserved ...string) {
	t.Helper()

	reservedTree := New(reserved...)
	actual := reservedTree.IsAllowed(input)

	if actual != expected {
		t.Errorf("\nExpected: %t\nActual:   %t\n", expected, actual)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BenchmarkTreeTest(b *testing.B) {
	b.ReportAllocs()
	// 256 chars
	stringToCheck := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
		"labore et dolore magna aliqua. Ut EMIM ad minim veniam, quis nostrud exercitation ULLAMCO laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure d RANDOM"
	reservedTree := New("ANOTHER", "RANDOM", "WORD")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = reservedTree.IsAllowed(stringToCheck)
	}
}
