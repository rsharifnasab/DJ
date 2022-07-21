
class A {
    int a(int x) {
        Print("entering method a of A");
        return x + 1;
    }

    int b(int x) {
        Print("entering method b of A");
        return x * 2;
    }
}

int main() {
    A a;
    a = new A;

    Print(a.a(a.b(a.a(a.b(1)))));
}
