// Code generated by "stringer -type SettingKey"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DestinationSetting-0]
	_ = x[FormatSetting-1]
	_ = x[OmitTimeSetting-2]
}

const _SettingKey_name = "DestinationSettingFormatSettingOmitTimeSetting"

var _SettingKey_index = [...]uint8{0, 18, 31, 46}

func (i SettingKey) String() string {
	if i < 0 || i >= SettingKey(len(_SettingKey_index)-1) {
		return "SettingKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SettingKey_name[_SettingKey_index[i]:_SettingKey_index[i+1]]
}
