
class A {

    int counter;

    void init() {
        counter = 0;
    }

    A a() {
        Print("entering method a of A");
        counter = counter + 1;
        print();
    }

    void print() {
        Print("Value of counter is: ", counter);
    }
}

class Z {
    void c(A a) {
        Print("entering method c of Z");
        a.a();
    }
}


int main() {
    A a;
    Z z;
    a = new A;
    z = new Z;
    a.init();

    z.c(a);
}
