package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

var fileSectionMap = map[string]string{
	"storage.go": "## Storage",
	"memory.go":  "## Memory",
	"stack.go":   "## Stack",
	"evm.go":     "## EVM",
}

func main() {
	readme, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println("Error reading README.md:", err)
		return
	}

	updatedReadme := string(readme)

	for file, section := range fileSectionMap {
		types, err := extractTypes(file)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", file, err)
			continue
		}

		updatedReadme = updateSection(updatedReadme, section, strings.Join(types, "\n\n"))
	}

	err = os.WriteFile("README.md", []byte(updatedReadme), 0644)
	if err != nil {
		fmt.Println("Error writing README.md:", err)
		return
	}

	fmt.Println("README.md has been updated successfully.")
}

func extractTypes(filename string) ([]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var types []string
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					types = append(types, formatTypeWithComments(fset, genDecl, typeSpec))
				}
			}
		}
	}

	return types, nil
}

func formatTypeWithComments(fset *token.FileSet, genDecl *ast.GenDecl, typeSpec *ast.TypeSpec) string {
	var buf strings.Builder

	// Print comments
	if genDecl.Doc != nil {
		for _, comment := range genDecl.Doc.List {
			fmt.Fprintln(&buf, comment.Text)
		}
	}

	// Print type declaration
	fmt.Fprintf(&buf, "type %s ", typeSpec.Name.Name)
	err := printer.Fprint(&buf, fset, typeSpec.Type)
	if err != nil {
		return fmt.Sprintf("Error formatting type %s: %v", typeSpec.Name.Name, err)
	}

	return buf.String()
}

func updateSection(readme, section, newContent string) string {
	parts := strings.SplitN(readme, section, 2)
	if len(parts) != 2 {
		return readme // Section not found
	}

	nextSection := findNextSection(parts[1])
	if nextSection == "" {
		// If no next section, update until the end of the file
		return fmt.Sprintf("%s%s\n\n```go\n%s\n```\n", parts[0], section, newContent)
	}

	// Update content between this section and the next
	return fmt.Sprintf("%s%s\n\n```go\n%s\n```\n\n%s%s", parts[0], section, newContent, nextSection, strings.SplitN(parts[1], nextSection, 2)[1])
}

func findNextSection(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			return line
		}
	}
	return ""
}
