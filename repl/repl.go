package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hi6okuni/writing-interpreter-in-go/lexer"
	"github.com/hi6okuni/writing-interpreter-in-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		fmt.Printf(PROMPT)
		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
