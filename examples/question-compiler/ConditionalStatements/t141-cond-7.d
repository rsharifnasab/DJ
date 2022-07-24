
int main() {
    int x;
    int y;

    x = ReadInteger(); y = ReadInteger();
    if (x > y) {
        Print("true");
    }

    x = ReadInteger(); y = ReadInteger();
    if (x < y) {
        Print("true");
    }

    x = ReadInteger(); y = -ReadInteger();
    if (y < 0 + x) {
        Print("true");
    }

    x = ReadInteger(); y = x;
    if (x == y) {
        Print("true");
    }
}
