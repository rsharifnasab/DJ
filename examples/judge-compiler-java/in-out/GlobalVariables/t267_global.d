int x;
int[] y;
int[] t;
void anotherFunction(int n){
	Print("global value of x is ", x);
}
void assignValues(int n){
	bool x;
	x = true;
	Print("local value of x is ", x);
}
void main(){
    y = NewArray(100, int);
    t = NewArray(100, int);
	x = 100;
	assignValues(99);
	anotherFunction(99);
}