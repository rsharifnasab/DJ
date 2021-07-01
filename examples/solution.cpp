#include <iostream>
int main()
{
    int i;
    std::cin >> i;
    int *p = NULL;
    switch (i) {
        case 1: // time limit 
            while(1)
                ;
            break;
        case 2: // non-zero exit code
            return 1;
        case 3: // runtime error
            return 1/0;
        case 4: // runtime error again
            return *(p);
        default: // ok
            std::cout << "input : " << i << std::endl;
    }
    return 0;
}
