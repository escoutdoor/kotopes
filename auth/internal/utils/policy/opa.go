package policy

import (
	"fmt"
	"os"
	"path"

	"github.com/open-policy-agent/opa/ast"
)

type Policy struct {
	Compiler *ast.Compiler
}

func New(path string) (*Policy, error) {
	compiler, err := load(path)
	if err != nil {
		return nil, err
	}

	return &Policy{
		Compiler: compiler,
	}, nil
}

func load(file string) (*ast.Compiler, error) {
	modules := map[string]string{}

	content, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error load: %w", err)
	}

	modules[path.Base(file)] = string(content)

	return ast.CompileModules(modules)
}
