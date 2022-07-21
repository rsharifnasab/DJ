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
    int res;
    int i;
    int j;
    int n;
    string s1;
    string s2;
    A[] arr;
    n = ReadInteger();
    arr = NewArray(n, A);
    for (i=0; i<n; i=i+1){
        arr[i] = new A;
        arr[i].set_a(ReadLine());
    }

    res = 0;
    for (i=0; i<n; i=i+1){
        for(j=i+1; j<n; j=j+1){
            if(arr[i].comp(arr[j])) res = res + 1;
        }
    }
    Print(res);
}