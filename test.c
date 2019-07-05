#include <stdio.h>
#include <math.h>
//#include "test2.c"
#include <ncurses.h>

int main () {
	// printf("Hello world!\n");
	// printf("Hello world3\n");
	init_pair(1,2,0);
	hello();	
	printf("%lf", sqrt(42));
	return 0;
}
