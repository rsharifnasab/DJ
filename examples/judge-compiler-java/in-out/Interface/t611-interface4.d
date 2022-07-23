interface X {
	void print();
}

class A implements X {
	void print() {
		Print("A");
	}
}

class B implements X {
	void print() {
		Print("B");
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
	callPrintX(a);
	callPrintX(b);
}