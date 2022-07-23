
class A {
    void test() {
        Print("A.test()");
    }
}

class B extends A {
	void test() {
        Print("B.test()");
    }
}


int main() {
    A a;
    B b;
    a = new A;
    a.test();
    b = new B;
    b.test();
    a = b;
    a.test();
}
