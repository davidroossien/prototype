package simplelog

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

/*
   Configures the simple logger.
*/
func Config(filename string) (int, error) {
	fmt.Println("log program Started.")

	dt := time.Now().UTC()
	fmt.Println("Current date and time is: ", dt.String())

	uuidWithHyphen := uuid.New()
	fmt.Println("Generated uuid: " + uuidWithHyphen.String())

	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, os.ErrNotExist) {
			logFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				// log.Fatal("Could not open " + filename)
				return -1, errors.New("could not open filename " + filename)
			}
		}
	}

	newLine := dt.String() + ":" + uuidWithHyphen.String() + ": log file created."

	_, err = fmt.Fprintln(logFile, newLine)
	if err != nil {
		fmt.Println(err)
		return -2, errors.New("could not write to filename " + filename)
	}

	defer logFile.Close()

	fmt.Println("log program Exiting.")

	return 0, nil
}

// func log(statement string)
// {

// }
