package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/maiyama18/immonkey/token"

	"github.com/maiyama18/immonkey/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)

	_, _ = fmt.Fprint(out, PROMPT)
	for sc.Scan() {
		input := sc.Text()
		l := lexer.New(input)

		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tk)
		}

		_, _ = fmt.Fprint(out, PROMPT)
	}
}
