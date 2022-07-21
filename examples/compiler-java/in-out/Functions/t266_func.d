int f(int x){
    if ( x == 0 ){
        Print(1);
        return 1;
    }
    return f(x-1) + x;
}

int main()  {
    Print(f(5));
}