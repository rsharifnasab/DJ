class vector {
    int actual_size;
    int fixed_size;
    int[] a;

    void init() {
        fixed_size = 1;
        actual_size = 0;
        a = NewArray(1, int);
    }

    void push(int x) {
        if(fixed_size == actual_size) {
            int[] b;
            int i;

            b = NewArray(2 * fixed_size, int);
            for (i=0; i<fixed_size; i=i+1) {
                b[i] = a[i];
            }
            a = b;
            fixed_size = 2 * fixed_size;
        }
        a[actual_size] = x;
        actual_size = actual_size + 1;
    }

    void pop() {
        if(actual_size == 0) return ;
        actual_size = actual_size - 1;
    }

    int get(int p) {
        if (p >= actual_size) return -1;
        return a[p];
    }

    int[] get_vector() {
        int[] res;
        int i;
        res = NewArray(actual_size, int);
        for (i=0; i<actual_size; i=i+1)
            res[i] = a[i];
        return res;
    }

}

void DFS(int v) {
    int[] neigh;
    int i;

    mark[v] = true;

    neigh = adj[v].get_vector();

    for (i=0; i<neigh.length(); i=i+1) {
        int u;
        u = neigh[i];
        if(!mark[u])
            DFS(u);
    }

    topol.push(v);

}

void SFD(int v) {
    int[] neigh;
    int i;

    mark[v] = true;

    neigh = rev[v].get_vector();

    for (i=0; i<neigh.length(); i=i+1) {
        int u;
        u = neigh[i];

        if(!mark[u])
            SFD(u);
    }

    Print(v);

}

vector[] adj;
vector[] rev;

bool[] mark;

vector topol;

int main() {
    int n;
    int m;
    int i;
    int res;
    int[] rtopol;
    int[] temp;

    topol = new vector;
    topol.init();

    n = ReadInteger();
    m = ReadInteger();

    adj = NewArray(n, vector);
    rev = NewArray(n, vector);

    mark = NewArray(n, bool);

    for (i=0; i<n; i=i+1)
        mark[i] = false;

    for (i=0; i<n; i=i+1)
    {
        adj[i] = new vector;
        rev[i] = new vector;
        adj[i].init();
        rev[i].init();
    }

    for (i=0; i<m; i=i+1) {
        int u;
        int v;
        u = ReadInteger();
        v = ReadInteger();

        adj[u].push(v);
        rev[v].push(u);

        //adj[v].push(u);

    }

    for (i=0; i<n; i=i+1) {
        if(!mark[i]) {
            DFS(i);
        }
    }

    for (i=0; i<n; i=i+1)
        mark[i] = false;


    temp = topol.get_vector();
    rtopol = NewArray(n, int);

    for (i=0; i<n; i=i+1) {
        rtopol[i] = temp[n - i - 1];
        Print(rtopol[i]);
    }

    for (i=0; i<n; i=i+1) {
        int v;
        v = rtopol[i];
        if(!mark[v]) {
            Print("New SCC component");
            SFD(v);
            Print("--------------");
         }
    }
}

