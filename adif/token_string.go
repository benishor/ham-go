// Code generated by "stringer -type=Token"; DO NOT EDIT.

package adif

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tokenError-0]
	_ = x[tokenLexError-1]
	_ = x[tokenEOF-2]
	_ = x[tokenEOH-3]
	_ = x[tokenEOR-4]
	_ = x[tokenColon-5]
	_ = x[tokenNL-6]
	_ = x[tokenLAngle-7]
	_ = x[tokenRAngle-8]
	_ = x[tokenOther-9]
}

const _Token_name = "tokenErrortokenLexErrortokenEOFtokenEOHtokenEORtokenColontokenNLtokenLAngletokenRAngletokenOther"

var _Token_index = [...]uint8{0, 10, 23, 31, 39, 47, 57, 64, 75, 86, 96}

func (i Token) String() string {
	if i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
