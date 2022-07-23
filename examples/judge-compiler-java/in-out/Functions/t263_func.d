int h() {
    Print("h");
    return 1;
}

int g() {
    Print("g");
    h();
    return 1;
}

int f(){
    Print("f");
    g();
    h();
    return 3;
}

int main()  {
    f();
    g();
    h();
}