
class A {
    int fA;
}

class B extends A {
	int fB;
    void test() {
        fA = 1;
        fB = 2;
        Print(fA);
        Print(fB);
    }
}

int main() {
    B b;
    b = new B;
    b.test();
}
