// Package lexer provides a type for scanning an io.Reader for token.Token values. It reports syntax errors
// while scanning.
package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/halimath/calc/token"
)

// ErrLexer is returned when the lexer hits invalid input.
var ErrLexer = errors.New("lexer error")

// Lexer implements scanning an io.Reader for tokens.
type Lexer struct {
	r     bufio.Reader
	value strings.Builder
}

// New creates a new Lexer consuming input from r.
func New(r io.Reader) *Lexer {
	l := Lexer{
		r: *bufio.NewReader(r),
	}
	return &l
}

// Next consumes the next token from l and returns it. If no more tokens are available, the returned token
// is nil and io.EOF is returned as the error. In any other non-nil value represents an scanning error.
func (l *Lexer) Next() (token.Token, error) {
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return l.consumeNumber(io.EOF)
			}

			return nil, fmt.Errorf("%w: %v", ErrLexer, err)
		}

		if unicode.IsSpace(r) {
			if l.value.Len() == 0 {
				// If nothing has been consumed so far, simply skip whitespace
				continue
			}

			// Otherwise sb must contain digits, so it must be a number
			return l.consumeNumber(ErrLexer)
		}

		// If r is a digit or a dot, append it to the buffer and continue consuming runes
		if unicode.IsDigit(r) || r == '.' {
			l.value.WriteRune(r)
			continue
		}

		if l.value.Len() > 0 {
			// If so, unread r and return a number
			if err = l.r.UnreadRune(); err != nil {
				return nil, fmt.Errorf("%w: %v", ErrLexer, err)
			}

			return l.consumeNumber(ErrLexer)
		}

		switch r {
		case '+':
			return token.Add, nil
		case '-':
			return token.Sub, nil
		case '*':
			return token.Mul, nil
		case '/':
			return token.Div, nil
		case '(':
			return token.LParen, nil
		case ')':
			return token.RParen, nil
		default:
			return nil, fmt.Errorf("%w: invalid input rune: %c", ErrLexer, r)
		}
	}
}

func (l *Lexer) consumeNumber(errToReturn error) (token.Token, error) {
	if l.value.Len() == 0 {
		return nil, errToReturn
	}

	val, err := strconv.ParseFloat(l.value.String(), 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrLexer, err)
	}

	tok := token.Number(val)

	l.value.Reset()

	return tok, nil
}
