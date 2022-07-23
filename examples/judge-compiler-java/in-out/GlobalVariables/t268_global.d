int a;

int foo(int a,  bool c) {
    if (c){
	  Print("ok");
      return a + 2;
    }
    else
      Print(a, " wacky.");
   return 18;
}

void main() {
	int b;

	a = 10;
	b = a/2;
	foo(a,  true);
	foo(b + 2, a <= b);
	foo(foo(3, true && false),  !true);
}