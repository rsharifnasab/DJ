
int main() {
    bool b;

    b = 5 < 6;
    if (b) Print("true");
    else Print("false");

    b = 5 == 6 && true;
    if (b) Print("true");
    else Print("false");

    b = 5 != 6;
    if (b) Print("true");
    else Print("false");

    b = 5 >= 6;
    if (b) Print("true");
    else Print("false");

    b = 5 < 6 || false;
    if (b) Print("true");
    else Print("false");

    b = 5 < 6 && false;
    if (b) Print("true");
    else Print("false");
}
