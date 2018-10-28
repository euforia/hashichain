package compose

import (
	"errors"
	"fmt"
	"strings"
)

func splitImageNameTag(name string) (string, string, error) {
	parts := strings.Split(name, ":")
	l := len(parts)
	switch l {
	case 0:
		return "", "", errors.New("image name required")
	case 1:
		return parts[0], "", nil
	case 2:
		return parts[0], parts[1], nil
	}

	return "", "", fmt.Errorf("invalid image name or tag: '%s'", name)
}
