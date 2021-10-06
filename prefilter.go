package prefilter

type Prefilter interface {
	IsAllowed(value string) bool
}

type treeNode struct {
	wordFragmentUpper byte
	wordFragmentLower byte
	isWord            bool
	children          []*treeNode
}

func New(reserved ...string) Prefilter {
	this := &treeNode{}

	for _, item := range reserved {
		this.add(item)
	}

	return this
}

func (this *treeNode) add(value string) {
	if len(value) == 0 {
		this.isWord = true
		return
	}

	wordFragment := value[0]

	// check if lowercase letter
	if !('a' <= wordFragment && wordFragment <= 'z') {
		// check if uppercase letter, change to lower
		if 'A' <= wordFragment && wordFragment <= 'Z' {
			wordFragment += 'a' - 'A'
		}
	}

	remainingWord := value[1:]
	for _, child := range this.children {
		if child.wordFragmentLower == wordFragment {
			child.add(remainingWord)
			return
		}
	}

	child := &treeNode{
		wordFragmentLower: wordFragment,
		wordFragmentUpper: wordFragment - ('a' - 'A'),
		children:          nil,
		isWord:            false,
	}

	this.children = append(this.children, child)
	child.add(remainingWord)
}

func (this *treeNode) IsAllowed(input string) bool {
	allowed := true
	raw := []byte(input)
	inputLength := len(input)

	for index := 0; index < inputLength; index++ {
		if allowed, index = this.isAllowed(raw, index); !allowed {
			return false
		}
	}

	return true
}
func (this *treeNode) isAllowed(input []byte, index int) (bool, int) {
	if len(input) == index || input[index] == ' ' || input[index] == '\n' || input[index] == '\t' {
		if this.isWord == true {
			return false, index
		}

		return true, index
	}

	for _, child := range this.children {
		if input[index] == child.wordFragmentLower {
			return child.isAllowed(input, index+1)

		} else if input[index] == child.wordFragmentUpper {
			return child.isAllowed(input, index+1)

		} else {
			for input[index] != ' ' && input[index] != '\n' && input[index] != '\t' && len(input) != index+1 {
				index += 1
			}

			return true, index
		}
	}

	return true, index
}
