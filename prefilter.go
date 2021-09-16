package prefilter

func main() {
}

type Prefilter interface {
	IsAllowed(string) bool
}

type prefilter struct {
	buffer     []byte
	restricted map[string]bool
	Prefilter
}

func newPrefilter(bufferSize int, reserved ...string) Prefilter {
	b := make([]byte, 0, bufferSize)
	m := make(map[string]bool)
	for _, r := range reserved {
		m[r] = true
	}
	return &prefilter{
		buffer:     b,
		restricted: m,
	}
}

func (this *prefilter) IsAllowed(input string) bool {
	this.buffer = []byte(input)
	start := 0
	for end, r := range this.buffer {
		//check if uppercase letter
		if !('A' <= r && r <= 'Z') {
			//check if lowercase letter, change to upper
			if 'a' <= r && r <= 'z' {
				this.buffer[end] -= 'a' - 'A'
			}
		}

		if r == ' ' {
			word := this.buffer[start:end]
			_, restricted := this.restricted[string(word)] //FIXME does this allocate?
			if restricted {
				return false
			}
			start = end + 1
		}
	}
	return true
}
