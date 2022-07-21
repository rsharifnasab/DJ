class A{
    int a;
    void set_a(int a) {
        this.a = a;
    }
    int get_a(){
        return a;
    }
}

int main() {
    int i;
    string s1;
    string s2;
    A[] arr;
    arr = NewArray(10, A);
    for (i=0; i<10; i=i+1)
    {
        arr[i] = new A;
        arr[i].set_a(i);
    }
    for (i=0; i<10; i=i+1)
        Print(arr[i].get_a());
    s1 = ReadLine();
    s2 = ReadLine();
    if(s1 == s2)
        Print("YES");
    else
        Print("NO");
}