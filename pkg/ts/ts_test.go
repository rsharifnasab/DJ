package ts

import (
	"os"
	"testing"

	"github.com/rsharifnasab/DJ/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestC1(t *testing.T) {

	source := `
#include "stdio.h"
#include <math.h>

int main(){
	const char * s = "string.h"; // shouldn't be a problem
	printf("%s\n", s);
}

`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/a.c", []byte(source), 0777)

	yaml := `
c:
  import-white:
    - stdio.h
    - math.h
  import-black:
    - strings.h
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.NoError(t, CheckSource(question, submission, "c"))
}
func TestC2(t *testing.T) {

	source := `
#include "stdio.h"
#include <math.h>
#include <string.h>
`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/a.c", []byte(source), 0777)

	yaml := `
c:
  import-black:
    - string.h
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.Error(t, CheckSource(question, submission, ""))
}

// A whitelist entry naming a PACKAGE must cover the classes imported
// from it (java.util should allow java.util.List) - exact matching
// rejected legitimate submissions.
func TestJavaWhitelistPackage(t *testing.T) {
	source := `
import java.util.List;
public class Main {
		public static void main(String[] a){
			System.out.println(List.of(1));
		}
}
`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/Main.java", []byte(source), 0777)

	yaml := `
java:
    import-white:
          - java.util
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.NoError(t, CheckSource(question, submission, "java"))
}

// A blacklist entry naming a PACKAGE must also catch classes from
// that package, not just verbatim imports.
func TestJavaBlacklistPackage(t *testing.T) {
	source := `
import java.net.Socket;
`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/Main.java", []byte(source), 0777)

	yaml := `
java:
    import-black:
          - java.net
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.Error(t, CheckSource(question, submission, "java"))
}

// Exact blacklist entries keep working.
func TestJavaBlacklistExactClass(t *testing.T) {
	source := `
import java.net.Socket;
`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/Main.java", []byte(source), 0777)

	yaml := `
java:
    import-black:
          - java.net.Socket
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.Error(t, CheckSource(question, submission, "java"))
}

func TestJava1(t *testing.T) {

	source := `
import java.util.*;
import java.net.Socket;
`
	submission := util.MakeTempfolder()
	defer os.RemoveAll(submission)
	os.WriteFile(submission+"/Main.java", []byte(source), 0777)

	yaml := `
java:
  import-white:
    - java.util
`
	question := util.MakeTempfolder()
	defer os.RemoveAll(question)
	os.WriteFile(question+"/ts.yaml", []byte(yaml), 0777)

	assert.Error(t, CheckSource(question, submission, "java"))
}
