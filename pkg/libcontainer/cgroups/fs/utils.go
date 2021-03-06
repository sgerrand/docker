package fs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	ErrNotSupportStat = errors.New("stats are not supported for subsystem")
	ErrNotValidFormat = errors.New("line is not a valid key value format")
)

// Parses a cgroup param and returns as name, value
//  i.e. "io_service_bytes 1234" will return as io_service_bytes, 1234
func getCgroupParamKeyValue(t string) (string, int64, error) {
	parts := strings.Fields(t)
	switch len(parts) {
	case 2:
		value, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return "", 0, fmt.Errorf("Unable to convert param value to int: %s", err)
		}
		return parts[0], value, nil
	default:
		return "", 0, ErrNotValidFormat
	}
}

// Gets a single int64 value from the specified cgroup file.
func getCgroupParamInt(cgroupPath, cgroupFile string) (int64, error) {
	contents, err := ioutil.ReadFile(filepath.Join(cgroupPath, cgroupFile))
	if err != nil {
		return -1, err
	}
	return strconv.ParseInt(strings.TrimSpace(string(contents)), 10, 64)
}
