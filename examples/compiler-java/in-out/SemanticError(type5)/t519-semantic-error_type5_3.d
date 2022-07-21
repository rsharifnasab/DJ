double f () {
	return 2.3;
}

int main() {
	int [] a;
	a = NewArray(10, int);
	a[f()] = 23;
	Print("Hi");
}