package prefilter

type treeNode struct {
	wordFragmentUpper byte
	wordFragmentLower byte
	children          []*treeNode
	isWord            bool
}

func (this *treeNode) Add(word []byte) error {
	if len(word) == 0 {
		this.isWord = true
		return nil
	}

	wordFragment := word[0]

	//check if lowercase letter
	if !('a' <= wordFragment && wordFragment <= 'z') {
		//check if uppercase letter, change to lower
		if 'A' <= wordFragment && wordFragment <= 'Z' {
			wordFragment += 'a' - 'A'
		}
	}

	remainingWord := word[1:]
	for _, child := range this.children {
		if child.wordFragmentLower == wordFragment {
			return child.Add(remainingWord)
		}
	}

	child := &treeNode{
		wordFragmentLower: wordFragment,
		wordFragmentUpper: wordFragment - ('a' - 'A'),
		children:          nil,
		isWord:            false,
	}
	if err := child.Add(remainingWord); err != nil {
		return err
	}

	this.children = append(this.children, child)
	return nil
}

func (this *treeNode) IsAllowed(input []byte) bool {
	index := 0
	allowed := true

	inputLength := len(input)

	for index < inputLength {
		if allowed, index = this.isAllowedHelper(input, index); allowed == false {
			return false
		}
		index += 1
	}
	return true
}

func (this *treeNode) isAllowedHelper(input []byte, index int) (bool, int) {
	if len(input) == index || input[index] == ' ' || input[index] == '\n' || input[index] == '\t' {
		if this.isWord == true {
			return false, index
		}
		return true, index
	}

	for _, child := range this.children {
		if input[index] == child.wordFragmentLower {
			return child.isAllowedHelper(input, index+1)
		} else if input[index] == child.wordFragmentUpper {
			return child.isAllowedHelper(input, index+1)
		} else {
			for input[index] != ' ' && input[index] != '\n' && input[index] != '\t' && len(input) != index+1 {
				index += 1
			}
			return true, index
		}
	}
	return true, index
}
