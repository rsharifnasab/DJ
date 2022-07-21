
int factorial(int x) {
    if (x <= 1) return 1;
    return x * factorial(x - 1);
}

int main() {
    Print(factorial(5));
}
