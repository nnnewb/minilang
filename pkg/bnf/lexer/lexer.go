// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/nnnewb/scsh/pkg/bnf/token"
)

const (
	NoState    = -1
	NumStates  = 37
	NumSymbols = 173
)

type Lexer struct {
	src     []byte
	pos     int
	line    int
	column  int
	Context token.Context
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:     src,
		pos:     0,
		line:    1,
		column:  1,
		Context: nil,
	}
	return lexer
}

// SourceContext is a simple instance of a token.Context which
// contains the name of the source file.
type SourceContext struct {
	Filepath string
}

func (s *SourceContext) Source() string {
	return s.Filepath
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	lexer := NewLexer(src)
	lexer.Context = &SourceContext{Filepath: fpath}
	return lexer, nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = &token.Token{}
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		tok.Pos.Context = l.Context
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn
	tok.Pos.Context = l.Context

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '('
1: ')'
2: '#'
3: '('
4: '''
5: '`'
6: ','
7: ','
8: '@'
9: '.'
10: 'q'
11: 'u'
12: 'o'
13: 't'
14: 'e'
15: 'l'
16: 'a'
17: 'm'
18: 'b'
19: 'd'
20: 'a'
21: 'i'
22: 'f'
23: 's'
24: 'e'
25: 't'
26: '!'
27: 'b'
28: 'e'
29: 'g'
30: 'i'
31: 'n'
32: 'c'
33: 'o'
34: 'n'
35: 'd'
36: 'a'
37: 'n'
38: 'd'
39: 'o'
40: 'r'
41: 'c'
42: 'a'
43: 's'
44: 'e'
45: 'l'
46: 'e'
47: 't'
48: 'l'
49: 'e'
50: 't'
51: '*'
52: 'l'
53: 'e'
54: 't'
55: 'r'
56: 'e'
57: 'c'
58: 'd'
59: 'o'
60: 'd'
61: 'e'
62: 'l'
63: 'a'
64: 'y'
65: 'q'
66: 'u'
67: 'a'
68: 's'
69: 'i'
70: 'q'
71: 'u'
72: 'o'
73: 't'
74: 'e'
75: 'e'
76: 'l'
77: 's'
78: 'e'
79: '='
80: '>'
81: 'd'
82: 'e'
83: 'f'
84: 'i'
85: 'n'
86: 'e'
87: 'u'
88: 'n'
89: 'q'
90: 'u'
91: 'o'
92: 't'
93: 'e'
94: 'u'
95: 'n'
96: 'q'
97: 'u'
98: 'o'
99: 't'
100: 'e'
101: '-'
102: 's'
103: 'p'
104: 'l'
105: 'i'
106: 'c'
107: 'i'
108: 'n'
109: 'g'
110: '!'
111: '$'
112: '%'
113: '&'
114: '*'
115: '/'
116: ':'
117: '<'
118: '='
119: '>'
120: '?'
121: '~'
122: '_'
123: '^'
124: '.'
125: '+'
126: '-'
127: '+'
128: '-'
129: '.'
130: '.'
131: '.'
132: '#'
133: 't'
134: '#'
135: 'f'
136: '#'
137: '\'
138: '#'
139: '\'
140: 's'
141: 'p'
142: 'a'
143: 'c'
144: 'e'
145: '#'
146: '\'
147: 'n'
148: 'e'
149: 'w'
150: 'l'
151: 'i'
152: 'n'
153: 'e'
154: '\'
155: '"'
156: '\'
157: '\'
158: '"'
159: '"'
160: '+'
161: '-'
162: '.'
163: '.'
164: ';'
165: ' '
166: '\t'
167: '\r'
168: '\n'
169: 'a'-'z'
170: 'A'-'Z'
171: '0'-'9'
172: .
*/
