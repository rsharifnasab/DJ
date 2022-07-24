
class A {
    int inher;
}

class B extends A {

}

class C extends B {
    void test() {
        inher = -2;
        Print(inher);
    }
}

int main() {
    A a;
    B b;
    C c;
    c = new C;
    c.test();
}
