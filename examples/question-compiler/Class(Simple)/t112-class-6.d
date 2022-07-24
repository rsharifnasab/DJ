
class A {
    void a() {
        Print("entering method a of A");
        b();
        Print("exiting method a of A");
    }

    void b() {
        Print("inside method b");
    }
}

int main() {
    A a;
    a = new A;

    a.a();
}
