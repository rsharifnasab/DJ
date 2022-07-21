
class Test {
    void test() {
        int t;
        t = 0;
        {
            int t;
            t = 5;
            Print(t);
        }
        Print(t);
    }
}

int main() {
    Test t;
    t = new Test;
    t.test();
}