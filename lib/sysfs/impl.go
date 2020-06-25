package sysfs

import (
	"fmt"
	"os"
)

func readUint64(filename string) (uint64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var value uint64
	nScanned, err := fmt.Fscanf(file, "%d", &value)
	if err != nil {
		return 0, err
	}
	if nScanned < 1 {
		return 0, fmt.Errorf("only read %d values from: %s", nScanned, filename)
	}
	return value, nil
}

func readBool(filename string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()
	var ivalue uint
	nScanned, err := fmt.Fscanf(file, "%d", &ivalue)
	if err != nil {
		return false, err
	}
	if nScanned < 1 {
		return false, fmt.Errorf("only read %d values from: %s",
			nScanned, filename)
	}
	if ivalue == 0 {
		return false, nil
	}
	return true, nil
}

func readString(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var value string
	nScanned, err := fmt.Fscanf(file, "%s", &value)
	if err != nil {
		return "", err
	}
	if nScanned < 1 {
		return "", fmt.Errorf("only read %d values from: %s",
			nScanned, filename)
	}
	return value, nil
}
