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
