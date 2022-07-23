
class Test {
    int field;

    void init() {
        field = 851236;
    }

    int get_field() {
        return field;
    }
}

int main() {
    Test t;
    t = new Test;
    t.init();

    Print(t.get_field());
}
