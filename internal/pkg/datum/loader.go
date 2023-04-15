package datum

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) (Index, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("faeild to read file <%s>: %w", filename, err)
	}

	index := make(Index, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		index[strings.TrimSpace(scanner.Text())] = true
	}

	return index, nil
}

func ReadFiles(filenames ...string) (Index, error) {
	index := make(Index, 0)
	for _, name := range filenames {
		i, err := ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("faield to read: %w", err)
		}

		for k, _ := range i {
			index[k] = true
		}
	}

	return index, nil
}
