package transpiler_test

import (
	"os"
	"quill/transpiler"
	"testing"
)

func TestTranspilerRun(t *testing.T) {
	buildDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(buildDir)

	mustCreate := func(name, content string) {
		if err := os.WriteFile(name, []byte(content), 0644); err != nil {
			t.Fatalf("failed to create %s: %v", name, err)
		}
	}

	if err := os.MkdirAll("posts", 0755); err != nil {
		t.Fatalf("failed to create posts dir: %v", err)
	}
	mustCreate("posts/0001_hello_world.md", "# Hello World\n\nThis is a test post.")
	mustCreate("styles.css", "body { background: #fff; }")
	mustCreate("index.html.tmpl", "<html><body>{{.DevMode}}</body></html>")

	transpiler.Run(false, buildDir)

	expectedFiles := []string{
		"index.html",
		"styles.css",
		"posts/_home.html",
	}
	for _, f := range expectedFiles {
		path := buildDir + "/" + f
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected %s to exist after build, got error: %v", path, err)
		}
	}

	if fi, err := os.Stat(buildDir + "/posts"); err != nil || !fi.IsDir() {
		t.Errorf("expected posts directory to exist in build output")
	}

	transpiler.Run(false, buildDir)
}
