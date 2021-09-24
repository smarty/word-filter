package prefilter

type treeNode struct {
	wordFragment string
	parent       *treeNode
	children     []*treeNode
	isWord       bool
}

func (this *treeNode) Add(word []byte) error {
	if len(word) == 0 {
		this.isWord = true
		return nil
	}

	wordFragment := word[0]

	//check if uppercase letter
	if !('a' <= wordFragment && wordFragment <= 'z') {
		//check if lowercase letter, change to upper
		if 'A' <= wordFragment && wordFragment <= 'Z' {
			wordFragment += 'a' - 'A'
		}
	}

	remainingWord := word[1:]
	for _, child := range this.children {
		if child.wordFragment == string(wordFragment) {
			return child.Add(remainingWord)
		}
	}

	child := &treeNode{
		wordFragment: string(wordFragment),
		parent:       this,
		children:     nil,
		isWord:       false,
	}
	if err := child.Add(remainingWord); err != nil {
		return err
	}

	this.children = append(this.children, child)
	return nil
}

func (this *treeNode) IsAllowed(input []byte) bool {
	// pass in whole string, if first letter doesn't match find next space, look at next letters

	return true
}

func (this *treeNode) isAllowedHelper(word []byte, index int) (bool, int) {

	if len(word) == 0 || word[0] == ' ' || word[0] == '\n' || word[0] == '\t' {
		if this.isWord == true {
			return false, 0
		}
		return true, index
	}

	//check if uppercase letter
	if !('a' <= word[0] && word[0] <= 'z') {
		//check if lowercase letter, change to upper
		if 'A' <= word[0] && word[0] <= 'Z' {
			word[0] += 'a' - 'A'
		} else {
			return true, index
		}
	}


	for _, child := range this.children {
		if string(word) == child.wordFragment {
			index += len(child.wordFragment)
			return child.isAllowedHelper(word[len(child.wordFragment):], index)
		}
	}
	return true, index
}
