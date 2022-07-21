
class Test {
    int a;
    void init() {
        a = 1;
    }
    void test() {
        int a;
        a = 4;
        Print(a);
        Print(this.a);
    }
}

int main() {
    Test t;
    t = new Test;
    t.init();
    t.test();
}