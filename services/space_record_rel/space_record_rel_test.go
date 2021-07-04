package space_record_rel

import "testing"

func TestAppendRel(t *testing.T) {
	err := AppendRel("pyqCJMRebP", "abc")
	if err != nil {
		t.Errorf(err.Error())
	}
	err = AppendRel("pyqCJMRebP", "123456")
	if err != nil {
		t.Errorf(err.Error())
	}
}
