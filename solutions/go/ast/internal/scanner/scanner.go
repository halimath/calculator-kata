// Package scabber provides a type for scanning an io.Reader for token.Token values. It reports syntax errors
// while scanning.
package scanner

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/halimath/calc/internal/token"
)

// ErrScanFailed is returned when the lexer hits invalid input.
var ErrScanFailed = errors.New("scan failed")

// Scanner implements scanning an io.Reader for tokens.
type Scanner struct {
	r     bufio.Reader
	value strings.Builder
}

// New creates a new Scanner consuming input from r.
func New(r io.Reader) *Scanner {
	l := Scanner{
		r: *bufio.NewReader(r),
	}
	return &l
}

// Next consumes the next token from l and returns it. If no more tokens are available, the returned token
// is nil and io.EOF is returned as the error. In any other non-nil value represents an scanning error.
func (s *Scanner) Next() (token.Token, error) {
	for {
		r, _, err := s.r.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return s.consumeNumber(io.EOF)
			}

			return nil, fmt.Errorf("%w: %v", ErrScanFailed, err)
		}

		if unicode.IsSpace(r) {
			if s.value.Len() == 0 {
				// If nothing has been consumed so far, simply skip whitespace
				continue
			}

			// Otherwise sb must contain digits, so it must be a number
			return s.consumeNumber(ErrScanFailed)
		}

		// If r is a digit or a dot, append it to the buffer and continue consuming runes
		if unicode.IsDigit(r) || r == '.' {
			s.value.WriteRune(r)
			continue
		}

		if s.value.Len() > 0 {
			// If so, unread r and return a number
			if err = s.r.UnreadRune(); err != nil {
				return nil, fmt.Errorf("%w: %v", ErrScanFailed, err)
			}

			return s.consumeNumber(ErrScanFailed)
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
			return nil, fmt.Errorf("%w: invalid input rune: %c", ErrScanFailed, r)
		}
	}
}

func (s *Scanner) consumeNumber(errToReturn error) (token.Token, error) {
	if s.value.Len() == 0 {
		return nil, errToReturn
	}

	tok := token.Number(s.value.String())

	s.value.Reset()

	return tok, nil
}
