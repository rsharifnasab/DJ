int main() {
    double d1;
    double d2;
    int i1;
    int i2;
    bool b1;
    bool b2;

    double rd;
    int ri;

    d1 = 1.5741691;
    d2 = -0.1245714;

    i1 = ReadInteger();
    i2 = ReadInteger();

    b1 = i1 > i2;
    b2 = itod(i2) >= d1;
    rd = btoi(b1) * d1 + btoi(b2) * d2;
    ri = dtoi(rd);
    ri = ri * i1 - i2 * i2 * ri;
    rd = ri % 2 + rd * 0.5;
    Print(b1 || b2);
    Print(ri);
    Print(rd);

    i1 = ReadInteger();
    i2 = ReadInteger();

    b1 = i1 > i2;
    b2 = itod(i2) >= d1;
    rd = btoi(b1) * d1 + btoi(b2) * d2;
    ri = dtoi(rd);
    ri = ri * i1 - i2 * i2 * ri;
    rd = ri % 2 + rd * 0.5;
    Print(b1 || b2);
    Print(ri);
    Print(rd);

    i1 = ReadInteger();
    i2 = ReadInteger();

    b1 = i1 > i2;
    b2 = itod(i2) >= d1;
    rd = btoi(b1) * d1 + btoi(b2) * d2;
    ri = dtoi(rd);
    ri = ri * i1 - i2 * i2 * ri;
    rd = ri % 2 + rd * 0.5;
    Print(b1 || b2);
    Print(ri);
    Print(rd);

    i1 = ReadInteger();
    i2 = ReadInteger();

    b1 = i1 > i2;
    b2 = itod(i2) >= d1;
    rd = btoi(b1) * d1 + btoi(b2) * d2;
    ri = dtoi(rd);
    ri = ri * i1 - i2 * i2 * ri;
    rd = ri % 2 + rd * 0.5;
    Print(b1 || b2);
    Print(ri);
    Print(rd);
}
