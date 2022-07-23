
int func1(int x) {
    return x * 2;
}

int func2(int x) {
    return x + 1;
}

int main() {
    Print(func1(func2(func1(func1(func2(ReadInteger()))))));
    Print(func2(func2(func1(func1(func2(ReadInteger()))))));
}
