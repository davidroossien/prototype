package semanticsort

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Represents a semantic version we can use to sort the input array.
type SemanticVersion struct {
	Version string
	major   int
	minor   int
	patch   int
}

var errEmptyData = errors.New("empty input data detected. Please provide an array of inputSemanticVersions")

/*
	Initialize accepts an array of strings that are formatted with semantic versions.
	It iterates through the list.
	It parses the strings for major, minor and patch, and initializes an array of
	semanticVersion structures.

	@param an array of strings that are formatted as semantic versions i.e. 1.0.0
	@return array of semanticVersion structures
*/
func Initialize(inputSemanticVersions []string) ([]SemanticVersion, error) {
	if len(inputSemanticVersions) == 0 {
		return nil, errEmptyData
	}

	// Make a slice we can build and then output.
	outputSemanticVersions := make([]SemanticVersion, len(inputSemanticVersions))

	// Initialize the structs.
	for i := 0; i < len(inputSemanticVersions); i++ {
		// split each string (i.e. 1.0.0) into an array {i.e. 1, 0, 0}
		versions := strings.Split(inputSemanticVersions[i], ".")

		if len(versions) != 3 {
			return nil, errors.New("Invalid inputSemanticVersions in row " + fmt.Sprint(i) +
				", unable to split into major, minor, patch versions.")
		}

		var semver SemanticVersion

		var err error

		// Put the whole string in version (i.e. "1.0.0").
		semver.Version = inputSemanticVersions[i]
		semver.major, err = strconv.Atoi(versions[0])

		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " +
				fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// Convert and assign minor.
		semver.minor, err = strconv.Atoi(versions[1])
		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " +
				fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// Convert and assign patch.
		semver.patch, err = strconv.Atoi(versions[2])
		if err != nil {
			return nil, errors.New("Invalid patch version in input inputSemanticVersions row " +
				fmt.Sprint(i) + ", perhaps non-numeric.")
		}

		// store in our output slice
		outputSemanticVersions[i] = semver
	}

	return outputSemanticVersions, nil
}

/*
	SemanticSort sorts an array of semantic versions that have been pre-loaded into
	an array of semanticVersion structures.

	Uses sort.Slice with a custom algorith to sort a semanticVersion structs.

	@param inputSemanticVersions an array of semanticVersion structures to be sorted
	@return a sorted array of semanticVersion structures
*/
func SemanticSort(inputSemanticVersions []SemanticVersion) []SemanticVersion {
	sortedSemanticVersions := make([]SemanticVersion, len(inputSemanticVersions))
	sortedSemanticVersions = inputSemanticVersions

	// Do the sortation.
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

		// It should never get here unless there is invalid input
		// data that contains duplicates.
		return sortedSemanticVersions[j].major < sortedSemanticVersions[k].major
	})

	return sortedSemanticVersions
}
