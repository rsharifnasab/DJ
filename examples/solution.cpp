#include <iostream>

#include <sys/ioctl.h>
#include <termios.h>
#include <unistd.h>

void f()
{
    f();
}

void printMuch(void){
    long t = 900L;
    while(t--){
        printf("1\n");
    }
}

void mem_use(void){
    const long size = 1024L * 1024 * 10;
    int *p = (int*)malloc(size); // 20MB
    for (long unsigned int i = 0; i < (size/sizeof(int)); ++i) {
        p[i] = i;
    }
    long milli_seconds = 1000;
    usleep(1000 * milli_seconds); // 1s
}

int main()
{
    int i;
    std::cin >> i;
    int* p = NULL;
    switch (i) {
    case 1: // time limit
        while (1)
            ;
        break;
    case 2: // non-zero exit code
        mem_use();
        return 0;
    case 3: // runtime error
        return 1 / 0;
    case 4: // runtime error again
        return *(p);
    case 5: // stackoverflow
        f();
    case 6: // print much
        printMuch();
        break;
    default: // ok
        std::cout << "input : " << i << std::endl;
    }
    return 0;
}
