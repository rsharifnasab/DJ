class A {

    public int getC() {
        return this.c;
    }

    public int setC(int newC) {
        this.c = newC;
    }

    protected void func() {
        Print("Main");
    }
}

class B extends A {

    protected void func() {
        Print("Override");
    }
}

int main() {
    A a;
    B b;
    b = new B;
    a = b;


    a.func();

    return 0;
}