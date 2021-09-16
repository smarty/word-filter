package prefilter

import (
	"reflect"
	"testing"
)

func TestPrefiltering(t *testing.T) {
	stringToCheck := "Another thing to check"
	myPrefilter := newPrefilter(len(stringToCheck), "ANOTHER", "RANDOM", "WORD")
	Assert(t).That(myPrefilter.IsAllowed(stringToCheck)).Equals(false)

	stringToCheck = "123 what if there are numbers"
	Assert(t).That(myPrefilter.IsAllowed(stringToCheck)).Equals(true)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type That struct{ t *testing.T }
type Assertion struct {
	*testing.T
	actual interface{}
}

func Assert(t *testing.T) *That                       { return &That{t: t} }
func (this *That) That(actual interface{}) *Assertion { return &Assertion{T: this.t, actual: actual} }

func (this *Assertion) IsNil() {
	this.Helper()
	if this.actual != nil && !reflect.ValueOf(this.actual).IsNil() {
		this.Equals(nil)
	}
}
func (this *Assertion) Equals(expected interface{}) {
	this.Helper()
	if !reflect.DeepEqual(this.actual, expected) {
		this.Errorf("\nExpected: %#v\nActual:   %#v", expected, this.actual)
	}
}
