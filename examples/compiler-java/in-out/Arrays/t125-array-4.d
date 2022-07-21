
int main() {
    int[] ints;

    ints = NewArray(5, int);

    ints[0] = ReadInteger();
    ints[1] = ReadInteger();
    ints[2] = ReadInteger();
    ints[3] = ReadInteger();
    ints[4] = ReadInteger();

    Print(ints[ints[1] % ints[2] % ints[4] % ints[0] % ints[3] % 5]);
}
