# Dejagnu
+ a program for test automation, but generic
+ uses TCL and expect (in debian named expect-dev package)
+ wikipedia page: [+](https://en.wikipedia.org/wiki/DejaGnu)
+ [gnu's manual](https://www.gnu.org/software/dejagnu/manual/index.html)
+ [quick start guide](https://web.archive.org/web/20120322145747/http://www.kalycito.com/documents/Quick_Start_DejaGNU.pdf)
+ mostly for interactive programs like gdb
+ use regex for pattern match output
+ create a log file (int text) and xml file from the test suit.
```tcl
set test "local another test"
spawn bash -c "echo 'salam donya'"
expect {
    -re "salam donya" { pass "$test" }
        default { fail "$test" }
}
```

+ another good tutorial with example [+](https://www.math.utah.edu/docs/info/dejagnu_1.html)


