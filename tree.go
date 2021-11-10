package wordfilter

type Filter interface {
	Contains(value string) bool
}

type treeNode struct {
	uppercase byte
	lowercase byte
	isWord    bool
	children  []*treeNode
}

func New(reserved ...string) Filter {
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

	character := value[0]
	if !('a' <= character && character <= 'z') {
		if 'A' <= character && character <= 'Z' {
			character += 'a' - 'A'
		}
	}

	remainingWord := value[1:]
	for _, child := range this.children {
		if child.lowercase == character {
			child.add(remainingWord)
			return
		}
	}

	child := &treeNode{lowercase: character, uppercase: character - ('a' - 'A')}
	this.children = append(this.children, child)
	child.add(remainingWord)
}

func (this *treeNode) Contains(input string) bool {
	contains := true
	inputLength := len(input)

	for index := 0; index < inputLength; index++ {
		if contains, index = this.contains(input, index); contains {
			return true
		}
	}

	return false
}
func (this *treeNode) contains(input string, index int) (bool, int) {
	if len(input) == index || input[index] == ' ' || input[index] == '\n' || input[index] == '\t' {
		if this.isWord == true {
			return true, index
		}

		return false, index
	}

	for _, child := range this.children {
		if input[index] == child.lowercase {
			return child.contains(input, index+1)

		} else if input[index] == child.uppercase {
			return child.contains(input, index+1)

		} else {
			for input[index] != ' ' && input[index] != '\n' && input[index] != '\t' && len(input) != index+1 {
				index += 1
			}

			return false, index
		}
	}

	return false, index
}
