# compile and run junit tests

junit5 standalone test runner can produce xml report with specifying report dir

```bash
java -jar junit-platform-console.jar --cp "." --scan-class-path \
    --fail-if-no-tests  --reports-dir=reportdir --disable-banner
```

parse this xml with:
+ bash [+](https://unix.stackexchange.com/questions/83385/parse-xml-to-get-node-value-in-bash-script) [+](https://unix.stackexchange.com/questions/83385/parse-xml-to-get-node-value-in-bash-script)
ugly parsers in pure bash or add dependancy

+ python [+](https://www.edureka.co/blog/python-xml-parser-tutorial/)
do they have python? we should force them for java?!

+ java (manually select first line of xml string or ..)


# in java 
+ custom test runner in java

+ run only one test at time [+](https://stackoverflow.com/questions/9288107/run-single-test-from-a-junit-class-using-command-line)

+ add listener [+](https://howtodoinjava.com/junit/how-to-add-listner-in-junit-testcases/)  [+](https://stackoverflow.com/questions/27038351/getting-list-of-tests-from-junit-command-line)




# (unread) junit 5 stuff
+ custom extensions [+](https://www.mscharhag.com/java/junit5-custom-extensions)

+ architecture [+](https://nipafx.dev/junit-5-architecture-jupiter/#JUnit-5)

+ user guide [+](https://junit.org/junit5/docs/current/user-guide/#writing-tests-assumptions)




