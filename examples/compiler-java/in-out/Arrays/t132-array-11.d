
int main() {
    int[][][][][][][][][][][][] multi_table;

    multi_table = NewArray(1, int[][][][][][][][][][][]);
    multi_table[0] = NewArray(1, int[][][][][][][][][][]);
    multi_table[0][0] = NewArray(1, int[][][][][][][][][]);
    multi_table[0][0][0] = NewArray(1, int[][][][][][][][]);
    multi_table[0][0][0][0] = NewArray(1, int[][][][][][][]);
    multi_table[0][0][0][0][0] = NewArray(1, int[][][][][][]);
    multi_table[0][0][0][0][0][0] = NewArray(1, int[][][][][]);
    multi_table[0][0][0][0][0][0][0] = NewArray(1, int[][][][]);
    multi_table[0][0][0][0][0][0][0][0] = NewArray(1, int[][][]);
    multi_table[0][0][0][0][0][0][0][0][0] = NewArray(1, int[][]);
    multi_table[0][0][0][0][0][0][0][0][0][0] = NewArray(1, int[]);
    multi_table[0][0][0][0][0][0][0][0][0][0][0] = NewArray(1, int);
    multi_table[0][0][0][0][0][0][0][0][0][0][0][0] = 42;
    // xD

    Print(multi_table[0][0][0][0][0][0][0][0][0][0][0][0]);
}
