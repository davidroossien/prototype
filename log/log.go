package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {

	fmt.Println("log program Started.")

	dt := time.Now().UTC()
	fmt.Println("Current date and time is: ", dt.String())

	uuidWithHyphen := uuid.New()
	fmt.Println("Generated uuid: " + uuidWithHyphen.String())

	var file, err = os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, os.ErrNotExist) {
			file, err = os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				log.Fatal("Could not open log.txt")
			}
		}
	}

	newLine := dt.String() + ":" + uuidWithHyphen.String() + ": line appended."

	_, err = fmt.Fprintln(file, newLine)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	fmt.Println("log program Exiting.")
}
