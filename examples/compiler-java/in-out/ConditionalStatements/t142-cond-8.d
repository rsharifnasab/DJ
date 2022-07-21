
int main() {
    int x;
    int y;

    y = -10;
    if ((x = y) < 0) Print("true");
    else {
        Print("false");
    }

    y = -10;
    if ((x = y) * 2 > -10) Print("true");
    else {
        Print("false");
    }
}
