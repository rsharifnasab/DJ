
int main() {
    int[][] table;
    table = NewArray(5, int[]);

    table[0] = NewArray(5, int);
    table[1] = NewArray(5, int);
    table[2] = NewArray(5, int);
    table[3] = NewArray(5, int);
    table[4] = NewArray(5, int);

    table[0][0] = ReadInteger();
    table[0][1] = ReadInteger();
    table[0][2] = ReadInteger();
    table[0][3] = ReadInteger();
    table[0][4] = ReadInteger();
    table[1][0] = ReadInteger();
    table[1][1] = ReadInteger();
    table[1][2] = ReadInteger();
    table[1][3] = ReadInteger();
    table[1][4] = ReadInteger();
    table[2][0] = ReadInteger();
    table[2][1] = ReadInteger();
    table[2][2] = ReadInteger();
    table[2][3] = ReadInteger();
    table[2][4] = ReadInteger();
    table[3][0] = ReadInteger();
    table[3][1] = ReadInteger();
    table[3][2] = ReadInteger();
    table[3][3] = ReadInteger();
    table[3][4] = ReadInteger();
    table[4][0] = ReadInteger();
    table[4][1] = ReadInteger();
    table[4][2] = ReadInteger();
    table[4][3] = ReadInteger();
    table[4][4] = ReadInteger();

    Print(table[0][0]);
    Print(table[0][1]);
    Print(table[1][1]);
    Print(table[1][0]);
    Print(table[4][3]);
    Print(table[2][1]);
}
