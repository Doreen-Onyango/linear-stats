package linears

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadData(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing value: %s", line)
		}
		data = append(data, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
