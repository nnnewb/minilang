// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // ␚
			shift(4), // (
			nil,      // define
			nil,      // )
			nil,      // lambda
			nil,      // boolean_t
			nil,      // boolean_f
			nil,      // identifier
			nil,      // uint
			nil,      // string
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // ␚
			shift(4),     // (
			nil,          // define
			nil,          // )
			nil,          // lambda
			nil,          // boolean_t
			nil,          // boolean_f
			nil,          // identifier
			nil,          // uint
			nil,          // string
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(3), // ␚, reduce: Program
			reduce(3), // (, reduce: Program
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(4), // ␚, reduce: Program
			reduce(4), // (, reduce: Program
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			shift(7),  // define
			nil,       // )
			shift(9),  // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(10), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // ␚, reduce: Program
			reduce(1), // (, reduce: Program
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // ␚, reduce: Program
			reduce(2), // (, reduce: Program
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(12), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(16), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(24), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(15), // (, reduce: Identifier
			nil,        // define
			reduce(15), // ), reduce: Identifier
			nil,        // lambda
			reduce(15), // boolean_t, reduce: Identifier
			reduce(15), // boolean_f, reduce: Identifier
			reduce(15), // identifier, reduce: Identifier
			reduce(15), // uint, reduce: Identifier
			reduce(15), // string, reduce: Identifier
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(26), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(15), // (, reduce: Identifier
			nil,        // define
			nil,        // )
			nil,        // lambda
			nil,        // boolean_t
			nil,        // boolean_f
			nil,        // identifier
			nil,        // uint
			nil,        // string
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(20), // (, reduce: Value
			nil,        // define
			reduce(20), // ), reduce: Value
			nil,        // lambda
			reduce(20), // boolean_t, reduce: Value
			reduce(20), // boolean_f, reduce: Value
			reduce(20), // identifier, reduce: Value
			reduce(20), // uint, reduce: Value
			reduce(20), // string, reduce: Value
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			shift(28), // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(10), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(16), // (, reduce: Value
			nil,        // define
			reduce(16), // ), reduce: Value
			nil,        // lambda
			reduce(16), // boolean_t, reduce: Value
			reduce(16), // boolean_f, reduce: Value
			reduce(16), // identifier, reduce: Value
			reduce(16), // uint, reduce: Value
			reduce(16), // string, reduce: Value
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(8), // ␚, reduce: Combination
			reduce(8), // (, reduce: Combination
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(29), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(11), // (, reduce: Operand
			nil,        // define
			reduce(11), // ), reduce: Operand
			nil,        // lambda
			reduce(11), // boolean_t, reduce: Operand
			reduce(11), // boolean_f, reduce: Operand
			reduce(11), // identifier, reduce: Operand
			reduce(11), // uint, reduce: Operand
			reduce(11), // string, reduce: Operand
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(17), // (, reduce: Value
			nil,        // define
			reduce(17), // ), reduce: Value
			nil,        // lambda
			reduce(17), // boolean_t, reduce: Value
			reduce(17), // boolean_f, reduce: Value
			reduce(17), // identifier, reduce: Value
			reduce(17), // uint, reduce: Value
			reduce(17), // string, reduce: Value
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(13), // (, reduce: BooleanLit
			nil,        // define
			reduce(13), // ), reduce: BooleanLit
			nil,        // lambda
			reduce(13), // boolean_t, reduce: BooleanLit
			reduce(13), // boolean_f, reduce: BooleanLit
			reduce(13), // identifier, reduce: BooleanLit
			reduce(13), // uint, reduce: BooleanLit
			reduce(13), // string, reduce: BooleanLit
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(14), // (, reduce: BooleanLit
			nil,        // define
			reduce(14), // ), reduce: BooleanLit
			nil,        // lambda
			reduce(14), // boolean_t, reduce: BooleanLit
			reduce(14), // boolean_f, reduce: BooleanLit
			reduce(14), // identifier, reduce: BooleanLit
			reduce(14), // uint, reduce: BooleanLit
			reduce(14), // string, reduce: BooleanLit
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(18), // (, reduce: Value
			nil,        // define
			reduce(18), // ), reduce: Value
			nil,        // lambda
			reduce(18), // boolean_t, reduce: Value
			reduce(18), // boolean_f, reduce: Value
			reduce(18), // identifier, reduce: Value
			reduce(18), // uint, reduce: Value
			reduce(18), // string, reduce: Value
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(19), // (, reduce: Value
			nil,        // define
			reduce(19), // ), reduce: Value
			nil,        // lambda
			reduce(19), // boolean_t, reduce: Value
			reduce(19), // boolean_f, reduce: Value
			reduce(19), // identifier, reduce: Value
			reduce(19), // uint, reduce: Value
			reduce(19), // string, reduce: Value
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(34), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			shift(36), // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(10), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(37), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(39), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(7), // ␚, reduce: Combination
			reduce(7), // (, reduce: Combination
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			reduce(12), // (, reduce: Operand
			nil,        // define
			reduce(12), // ), reduce: Operand
			nil,        // lambda
			reduce(12), // boolean_t, reduce: Operand
			reduce(12), // boolean_f, reduce: Operand
			reduce(12), // identifier, reduce: Operand
			reduce(12), // uint, reduce: Operand
			reduce(12), // string, reduce: Operand
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			nil,        // (
			nil,        // define
			reduce(10), // ), reduce: Formals
			nil,        // lambda
			nil,        // boolean_t
			nil,        // boolean_f
			reduce(10), // identifier, reduce: Formals
			nil,        // uint
			nil,        // string
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(41), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			nil,        // (
			nil,        // define
			reduce(15), // ), reduce: Identifier
			nil,        // lambda
			nil,        // boolean_t
			nil,        // boolean_f
			reduce(15), // identifier, reduce: Identifier
			nil,        // uint
			nil,        // string
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(5), // ␚, reduce: Define
			reduce(5), // (, reduce: Define
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(42), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(44), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			reduce(8), // (, reduce: Combination
			nil,       // define
			reduce(8), // ), reduce: Combination
			nil,       // lambda
			reduce(8), // boolean_t, reduce: Combination
			reduce(8), // boolean_f, reduce: Combination
			reduce(8), // identifier, reduce: Combination
			reduce(8), // uint, reduce: Combination
			reduce(8), // string, reduce: Combination
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(45), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			reduce(9), // ), reduce: Formals
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			reduce(9), // identifier, reduce: Formals
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(26), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			reduce(8), // ), reduce: Combination
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(14), // (
			nil,       // define
			shift(48), // )
			nil,       // lambda
			shift(20), // boolean_t
			shift(21), // boolean_f
			shift(10), // identifier
			shift(22), // uint
			shift(23), // string
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			reduce(7), // (, reduce: Combination
			nil,       // define
			reduce(7), // ), reduce: Combination
			nil,       // lambda
			reduce(7), // boolean_t, reduce: Combination
			reduce(7), // boolean_f, reduce: Combination
			reduce(7), // identifier, reduce: Combination
			reduce(7), // uint, reduce: Combination
			reduce(7), // string, reduce: Combination
		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(50), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(51), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			reduce(7), // ), reduce: Combination
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(52), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			shift(33), // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(26), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(6), // ␚, reduce: Combination
			reduce(6), // (, reduce: Combination
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			shift(26), // (
			nil,       // define
			nil,       // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(55), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			shift(56), // )
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			reduce(6), // (, reduce: Combination
			nil,       // define
			reduce(6), // ), reduce: Combination
			nil,       // lambda
			reduce(6), // boolean_t, reduce: Combination
			reduce(6), // boolean_f, reduce: Combination
			reduce(6), // identifier, reduce: Combination
			reduce(6), // uint, reduce: Combination
			reduce(6), // string, reduce: Combination
		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // (
			nil,       // define
			reduce(6), // ), reduce: Combination
			nil,       // lambda
			nil,       // boolean_t
			nil,       // boolean_f
			nil,       // identifier
			nil,       // uint
			nil,       // string
		},
	},
}
