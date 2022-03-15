@0 = global [9 x i8] c"result %d"

define i32 @main() {
0:
	%1 = add i32 2, 3
	%2 = sub i32 3, 2
	%3 = sub i32 %2, 1
	%4 = add i32 1, %1
	%5 = add i32 %4, %3
	%6 = mul i32 %5, 100
	%7 = sdiv i32 %6, 100
	%8 = bitcast [9 x i8]* @0 to i8*
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
