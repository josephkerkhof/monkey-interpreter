// Package repl manages the read-eval-print-loop.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
