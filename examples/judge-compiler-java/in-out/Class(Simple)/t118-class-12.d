
class A {

    int counter;

    void init() {
        counter = 0;
    }

    A a(A a1) {
        Print("entering method a of A");
        a1.counter = a1.counter + 1;
        a1.print();
        print();
    }

    void print() {
        Print("Value of counter is: ", counter);
    }
}


int main() {
    A a1;
    A a2;
    a1 = new A;
    a2 = new A;
    a1.init();
    a2.init();

    a1.a(a2);
}
