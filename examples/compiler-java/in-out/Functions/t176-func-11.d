
int func(int x) {
    if (x <= 1) return 1;
    return 1 + func(x - 1);
}

int main() {
    Print(func(1000));
}
