define i32 @main() {
0:
	%1 = fadd float 2.0, 3.0
	%2 = fsub float 3.0, 2.0
	%3 = fsub float %2, 1.0
	%4 = fadd float 1.0, %1
	%5 = fadd float %4, %3
	%6 = fmul float %5, 100.0
	%7 = fdiv float %6, 100.0
	ret i32 0
}
