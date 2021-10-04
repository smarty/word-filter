package prefilter

type prefilterTree struct {
	buffer         []byte
	restrictedRoot *treeNode
}

func newPrefilterTree(bufferSize int, reserved ...string) *treeNode {
	b := make([]byte, bufferSize, bufferSize)
	root := treeNode{}
	for _, word := range reserved {
		if err := root.Add([]byte(word)); err != nil { //TODO error handling
			return nil
		}
	}
	tree := prefilterTree{
		buffer:         b,
		restrictedRoot: &root,
	}
	return tree.restrictedRoot
}
