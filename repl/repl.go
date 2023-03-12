// REPLは、Read-Evel-Print Loopの略。
// 直訳するなら、入力を受け取って、評価して、結果を出力するループ。
// よくある人間がCLIでプログラムを1行ずつ入力しながら実行するアレ。
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tkn := l.NextToken(); tkn.Type != token.EOF; tkn = l.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}
}