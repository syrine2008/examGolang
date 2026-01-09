package services

import (
	"bufio"
	"fmt"
	"gestAdressNetwork/internal/model"
	"os"
	// "strconv"
	"strings"
)

func Parse(path string) ([]*model.Operateur, error) {



	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var operateurs []*model.Operateur

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var line string = scanner.Text()
		var parts []string = strings.Split(line, ",")

		// x, _ := strconv.ParseFloat(parts[1], 64)
		// y, _ := strconv.ParseFloat(parts[2], 64)

		fmt.Printf("Latitude csv :", parts[1])
		fmt.Printf("Latitude csv :", parts[2])
		//    deuxG , _ := strconv.ParseInt(parts[1])
		//    troisG , _ := strconv.ParseFloat(parts[1], 64)
		//    quatreG , _ := strconv.ParseFloat(parts[1], 64)

		//    o, err := model.New(parts[0],x,y,deuxG,parts[4],parts[5])

		//     operateurs = append(operateurs, o)
	}

	return operateurs, nil

}
