package prefilter

type prefilterTree struct {
	Prefilter
}

type treeNode struct {
	Letter rune
	Parent *treeNode
	Children []*treeNode
}

func newPrefilterTree(reserved ...string) Prefilter {

	return &prefilterTree{}
}

func (this *prefilter) IsAllowed3(input string) bool {
	return true
}