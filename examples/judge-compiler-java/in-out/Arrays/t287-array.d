int main(){
    int i;
    int j;
    int k;
    int[][][] a;

    a = NewArray(3, int[][]);
    for(i = 0; i < 3; i = i+1){
        a[i] = NewArray(i+1, int[]);
    }
    for(i = 0; i < 3; i = i+1){
        for(j = 0; j <= i; j = j+1){
            a[i][j] = NewArray(3, int);
            for(k = 0; k < 3; k = k+1){
                a[i][j][k] = k;
            }
        }
    }

    for(i = 0; i < 3; i = i+1){
        Print("i ", i);
        for(j = 0; j <= i; j = j+1){
            Print("j ", j);
            for(k = 0; k < 3; k = k+1){
                Print(a[i][j][k]);
            }
        }
    }


}