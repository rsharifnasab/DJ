package run

var CompileCommand map[string]string = map[string]string{
	// with the help of https://www.hackerearth.com/challenges/college/code-with-c/faq/
	//"c":   "gcc -std=gnu99 -fno-asm -Dasm=error -w -O2 -fomit-frame-pointer -lm %v -o compiled/a.out",
	//"cpp": "g++ -std=c++0x -fno-asm -Dasm=error -w -O2 -fomit-frame-pointer -lm %v -o compiled/a.out",

	"c":   "gcc -std=c11   -w -O2 -fomit-frame-pointer -lm -lpthread %v -o compiled/a.out",
	"cpp": "g++ -std=c++17 -w -O2 -fomit-frame-pointer -lm  -pthread %v -o compiled/a.out",
	// https://stackoverflow.com/questions/2679885/using-pthread-in-c

	"java": "javac --release 11 -d compiled %v",
}
