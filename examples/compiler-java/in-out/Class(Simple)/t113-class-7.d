
class A {
    void a(Z z) {
        Print("entering method a of A");
        z.c();
        Print("exiting method a of A");
    }
}

class Z {
    void c() {
        Print("inside method c of Z");
    }
}

int main() {
    A a;
    Z z;
    a = new A;
    z = new Z;

    a.a(z);
}
