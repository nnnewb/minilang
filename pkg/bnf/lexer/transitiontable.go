// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 35: // ['#','#']
			return 3
		case r == 40: // ['(','(']
			return 4
		case r == 41: // [')',')']
			return 5
		case r == 43: // ['+','+']
			return 6
		case r == 45: // ['-','-']
			return 6
		case r == 46: // ['.','.']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 59: // [';',';']
			return 9
		case 65 <= r && r <= 90: // ['A','Z']
			return 10
		case 97 <= r && r <= 122: // ['a','z']
			return 10
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 11
		case r == 92: // ['\','\']
			return 12
		default:
			return 13
		}
	},
	// S3
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 14
		case r == 43: // ['+','+']
			return 15
		case r == 45: // ['-','-']
			return 15
		case r == 46: // ['.','.']
			return 16
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case 97 <= r && r <= 101: // ['a','e']
			return 17
		case r == 102: // ['f','f']
			return 18
		case 103 <= r && r <= 115: // ['g','s']
			return 17
		case r == 116: // ['t','t']
			return 19
		case 117 <= r && r <= 122: // ['u','z']
			return 17
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 20
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 21
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 23
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 24
		default:
			return 9
		}
	},
	// S10
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case r == 43: // ['+','+']
			return 25
		case r == 45: // ['-','-']
			return 25
		case r == 46: // ['.','.']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 63: // ['?','?']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 13
		case r == 92: // ['\','\']
			return 13
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 11
		case r == 92: // ['\','\']
			return 12
		default:
			return 13
		}
	},
	// S14
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 28
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 32
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 33
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case r == 43: // ['+','+']
			return 25
		case r == 45: // ['-','-']
			return 25
		case r == 46: // ['.','.']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 63: // ['?','?']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case r == 43: // ['+','+']
			return 25
		case r == 45: // ['-','-']
			return 25
		case r == 46: // ['.','.']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 63: // ['?','?']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case r == 43: // ['+','+']
			return 25
		case r == 45: // ['-','-']
			return 25
		case r == 46: // ['.','.']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 63: // ['?','?']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 15
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 29
		case r == 43: // ['+','+']
			return 29
		case r == 45: // ['-','-']
			return 29
		case r == 46: // ['.','.']
			return 29
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case r == 63: // ['?','?']
			return 29
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 33
		}
		return NoState
	},
}
