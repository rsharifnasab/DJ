int main()
{

    int u;
    int v;

    u = 5;

    for(u = 7;u < 10;u = u + 100 - 97)
    {
        Print(2 * u);
        for(v = 2;v < 4;v = v + 1)
        {
            Print(v);
            continue;
            v = u;
        }
        continue;
        u = u - 1;
    }

}
