package files

import (
	"fmt"
	"os"
)

func ReadFromFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return data, nil
}

func WriteToFile(data []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Print(err)
	}
	_, err = file.Write(data)
	defer file.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
}
