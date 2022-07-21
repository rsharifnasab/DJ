int main() {
    int a;
    int b;
    int i;

    b = 0;
    for(i = 1; true; i = i + 1) {
        Print("Please enter the #", i, " number:");
        a = ReadInteger();
        if (a < 0)
            break;
        b = b + a;
    }

    Print("Sum of ", i, " items is: ", b);
}