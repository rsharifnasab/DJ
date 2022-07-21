
int main() {
    int a;
    int b;

    a = ReadInteger();
    b = ReadInteger();

    Print(abs_mult(a, b));
}

int abs_mult(int a, int b) {
    int c;
    if (a > b)
        c = a - b;
    else
        c = b - a;
    return c * a * b;
}
