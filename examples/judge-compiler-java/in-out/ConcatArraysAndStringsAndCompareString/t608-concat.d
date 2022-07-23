int main()
{
	int[] a;
	int[] b;
    a = NewArray(8, int);
    b = NewArray(4, int);

    b[0] = 4;

    Print((a+b)[8]);

}
