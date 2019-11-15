package images

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string, output *string) error {
	if file, err := os.Open(filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		var tmp string
		if nScanned, err := fmt.Fscanf(reader, "%s", &tmp); err != nil {
			return err
		} else if nScanned != 1 {
			return fmt.Errorf("got: %d fields, expected 1", nScanned)
		} else {
			*output = tmp
			return nil
		}
	}
}

func (p *prober) probe() error {
	if err := readFile("/var/lib/initial-image", &p.initialImage); err != nil {
		return err
	}
	if err := readFile("/var/lib/patched-image", &p.patchedImage); err != nil {
		return err
	}
	return nil
}
