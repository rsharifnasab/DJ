
int factorial_helper(int n, int acc) {
    if (n == 0) return acc;
    else return factorial_helper(n - 1, acc * n);
}

int factorial(int n) {
    return factorial_helper(n, 1);
}

int main() {
    Print(factorial(5));
}
