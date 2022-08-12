# How to judge your project
+ Place your project in a folder and remember its path (solution)
+ source files should be in a folder named `src`
+ Choose the executable based on your OS (Darwin for mac and Linux for Linux)
+ run the script like this:
```bash
# for java:
./linux.out student -j ./judge-compiler-java -q ./question-compiler -s ./path/to/solution
# for python:
./linux.out student -j ./judge-compiler-python -q ./question-compiler -s ./path/to/solution

```
+ for Java, your program is being called like this:
```bash
javac -d out --release 11 -cp ".:lib/*" sources.txt
java -cp "out:lib/*" Main -i /path/to/src/file -o /path/to/write/result.asm
```
    + any jar file in the `lib` folder is included in Classpath
    + your main should be in class "Main" without any package


+ for Python, your program is being called like this:
```bash
python3 src/main.py -i /path/to/src/file -o /path/to/write/result.asm
```
    + your main.py file is called.
