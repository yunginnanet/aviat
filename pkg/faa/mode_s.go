package faa

import (
	"strconv"
	"strings"
)

func CompactModeS(modeS string) string {
	modeS = strings.TrimSpace(modeS)
	modeS = strings.TrimLeft(modeS, "0")
	if modeS == "" {
		return "0"
	}

	if num, err := strconv.ParseUint(modeS, 10, 64); err == nil {
		return strings.ToUpper(strconv.FormatUint(num, 16))
	}

	return modeS
}
