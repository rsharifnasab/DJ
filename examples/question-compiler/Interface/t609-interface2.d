interface x {
    public bool zz();
}

class A implements x {
    int b;

    public bool zz() {
        return true;
    }
}


int main() {
    A a;
    bool b;
    a = new A;
    b = a.zz();

    Print(b);

    return 0;
}