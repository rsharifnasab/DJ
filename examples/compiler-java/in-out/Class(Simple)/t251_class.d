class A{
    string a;
    void set_a(string a) {
        this.a = a;
    }
    string get_a(){
        return a;
    }
    bool comp(A oth){
        if (a == oth.get_a())
            return true;
        return false;
    }
}

int main() {
    double res;
    int a;
    int b;
    a = ReadInteger();
    b = ReadInteger();
    res = itod(a) / itod(b);
    Print(dtoi(res) == a/b);
}