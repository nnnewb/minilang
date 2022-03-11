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
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 35: // ['#','#']
			return 4
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 40: // ['(','(']
			return 5
		case r == 41: // [')',')']
			return 6
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 7
		case r == 45: // ['-','-']
			return 7
		case r == 46: // ['.','.']
			return 8
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		case r == 58: // [':',':']
			return 2
		case r == 59: // [';',';']
			return 10
		case r == 60: // ['<','<']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case 97 <= r && r <= 99: // ['a','c']
			return 11
		case r == 100: // ['d','d']
			return 12
		case 101 <= r && r <= 107: // ['e','k']
			return 11
		case r == 108: // ['l','l']
			return 13
		case 109 <= r && r <= 122: // ['m','z']
			return 11
		case r == 126: // ['~','~']
			return 2
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
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 19
		case r == 92: // ['\','\']
			return 20
		default:
			return 21
		}
	},
	// S4
	func(r rune) int {
		switch {
		case r == 102: // ['f','f']
			return 22
		case r == 116: // ['t','t']
			return 23
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
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 24
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 26
		case 48 <= r && r <= 57: // ['0','9']
			return 27
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 28
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 29
		default:
			return 10
		}
	},
	// S11
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 100: // ['a','d']
			return 18
		case r == 101: // ['e','e']
			return 30
		case 102 <= r && r <= 122: // ['f','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case r == 97: // ['a','a']
			return 31
		case 98 <= r && r <= 122: // ['b','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 21
		case r == 92: // ['\','\']
			return 21
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 19
		case r == 92: // ['\','\']
			return 20
		default:
			return 21
		}
	},
	// S22
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 27
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 28
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 32
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 27
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 33
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 101: // ['a','e']
			return 18
		case r == 102: // ['f','f']
			return 34
		case 103 <= r && r <= 122: // ['g','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 108: // ['a','l']
			return 18
		case r == 109: // ['m','m']
			return 35
		case 110 <= r && r <= 122: // ['n','z']
			return 18
		case r == 126: // ['~','~']
			return 15
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
	// S34
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 104: // ['a','h']
			return 18
		case r == 105: // ['i','i']
			return 36
		case 106 <= r && r <= 122: // ['j','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case r == 97: // ['a','a']
			return 18
		case r == 98: // ['b','b']
			return 37
		case 99 <= r && r <= 122: // ['c','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 109: // ['a','m']
			return 18
		case r == 110: // ['n','n']
			return 38
		case 111 <= r && r <= 122: // ['o','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 99: // ['a','c']
			return 18
		case r == 100: // ['d','d']
			return 39
		case 101 <= r && r <= 122: // ['e','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 100: // ['a','d']
			return 18
		case r == 101: // ['e','e']
			return 40
		case 102 <= r && r <= 122: // ['f','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case r == 97: // ['a','a']
			return 41
		case 98 <= r && r <= 122: // ['b','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 36: // ['$','$']
			return 15
		case r == 37: // ['%','%']
			return 15
		case r == 38: // ['&','&']
			return 15
		case r == 42: // ['*','*']
			return 15
		case r == 43: // ['+','+']
			return 16
		case r == 45: // ['-','-']
			return 16
		case r == 46: // ['.','.']
			return 16
		case r == 47: // ['/','/']
			return 15
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 58: // [':',':']
			return 15
		case r == 60: // ['<','<']
			return 15
		case r == 61: // ['=','=']
			return 15
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 18
		case r == 94: // ['^','^']
			return 15
		case r == 95: // ['_','_']
			return 15
		case 97 <= r && r <= 122: // ['a','z']
			return 18
		case r == 126: // ['~','~']
			return 15
		}
		return NoState
	},
}
