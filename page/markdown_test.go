package page_test

import (
	"quill/page"
	"strings"
	"testing"
)

func contains(s, substr string) bool {
	return substr == "" || strings.Contains(s, substr)
}

func TestMarkdownParsing(t *testing.T) {
	cases := []struct {
		name         string
		md           string
		wantContains []string
	}{
		{
			name:         "Header and paragraph",
			md:           "# Title\n\nSome text.",
			wantContains: []string{"<h1>Title</h1>", "<p>Some text."},
		},
		{
			name:         "Code block",
			md:           "```go\nfmt.Println(\"hi\")\n```",
			wantContains: []string{"<pre><code class=\"language-go\">", "fmt.Println"},
		},
		{
			name:         "List",
			md:           "- item 1\n- item 2",
			wantContains: []string{"<ul>", "<li>item 1</li>", "<li>item 2</li>"},
		},
		{
			name:         "Edge: empty input",
			md:           "",
			wantContains: []string{""},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			html := page.ParseMarkdownToHTML(tc.md)
			for _, want := range tc.wantContains {
				if want != "" && !contains(html, want) {
					t.Errorf("expected output to contain %q, got: %s", want, html)
				}
			}
		})
	}
}
