package linguist

import (
	"path/filepath"
	"testing"
)

var (
	appgolang       = filepath.Join("_testdata", "app-golang")
	appEmptydirPath = filepath.Join("_testdata", "app-emptydir")
	appPythonfile   = filepath.Join("_testdata", "app-python", "app.py")
)

func TestProcessFile(t *testing.T) {
	output, err := ProcessFile(appPythonfile)
	if err != nil {
		t.Error("expected detect to pass")
	}
	if output != "Python" {
		t.Errorf("expected output == 'Python', got '%s'", output)
	}
}

func TestProcessDir(t *testing.T) {
	output, err := ProcessDir(appgolang)
	if err != nil {
		t.Error("expected detect to pass")
	}
	if output[0].Language != "Go" {
		t.Errorf("expected output == 'Go', got '%s'", output[0].Language)
	}

	// test with a bad dir
	if _, err := ProcessDir(filepath.Join("/dir", "does", "not", "exist")); err == nil {
		t.Error("expected err when running detect with a dir that does not exist")
	}

	// test an application that should fail detection
	output, _ = ProcessDir(appEmptydirPath)
	if len(output) != 0 {
		t.Errorf("expected no languages detected, got '%d'", len(output))
	}
}

func TestGitAttributes(t *testing.T) {
	testCases := []struct {
		path         string
		expectedLang string
	}{
		{filepath.Join("_testdata", "app-duck"), "Duck"},
		{filepath.Join("_testdata", "app-vendored"), "Python"},
		{filepath.Join("_testdata", "app-not-vendored"), "HTML"},
		{filepath.Join("_testdata", "app-documentation"), "Python"},
		{filepath.Join("_testdata", "app-generated"), "Python"},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			output, err := ProcessDir(tc.path)
			if err != nil {
				t.Errorf("expected ProcessDir() to pass, got %s", err)
			}
			if output[0].Language != tc.expectedLang {
				t.Errorf("expected output == '%s', got '%s'", tc.expectedLang, output[0].Language)
			}
		})
	}
}

//TestDirectoryIsIgnored checks to see if directory paths such as 'docs/' are ignored from being classified by linguist when added to the "ignore" list.
func TestDirectoryIsIgnored(t *testing.T) {
	path := filepath.Join("_testdata", "app-documentation")
	// populate isIgnored
	ProcessDir(path)
	ignorePath := filepath.Join(path, "docs")
	if !isIgnored(ignorePath) {
		t.Errorf("expected dir '%s' to be ignored", ignorePath)
	}
}

func TestGetAlias(t *testing.T) {
	testcases := map[string]string{
		"maven pom": "Java",
		"mAvEn POM": "Java",
		"c#":        "csharp",
		"Python":    "Python",
	}

	for packName, expectedAlias := range testcases {
		alias := Alias(&Language{Language: packName})
		if alias.Language != expectedAlias {
			t.Errorf("Expected alias to be '%s', got '%s'", expectedAlias, alias.Language)
		}
	}
}
