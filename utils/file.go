package utils

import "os"

func ExistFile(fname string) bool {
	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
