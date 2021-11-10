package wordfilter

import "testing"

func TestTreeFiltering(t *testing.T) {
	assertContains(t, "WORD", true, "word")
	assertContains(t, "word", true, "word")
	assertContains(t, "woRd", true, "word")

	assertContains(t, "A sentence with a word in the middle", true, "word", "restricted")
	assertContains(t, "A sentence with a word\n in the middle", true, "word", "restricted")
	assertContains(t, "A sentence with a word	 in the middle", true, "word", "restricted")

	assertContains(t, "Only good words", false, "word", "restricted")
	assertContains(t, "A good sentence 123 with numbers", false, "word", "restricted")
	assertContains(t, "Trailing space ", false, "word", "restricted")

	assertContains(t, "A bad sentence with numbers word1", true, "word1", "restricted")
}
func assertContains(t *testing.T, input string, expected bool, reserved ...string) {
	t.Helper()

	tree := New(reserved...)
	actual := tree.Contains(input)

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
		_ = reservedTree.Contains(stringToCheck)
	}
}
