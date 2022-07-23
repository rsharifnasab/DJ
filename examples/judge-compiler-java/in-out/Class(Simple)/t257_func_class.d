class Fen {
    int[] fen;
    int N;
    void init(int N) {
        fen = NewArray(N + 10, int);
        this.N = N;
        return ;
    }

    void upd(int p, int x) {
        for(p=p+1; p<N; p=p+big(p)) {
            fen[p] = fen[p] + x;
            //Print(p, " debug ");
        }
    }

    int get(int p) {
        int res;
         res = 0;
        for (; p>0; p=p-big(p))
            res = res + fen[p];
        return res;
    }

}

int big(int x) {
    int p;
    p = 1;
    while(x % 2 == 0) {
        p = p * 2;
        x = x / 2;
    }
    return p;
}

int main() {
    Fen f;
    int n;
    int q;
    int i;

    n = ReadInteger();
    q = ReadInteger();

    //Print(and(n, q));

    f = new Fen;
    f.init(n);

    while(itob(q)) {
        int tp;
        q = q - 1;
        tp = ReadInteger();
        if(tp == 0){
            int l;
            int r;
            Print("l, r");
            l = ReadInteger();
            r = ReadInteger();
            l = l - 1;
            Print(f.get(r) - f.get(l));
        }else {
            int p;
            int x;
            Print("p, x");
            p = ReadInteger();
            x = ReadInteger();
            p = p - 1;
            f.upd(p, x);
        }
    }

}