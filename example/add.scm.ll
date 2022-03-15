define i32 @main() {
0:
	%1 = add i32 2, 3
	%2 = sub i32 3, 2
	%3 = sub i32 %2, 1
	%4 = add i32 1, %1
	%5 = add i32 %4, %3
	%6 = mul i32 %5, 100
	%7 = sdiv i32 %6, 100
	ret i32 0
}
