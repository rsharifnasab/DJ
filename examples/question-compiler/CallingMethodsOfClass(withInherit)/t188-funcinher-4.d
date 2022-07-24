
class Base {
    int counter;
    void action() {}
    void base_print() {
        Print("Base.print()");
    }
}

class A extends Base {
    void action() {
        base_print();
        print();
    }

    void print() {
        Print("a.print()");
    }
}

int main() {
    Base[] bs;
    // not that BS thou :)

    bs = NewArray(1, Base);
    bs[0] = new A;

    bs[0].action();
}