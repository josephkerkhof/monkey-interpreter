package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements should be 3 statements. got=%d",
			len(program.Statements),
		)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestErrors(t *testing.T) {
	input := `
let x 5;
let y = 10;
let 838383;
`
	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()

	errors := p.Errors()

	if len(errors) != 2 {
		t.Fatalf(
			"p.Errors() should have 2 errors. got=%d",
			len(errors),
		)
	}

	tests := []struct {
		expectedError string
	}{
		{"expected next token to be =, got INT instead"},
		{"expected next token to be IDENT, got INT instead"},
	}

	for i, tt := range tests {
		e := errors[i]
		if e != tt.expectedError {
			t.Errorf("error message not correct. got=%q", e)
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	t.Helper()

	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, err := range errors {
		t.Errorf("parser error: %s", err)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	t.Helper()

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %q. got=%q", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%q", name, letStmt.Name)
		return false
	}

	return true
}
