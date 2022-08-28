package judge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var examplePath = "../../examples"

func createSubmission(lang string) *Submission {
	return NewSubmission(
		examplePath+"/judge-simple",
		examplePath+"/question-add",
		lang,
		examplePath+"/sol-add-"+lang,
		"",
	)
}

func TestCheckReq(t *testing.T) {
	submission := createSubmission("c")
	submission.initSandboxWithoutTest()
	submission.checkReq()
}

func TestNoOfTest(t *testing.T) {
	submission := createSubmission("c")
	submission.initSandboxWithoutTest()
	submission.prepareForTestGroup(submission.exploreTestGroups()[0])
	n := submission.currentGroupTestCount()
	assert.Equal(t, 2, n)
}

func TestSpecificLang(t *testing.T) {
	for _, lang := range []string{"c", "java", "python"} {
		submission := createSubmission(lang)
		res := submission.RunSuite()
		assert.Len(t, res.TestGroupResults, 1)
		assert.Equal(t, res.TestGroupResults[0].TestCount, 2)
		assert.Len(t, res.TestGroupResults[0].TestResults, 2)
		assert.True(t, res.TestGroupResults[0].TestResults[0].isPassed())
		assert.True(t, res.TestGroupResults[0].TestResults[1].isPassed())
	}
}

func TestAutoDetectLang(t *testing.T) {
	for _, lang := range []string{"c", "java", "python"} {
		submission := NewSubmission(
			examplePath+"/judge-simple",
			examplePath+"/question-add",
			"", // let it detect
			examplePath+"/sol-add-"+lang,
			"",
		)
		res := submission.RunSuite()
		assert.Len(t, res.TestGroupResults, 1)
		assert.Equal(t, res.TestGroupResults[0].TestCount, 2)
		assert.Len(t, res.TestGroupResults[0].TestResults, 2)
		assert.True(t, res.TestGroupResults[0].TestResults[0].isPassed())
		assert.True(t, res.TestGroupResults[0].TestResults[1].isPassed())
	}
}

func TestWA(t *testing.T) {
	submission := NewSubmission(
		examplePath+"/judge-simple",
		examplePath+"/question-add",
		"", // let it detect
		examplePath+"/sol-wrong-c",
		"",
	)
	res := submission.RunSuite()
	assert.Len(t, res.TestGroupResults, 1)
	assert.Equal(t, res.TestGroupResults[0].TestCount, 2)
	assert.Len(t, res.TestGroupResults[0].TestResults, 2)
	assert.False(t, res.TestGroupResults[0].TestResults[0].isPassed())
	assert.True(t, res.TestGroupResults[0].TestResults[0].Run)
	assert.True(t, res.TestGroupResults[0].TestResults[0].Wrong)
	assert.False(t, res.TestGroupResults[0].TestResults[0].TimedOut)

	assert.False(t, res.TestGroupResults[0].TestResults[1].isPassed())
	assert.True(t, res.TestGroupResults[0].TestResults[1].Run)
	assert.True(t, res.TestGroupResults[0].TestResults[1].Wrong)
	assert.False(t, res.TestGroupResults[0].TestResults[1].TimedOut)
}
