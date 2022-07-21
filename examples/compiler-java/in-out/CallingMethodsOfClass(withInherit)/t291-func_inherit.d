class Base {
    void f(){
        Print("base");
    }
}

class Derived extends Base {
    void f(){
        Print("derived");
    }
}

class Derived2 extends Derived{
    void g(){
        Print("derived2");
    }
}

void H(Base b1, Base b2) {
    b1.f();
    b2.f();
}

void H2(Base b1, Derived2 d2) {
    b1.f();
    d2.g();
}

int main() {
    Derived2 d;
    Base b;
    Base x;
    b = new Base;
    b.f();
    b = new Derived;
    b.f();
    b = new Derived2;
    b.f();

    d = new Derived2;

    x = new Base;
    H(x, b);
    H2(b, d);
    H2(x, d);
}