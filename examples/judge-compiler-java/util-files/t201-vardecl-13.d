
int main() {
    int t;
    int c;
    c = 0;
    for ( t = 5; t > 0; t = t - 1 ) {
        // this one may get a lil' tricky :D
        int t;
        t = 5;
        Print(t);

        // avoid infinite loop in case of faulty compiler
        c = c + 1;
        if (c > 10) break;
    }
    Print(t);
}
