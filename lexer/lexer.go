package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/halimath/calc/token"
)

var ErrLexer = errors.New("lexer error")

type Lexer struct {
	r     bufio.Reader
	value strings.Builder
}

func New(r io.Reader) *Lexer {
	return &Lexer{
		r: *bufio.NewReader(r),
	}
}

func (l *Lexer) Next() (tok token.Token, err error) {
	l.value.Reset()
	var r rune

	for {
		r, _, err = l.r.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return l.consumeNumberOrErr(io.EOF)
			}

			err = fmt.Errorf("%w: %v", ErrLexer, err)
			return
		}

		if unicode.IsSpace(r) {
			if l.value.Len() == 0 {
				// If nothing has been consumed so far, simply skip whitespace
				continue
			}

			// Otherwise sb must contain digits, so it must be a number
			return l.consumeNumberOrErr(ErrLexer)
		}

		// If r is a digit or a dot, append it to the buffer and continue consuming runes
		if unicode.IsDigit(r) || r == '.' {
			l.value.WriteRune(r)
			continue
		}

		// Otherwise, r is either an operator or a parenthesis. Check if we have digits consumed so far.
		if l.value.Len() > 0 {
			// If so, unread r and return a number
			if err = l.r.UnreadRune(); err != nil {
				err = fmt.Errorf("%w: %v", ErrLexer, err)
				return
			}

			return l.consumeNumberOrErr(ErrLexer)
		}

		l.value.WriteRune(r)
		tok.Value = l.value.String()

		switch r {
		case '+':
			tok.Type = token.Add
		case '-':
			tok.Type = token.Sub
		case '*':
			tok.Type = token.Mul
		case '/':
			tok.Type = token.Div
		case '(':
			tok.Type = token.LParen
		case ')':
			tok.Type = token.RParen
		}

		return
	}
}

func (l *Lexer) consumeNumberOrErr(errToReturn error) (tok token.Token, err error) {
	var ok bool

	tok, ok = l.consumeNumber()
	if !ok {
		err = errToReturn
	}

	return
}

func (l *Lexer) consumeNumber() (tok token.Token, ok bool) {
	if l.value.Len() == 0 {
		return
	}

	tok.Type = token.Number
	tok.Value = l.value.String()
	ok = true

	l.value.Reset()

	return
}
