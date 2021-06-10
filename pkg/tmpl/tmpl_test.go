package tmpl

import (
	"testing"
)

func TestFileWalk(t *testing.T) {
	path := "/home/jasoet/Document/Template/vanilla-gcloud-compute-template"
	result, err := TemplateScan(path, IgnoreGit(), "tmpl")
	if err != nil {
		panic(err)
	}

	t.Logf("Result Path: %s, Extension: %s, Filter: %s \n", result.RootPath, result.Extension, result.FilterName)

	for key, value := range result.TemplateMap {
		t.Logf("Key: %s, is Template %v", key, value.IsTemplate)
	}

	// lookup template

	for _, value := range result.TemplateList() {
		t.Logf("Template: %s, is available: %v", value, result.Template.Lookup(value) != nil)
	}

	data := map[string]string{}
	executeResult, err := result.Execute(data)
	if err != nil {
		panic(err)
	}

	for key, value := range executeResult {
		t.Logf("Key: %s, value: %v", key, value)
	}

	err = result.ExecuteToPath(data, "/tmp/randomalpha")
	if err != nil {
		panic(err)
	}

}
