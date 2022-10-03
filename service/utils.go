package service

import (
	"fmt"
	"strconv"
)

func getIDFromString(idString string) (id int, err error) {
	id, err = strconv.Atoi(idString)
	if err != nil {
		fmt.Println("unable to convert idString into into, err: ", err)
		return id, fmt.Errorf("invalid id")
	}

	if id < 1 {
		return id, fmt.Errorf("invalid id")
	}

	return id, nil
}
