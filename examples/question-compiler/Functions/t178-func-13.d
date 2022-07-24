
double func1(double x) {
    return x * 2.0;
}

double func2(double x) {
    return x + 1.0;
}

int main() {
    Print(func1(func2(func1(func1(func2(2.5))))));
    Print(func2(func2(func1(func1(func2(1.2))))));
}
