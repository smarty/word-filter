package prefilter

import "strings"

func main() {
}

type Prefilter interface {
	IsAllowed(string) bool
}

type prefilter struct {
	buffer     []byte
	restricted []string
	Prefilter
}

func newPrefilter(bufferSize int, reserved ...string) Prefilter {
	b := make([]byte, bufferSize, bufferSize)
	m := make(map[string]struct{})
	for _, r := range reserved {
		m[r] = struct{}{}
	}
	restricted := make([]string, 0, len(m))
	for key, _ := range m {
		restricted = append(restricted, key)
	}
	return &prefilter{
		buffer:     b,
		restricted: restricted,
	}
}

func (this *prefilter) IsAllowed(input string) bool {
	copy(this.buffer[:], input)
	start := 0
	for end, r := range this.buffer {
		//check if lowercase letter
		if !('a' <= r && r <= 'z') {
			//check if uppercase letter, change to lower
			if 'A' <= r && r <= 'Z' {
				this.buffer[end] += 'a' - 'A'
			}
		}

		if r == ' ' || r == '\n' || r == '\t' {
			word := this.buffer[start:end]
			if contains(this.restricted, word) {
				return false
			}
			start = end + 1
		}
	}

	// restricted word may be at eof
	word := this.buffer[start:]
	if contains(this.restricted, word) {
		return false
	}

	return true
}

func contains(restricted []string, word []byte) bool {
	for _, s := range restricted {
		if s == string(word) {
			return true
		}
	}

	return false
}

func (this *prefilter) IsAllowed2(input string) bool {
	input = strings.ToUpper(input)
	for _, word := range this.restricted {
		if strings.Contains(input, word) {
			return false
		}
	}
	return true
}
