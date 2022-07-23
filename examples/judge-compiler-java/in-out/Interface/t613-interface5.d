interface X {
	void print();
}

class A implements X {
	int a;
	void print() {
		Print(this.a);
	}
}

class B implements X {
	int a;
	void print() {
		Print(this.a);
	}
}

void callPrintX(X x) {
	x.print();
}

int main(){
	A a;
	B b;
	a = new A;
	b = new B;
	a.a = 2;
	b.a = 3;
	callPrintX(a);
	callPrintX(b);
}