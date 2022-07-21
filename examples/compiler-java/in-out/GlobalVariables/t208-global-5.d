
string g;

void test() {
    Print(g);
}

int main() {
    g = "global for the win";
    test();
}
