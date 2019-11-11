package env

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// ReadFiles reads key=value pairs from the files in the glob expression and stores the values in a dictionary.
func ReadFiles(globPath string, dest map[string]string) error {
	if matches, err := filepath.Glob(globPath); err != nil {
		return err
	} else {
		for _, match := range matches {
			if err := ReadFile(match, dest); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReadFile reads key=value pairs from a file and stores the values in a dictionary.
func ReadFile(filePath string, dest map[string]string) error {
	if contents, err := ioutil.ReadFile(filePath); err != nil {
		return err
	} else {
		return ReadString(string(contents), dest)
	}
}

// ReadString reads key=value pairs from a string and stores the values in a dictionary.
func ReadString(contents string, dest map[string]string) error {
	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) == 2 && len(parts[0]) > 0 && parts[0][0] != '#' {
			key := parts[0]
			if _, ok := dest[key]; ok {
				return fmt.Errorf("key '%s' already set", key)
			}

			dest[key] = parts[1]
		}
	}

	return scanner.Err()
}
