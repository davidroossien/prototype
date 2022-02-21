package semanticSort

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var inputSemanticVersions = [6]string{"1.5.2", "1.0.1", "1.2.0", "1.1.1", "1.3.0", "1.2.11"}

// represents a semantic version we can use to sort the input array
type SemanticVersion struct {
	Version string
	major   int
	minor   int
	patch   int
}

/*
   Uses sort.Slice to sort a semanticVersion structs
*/
// func main() {

// 	fmt.Println("Start...")

// 	initalizedSemanticVersions, err := initialize(inputSemanticVersions[:])
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	sortedSemanticVersions := semanticSort(initalizedSemanticVersions)

// 	fmt.Println("Sorted semantic versions:")

// 	// print the result
// 	for i := 0; i < len(sortedSemanticVersions); i++ {
// 		fmt.Println(sortedSemanticVersions[i].version)
// 	}

// 	fmt.Println("Done...")
// }

/*
initialize accepts an array of strings that are formatted with semantic versions.
It iterates through the list
It parses the strings for major, minor and patch, and initializes an array of
semanticVersion structures.

@param an array of strings that are formatted as semantic versions i.e. 1.0.0
@return array of semanticVersion structures
*/
func Initialize(inputSemanticVersions []string) ([]SemanticVersion, error) {

	if len(inputSemanticVersions) == 0 {
		return nil, errors.New("Empty input data detected. Please provide an array of inputSemanticVersions.")
	}

	// make a slice we can build and then output
	outputSemanticVersions := make([]SemanticVersion, len(inputSemanticVersions))

	// initialize the structs
	for i := 0; i < len(inputSemanticVersions); i++ {

		// split each string (i.e. 1.0.0) into an array {i.e. 1, 0, 0}
		versions := strings.Split(inputSemanticVersions[i], ".")

		if len(versions) == 0 {
			return nil, errors.New("Invalid inputSemanticVersions in row " + fmt.Sprint(i) + ", unable to split into major, minor, patch versions.")
		} else if len(versions) != 3 {
			return nil, errors.New("Invalid inputSemanticVersions in row " + fmt.Sprint(i) + ", missing at least one semantic (major, minor or patch).")
		}

		var sv SemanticVersion
		var err error

		// put the whole string in version (i.e. "1.0.0")
		sv.Version = inputSemanticVersions[i]

		// convert and assign major
		sv.major, err = strconv.Atoi(versions[0])
		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " + fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// convert and assign minor
		sv.minor, err = strconv.Atoi(versions[1])
		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " + fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// convert and assign patch
		sv.patch, err = strconv.Atoi(versions[2])
		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " + fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// store in our output slice
		outputSemanticVersions[i] = sv
	}

	return outputSemanticVersions, nil
}

/*
semanticSort sorts an array of semantic versions that have been pre-loaded into
an array of semanticVersion structures

@param inputSemanticVersions an array of semanticVersion structures to be sorted
@return a sorted array of semanticVersion structures
*/
func SemanticSort(inputSemanticVersions []SemanticVersion) []SemanticVersion {

	sortedSemanticVersions := make([]SemanticVersion, len(inputSemanticVersions))
	sortedSemanticVersions = inputSemanticVersions

	// do the sortation
	sort.Slice(sortedSemanticVersions, func(j, k int) bool {

		if sortedSemanticVersions[j].major != sortedSemanticVersions[k].major {
			return sortedSemanticVersions[j].major < sortedSemanticVersions[k].major
		}

		if sortedSemanticVersions[j].minor != sortedSemanticVersions[k].minor {
			return sortedSemanticVersions[j].minor < sortedSemanticVersions[k].minor
		}

		if sortedSemanticVersions[j].patch != sortedSemanticVersions[k].patch {
			return sortedSemanticVersions[j].patch < sortedSemanticVersions[k].patch
		}

		return sortedSemanticVersions[j].major < sortedSemanticVersions[k].major
	})

	return sortedSemanticVersions
}
