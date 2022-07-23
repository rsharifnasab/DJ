
class Test {
    int field;

    void init() {
        field = 365214;
    }

    int test() {
        int field;
        field = 471569;
        Print(field);
        Print(this.field);
    }
}

int main() {
    Test t;
    t = new Test;
    t.init();

    t.test();
}
