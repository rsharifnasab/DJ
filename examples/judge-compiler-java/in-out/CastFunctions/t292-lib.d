int main(){
    Print(btoi(true));
    Print(btoi(false));
    Print(btoi(!true));
    Print(itob(2));
    Print(itob(0));
    Print(itob(1-1));
    Print(dtoi(2.5));
    Print(dtoi(2.1));
    Print(dtoi(2.6));
    Print(dtoi(2.3)*3);
    Print(itod(5));
    Print(itod(10));
    Print(itod(5)/2.5);
    Print(itod(5)/4.0);
    Print( dtoi(itod(btoi(true)) + 2.3) );
}
