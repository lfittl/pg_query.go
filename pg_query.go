package pg_query

import (
	"encoding/json"
	"runtime/debug"

	"github.com/lfittl/pg_query_go/parser"
)

// ParseToJSON - Parses the given SQL statement into an AST (JSON format)
func ParseToJSON(input string) (result string, err error) {
	return parser.ParseToJSON(input)
}

// Parse the given SQL statement into an AST (native Go structs)
func Parse(input string) (tree ParsetreeList, err error) {
	jsonTree, err := ParseToJSON(input)
	if err != nil {
		return
	}

	// JSON unmarshalling can panic in edge cases we don't support yet. This is
	// still a *bug that needs to be fixed*, but this way the caller can expect an
	// error to be returned always, instead of a panic
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err = r.(error)
		}
	}()

	err = json.Unmarshal([]byte(jsonTree), &tree)
	return
}

// Normalize the passed SQL statement to replace constant values with ? characters
func Normalize(input string) (result string, err error) {
	return parser.Normalize(input)
}
