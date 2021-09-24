package prefilter

type prefilterTree struct {
	Prefilter
}


func newPrefilterTree(reserved ...string) Prefilter {

	return &prefilterTree{}
}

func (this *prefilter) IsAllowed3(input string) bool {
	return true
}