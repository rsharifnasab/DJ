int main(){
    int i;
    int[] a;
    int len;

    len = 5;
    a = NewArray(len, int);
    a[0] = 1;
    a[1] = 2;
    a[2] = a[1];
    a[1] = 10;
    a[3] = 4;
    a[4] = 5;
    for(i = 0; i < len; i = i+1){
        Print(a[i]);
    }
    Print(a.length());
}