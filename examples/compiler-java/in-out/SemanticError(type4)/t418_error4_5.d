class A{
    int a;
}
class B extends A{
    int b;   
}
void f(B b){
    Print(b.a);
}
int main(){
    A a;
    a = new A;
    a.a = 2;
    f(a);
}