
int[] g;

void test() {
    Print(g[4]);
    Print(g[3]);
    Print(g[2]);
    Print(g[1]);
    Print(g[0]);
}

int main() {
    g = NewArray(5, int);
    g[0] = 1;
    g[1] = 2;
    g[2] = 4;
    g[3] = 8;
    g[4] = 16;
    test();
}
