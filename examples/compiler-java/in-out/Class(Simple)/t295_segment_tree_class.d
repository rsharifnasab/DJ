int inf;

class segment_tree{
    int n;
    int[] seg;

    void init(int n){
        int i;
        this.n=n;
        seg=NewArray(4*n+10, int);
        for(i=0;i<4*n+10;i=i+1){
            seg[i]=inf;
        }
    }
    void update2(int id,int val,int s,int e,int p){
        int mid;
        if(e-s<=1){
            seg[p]=val;
            return;
        }
        mid=(s+e)/2;
        if(id<mid)
            update2(id,val,s,mid,2*p);
        else
            update2(id,val,mid,e,2*p+1);

        if(seg[2*p]<seg[2*p+1])
            seg[p]=seg[2*p];
        else
            seg[p]=seg[2*p+1];
    }
    void update(int id,int val){
        update2(id,val,0,n,1);
    }
    int get_min2(int l,int r,int s,int e,int p){
        int mid;
        int r1;
        int r2;
        if(l>=e || r<=s){
            return inf;
        }
        if(l<=s && r>=e){
            return seg[p];
        }
        mid=(s+e)/2;
        r1=get_min2(l,r,s,mid,2*p);
        r2=get_min2(l,r,mid,e,2*p+1);
        if(r1<r2)
            return r1;
        return r2;
    }
    int get_min(int l,int r){
        return get_min2(l,r,0,n,1);
    }
}

int main() {
    segment_tree my_st;
    int n;
    int q;
    int i;

    inf = 1000000000;
    my_st=new segment_tree;
    n=ReadInteger();
    q=ReadInteger();

    my_st.init(n);
    for(i=0;i<q;i=i+1){
        //Query: ? l r -> min(a[l],...,a[r])  , + x val  -> a[x]=val;
        string t;
        int l;
        int r;
        t=ReadLine();
        l=ReadInteger();
        r=ReadInteger();

        if(t=="?"){
            Print(my_st.get_min(l-1,r));
        }else
            my_st.update(l-1,r);
    }
}
