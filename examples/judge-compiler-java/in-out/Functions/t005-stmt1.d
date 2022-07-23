
int test(int a, int b) {
    return a * b;
}

int main() {
    int a;
    int b;

    a = ReadInteger();
    b = ReadInteger();

    Print(test(a, b));
}
