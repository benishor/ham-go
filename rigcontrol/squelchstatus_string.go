// Code generated by "stringer -type=SquelchStatus"; DO NOT EDIT.

package rigcontrol

import "fmt"

const _SquelchStatus_name = "SquelchUnknownSquelchOffSquelchOn"

var _SquelchStatus_index = [...]uint8{0, 14, 24, 33}

func (i SquelchStatus) String() string {
	if i >= SquelchStatus(len(_SquelchStatus_index)-1) {
		return fmt.Sprintf("SquelchStatus(%d)", i)
	}
	return _SquelchStatus_name[_SquelchStatus_index[i]:_SquelchStatus_index[i+1]]
}
