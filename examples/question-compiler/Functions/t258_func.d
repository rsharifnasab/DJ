int f(int x){
    if ( x <= 2 ){
        return 1;
    }
    return f(x-1) + f(x-2);
}

int main()  {
    Print(f(3));
    Print(f(6));
}