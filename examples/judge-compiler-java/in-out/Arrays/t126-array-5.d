
int main() {
    int[] ints;

    ints = NewArray(5, int);

    ints[0] = 1;
    ints[1] = 2;
    ints[2] = 3;
    ints[3] = 4;
    ints[4] = 0;

    Print(ints[ints[ints[0]]]);
}
