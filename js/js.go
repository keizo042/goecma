package js

import (
	"fmt"
	"os"

	"github.com/keizo042/goecma/js/ast"
	"github.com/keizo042/goecma/js/token"
)

func New() (*State, error) {
	return nil, fmt.Errorf("TBD")
}

func (s *State) ReadFile(f *os.File) ([]token.Value, error) {
	return nil, fmt.Errorf("TBD")
}

func (s *State) Parse(tokens []token.Value) (ast.Tree, error) {
	return nil, fmt.Errorf("TBD")
}

func (s *State) Run(tree ast.Tree, argc int, argv **Value) (Value, error) {
	var v Value
	return v, fmt.Errorf("TBD")
}
