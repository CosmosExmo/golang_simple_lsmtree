package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFromFile(filename string) (map[int]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[int]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid data format in file: %s", filename)
		}
		key, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error converting key to integer: %v", err)
		}
		data[key] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func WriteToFile(filename string, data map[int]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range data {
		_, err := file.WriteString(fmt.Sprintf("%d:%s\n", key, value))
		if err != nil {
			return err
		}
	}

	return nil
}