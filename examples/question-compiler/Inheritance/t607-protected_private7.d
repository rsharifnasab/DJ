class A {
    public int a;
}

class B extends A {

}

class C extends B {

}

int main() {
    C c;
    c = new C;
    c.a = 5;

    Print(c.a);


    return 0;
}