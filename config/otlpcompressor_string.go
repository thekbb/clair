// Code generated by "stringer -type OTLPCompressor -linecomment"; DO NOT EDIT.

package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OTLPCompressUnset-0]
	_ = x[OTLPCompressNone-1]
	_ = x[OTLPCompressGzip-2]
}

const _OTLPCompressor_name = "nonegzip"

var _OTLPCompressor_index = [...]uint8{0, 0, 4, 8}

func (i OTLPCompressor) String() string {
	if i < 0 || i >= OTLPCompressor(len(_OTLPCompressor_index)-1) {
		return "OTLPCompressor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OTLPCompressor_name[_OTLPCompressor_index[i]:_OTLPCompressor_index[i+1]]
}