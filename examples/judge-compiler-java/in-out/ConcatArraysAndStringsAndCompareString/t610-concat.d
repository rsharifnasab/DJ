int main()
{
    string[] ss1;
	string[] ss2;
	ss1	= NewArray(5, string);
    ss2 = NewArray(6, string);

    ss1[3] = "ss1";
    ss2[2] = "ss2";

    Print((ss2 + ss1)[9] + (ss1 + ss2)[7]);


}
