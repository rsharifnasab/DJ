class A {
    public int a;
    private int c;
}

class B extends A {
    public int d;
     protected int b;
}

int main() {
    A a;
    B b;
    a = new A;
    a.a = 10;
    a.b = 20;
    a.c = 30;

    b = new B;
    b.a = 25;
    b.b = 40;
    b.d = 1;

    Print(b.a, " ", b.b, " ", b.d);

    return 0;
}