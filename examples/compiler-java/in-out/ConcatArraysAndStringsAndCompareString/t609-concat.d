int main()
{
	int[] a;
	int[] b;
	int[] d1;
	int[] d2;
	
    a = NewArray(6, int);
    b = NewArray(7, int);

    d1 = NewArray(5, double);
    d2 = NewArray(10, double);

    d1[3] = 6.25;



    a[3] = ReadInteger();

    Print((d2 + d1)[(b + a)[10]]);

}
