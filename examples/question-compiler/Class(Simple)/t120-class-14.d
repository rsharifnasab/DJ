
class A {
    Z a(Z z, int c) {
        Print("entering method a of A");
        return z.c(c + 1);
    }

    Z b(Z z, int c) {
        Print("entering method b of A");
        return z.c(c * 2);
    }
}

class Z {

    int counter;

    void init() {
        counter = 0;
    }

    Z c(int x) {
        Print("inside method c of Z");
        counter = counter + x;
        return this;
    }

    void print() {
        Print("Value of Z counter is: ", counter);
    }
}

int main() {
    A a;
    Z z;
    a = new A;
    z = new Z;
    z.init();

    a.b(a.a(a.b(z, 1), 2), 3);
    z.print();
}
