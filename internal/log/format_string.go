// Code generated by "stringer -type Format,Level"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FormatInvalid-0]
	_ = x[FormatText-1]
	_ = x[FormatJSON-2]
}

const _Format_name = "FormatInvalidFormatTextFormatJSON"

var _Format_index = [...]uint8{0, 13, 23, 33}

func (i Format) String() string {
	if i < 0 || i >= Format(len(_Format_index)-1) {
		return "Format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LevelInvalid-0]
}

const _Level_name = "LevelInvalid"

var _Level_index = [...]uint8{0, 12}

func (i Level) String() string {
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}
