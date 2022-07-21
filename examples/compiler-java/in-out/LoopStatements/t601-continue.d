int main() {
    int x;

    for(x=1;x<16;x = x * 2)
    {
        Print(x);
        continue;
        x = x * 2;
    }

    return 0;
}
