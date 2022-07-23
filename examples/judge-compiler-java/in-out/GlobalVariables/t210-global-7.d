
int g;

void test1() {
    g = 16;
}

void test2() {
    Print(g);
}

int main() {
    int g;
    g = 1;
    Print(g);
    test1();
    Print(g);
    test2();
    Print(g);
}
