package mattefixer

import (
	"fmt"
	"strconv"
)

func ParseArgs(args []string) (string, int, string, error) {
	filePath := args[1]
	strength := 100
	format := "BC3_UNORM"

	if len(args) > 2 {
		val, err := strconv.Atoi(args[2])
		if err == nil {
			if val < 0 {
				val = 0
			}
			if val > 100 {
				val = 100
			}
			strength = val
		} else {
			format = args[2]
		}
	}

	if len(args) > 3 {
		if strength != 100 {
			format = args[3]
		} else {
			val, err := strconv.Atoi(args[3])
			if err == nil {
				if val < 0 {
					val = 0
				}
				if val > 100 {
					val = 100
				}
				strength = val
			}
		}
	}

	if filePath == "" {
		return "", 0, "", fmt.Errorf("invalid file path")
	}

	return filePath, strength, format, nil
}
