// Package dbreport provides sysbench dbreport parser
package ioreport

import (
	"fmt"
	"io"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

type badStringError struct {
	what string
	str  string
}

func (e *badStringError) Error() string { return fmt.Sprintf("%s %q", e.what, e.str) }

// ParseDBReport is parse sysbench sql benchmark report
func ParseDBReport(r io.Reader) (*SysbenchInfo, error) {
	var sysbenchIOLexer = lexer.Must(lexer.Regexp(
		`(?m)` +
			`(\s+)` +
			`|(^Running the test with following options:$)` +
			`|(^Initializing random number generator from current time$)` +
			`|(^Initializing worker threads...$)` +
			`|(^Threads started!)` +
			`|(?P<VersionNumber>[\d]+\.[\d]+\.[\w\-]+)` +
			`|(?P<Float>\d+(?:\.\d+)+)` +
			`|(?P<Int>[+-]?[\d]+)` +
			`|(?P<Ident>[a-zA-Z0-9][a-zA-Z0-9%_:\./-]*)` +
			`|(?P<Punct>[()\[\]\/:])`,
	))

	parser, err := participle.Build(
		&SysbenchInfo{},
		participle.Lexer(sysbenchIOLexer),
	)

	if err != nil {
		return nil, &badStringError{"Failed to load struct", err.Error()}
	}

	sb := &SysbenchInfo{}

	err = parser.Parse(r, sb)
	if err != nil {
		return nil, &badStringError{"Failed to parse", err.Error()}
	}

	return sb, nil
}
