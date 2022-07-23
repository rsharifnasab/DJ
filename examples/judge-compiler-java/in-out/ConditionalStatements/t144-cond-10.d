
class A {
    int field;

    void init() {
        field = 22;
    }

    int get_field() {
        return field;
    }
}

int main() {
    A a;
    a = new A;
    a.init();

    if (a.get_field() < 20) Print("true");
    else Print("false");

    if (a.get_field() < 25) Print("true");
    else Print("false");
}
