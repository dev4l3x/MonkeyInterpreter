package repl

import (
	"bufio"
	"fmt"
	"monkeyinterpreter/lexer"
	"monkeyinterpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		l := lexer.New(scanner.Text())

		for tok := l.NextToken(); tok.Type != token.EOF ; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}

}
