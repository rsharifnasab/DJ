class Base {
    int a;
    void f(int x) {
        int a;
        this.a = x;
        a = 10;
        Print(a);
    }
    void print(){
        Print(this.a);
        Print(a);
    }
    void g(int x) {
        Print("HEELO");
    }
}

class Derived extends Base{
    int b;
    void g(int y) {
        this.print();
        Print(y);
    }
}

int main() {
    Base b;
    Derived d;
    d = new Derived;
    d.f(100);
    d.g(200);
    b = d;
    b.f(100);
    b.g(400);
    b = new Base;
    b.g(400);
}
