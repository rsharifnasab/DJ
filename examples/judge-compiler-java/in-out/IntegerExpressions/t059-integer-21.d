int main() {
    int x;
    int y;
    int z;

    int r;

    x = ReadInteger();
    y = ReadInteger();
    z = ReadInteger();

    r = x * (y + 5) * (z - x * y) * 3 / ((z + 0) - x) % y;
    Print(r);

    r = (x - (y + z) % (y % z));
    Print(r);

    r = (x + (y * z)) * (z - y) / (z);
    Print(r);

}
