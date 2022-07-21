
void sort(int[] items) {

    /* implementation of bubble sort */


    int i;
    int j;

    int n;
    n = items.length();
    for (i = 0; i < n-1; i = i + 1)
        for (j = 0; j < n - i - 1; j = j + 1)
            if (items[j] > items[j + 1]) {
                int t;
                t = items[j];
                items[j] = items[j + 1];
                items[j + 1] = t;
            }
}

int main() {
    int i;
    int j;
    int[] rawitems;
    int[] items;

    Print("Please enter the numbers (max count: 100, enter -1 to end sooner): ");

    rawitems = NewArray(100, int);
    for (i = 0; i < 100; i = i + 1) {
        int x;
        x = ReadInteger();
        if (x == -1) break;

        rawitems[i] = x;
    }

    items = NewArray(i, int);

    // copy to a more convenient location
    for (j = 0; j < i; j = j + 1) {
        items[j] = rawitems[j];
    }

    sort(items);


    Print("After sort: ");

    for (i = 0; i < items.length(); i = i + 1) {
        Print(items[i]);
    }
}
