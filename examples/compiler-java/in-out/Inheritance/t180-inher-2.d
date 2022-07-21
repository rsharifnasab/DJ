
class A {
    int inher;
}

class B extends A {
    void test() {
        inher = -1;
        Print(inher);
    }
}

int main() {
    A a;
    B b;
    b = new B;
    b.test();
}
