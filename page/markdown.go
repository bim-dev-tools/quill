package page

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func ParseMarkdownToHTML(md string) string {
	opts := html.RendererOptions{
		Flags: html.CommonFlags,
	}
	renderer := html.NewRenderer(opts)
	return string(markdown.ToHTML([]byte(md), nil, renderer))
}
