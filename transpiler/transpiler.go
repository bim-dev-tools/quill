package transpiler

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"spagen/config"
	"spagen/page"
	"strings"

	"golang.org/x/net/html"
)

//go:embed static/*
var StaticFiles embed.FS

var staticFileList = []string{
	"router.js",
	"prism.js",
	"prism.css",
}

type templateData struct {
	*config.Config
	DevMode bool
}

func Run(devMode bool, buildDir string) {

	if err := buildHomeComponent(buildDir); err != nil {
		fmt.Println("Error building home component:", err)
		return
	}

	if err := buildIndexFile(buildDir, devMode); err != nil {
		fmt.Println("Error building index file:", err)
		return
	}

	fp, err := os.ReadFile("styles.css")
	if err != nil {
		fmt.Println("Error reading styles.css:", err)
		return
	}
	stylesFile, err := os.Create(filepath.Join(buildDir, "styles.css"))
	if err != nil {
		fmt.Println("Error creating styles.css in build dir:", err)
		return
	}
	defer stylesFile.Close()
	_, err = stylesFile.Write(fp)
	if err != nil {
		fmt.Println("Error writing styles.css to build dir:", err)
		return
	}

	for _, file := range staticFileList {
		path := fmt.Sprintf("static/%s", file)
		if err := copyToBuild(path, buildDir, file); err != nil {
			fmt.Printf("Error copying %s: %v\n", file, err)
			return
		}
	}
}

func buildPostList(buildDir string) *html.Node {
	list := html.Node{
		Type: html.ElementNode,
		Data: "ul",
		Attr: []html.Attribute{
			{Key: "id", Val: "home__page-list"},
		},
	}

	files, err := os.ReadDir("./posts/")
	if err != nil {
		panic("Could not find dir")
	}

	os.MkdirAll(filepath.Join(buildDir, "posts"), os.ModePerm)
	os.Chdir("./posts")
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		filename := file.Name()
		if strings.HasPrefix(filename, "_") || !strings.HasSuffix(filename, ".md") {
			continue
		}

		mdBytes, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		htmlContent := page.ParseMarkdownToHTML(string(mdBytes))
		// Use first line or filename as title
		lines := strings.Split(string(mdBytes), "\n")
		var title string
		if len(lines) > 0 && strings.HasPrefix(lines[0], "#") {
			title = strings.TrimPrefix(lines[0], "#")
			title = strings.TrimSpace(title)
		} else {
			title = filename
		}

		outName := strings.TrimSuffix(filename, ".md") + ".html"
		postFile, err := os.Create(filepath.Join("..", buildDir, "posts", outName))
		if err != nil {
			panic(err)
		}
		defer postFile.Close()

		postFile.WriteString(htmlContent)

		appendToHomeList(outName, title, &list)
	}

	os.Chdir("..")

	body := html.Node{
		Type: html.ElementNode,
		Data: "body",
	}

	body.AppendChild(&list)

	return &body
}

func buildHomeComponent(buildDir string) error {
	body := buildPostList(buildDir)

	homeComponentPath := fmt.Sprintf("%s/posts/_home.html", buildDir)
	dir := filepath.Dir(homeComponentPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	homeComponent, err := os.Create(homeComponentPath)
	if err != nil {
		return err
	}
	defer homeComponent.Close()

	return html.Render(homeComponent, body)
}

func buildIndexFile(buildDir string, devMode bool) error {
	tmplFile := "index.html.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	fp, err := os.Create("" + buildDir + "/index.html")
	if err != nil {
		return err
	}
	defer fp.Close()

	templateData := templateData{
		Config:  config.Get(),
		DevMode: devMode,
	}
	if err = tmpl.Execute(fp, templateData); err != nil {
		return err
	}

	return nil
}

func appendToHomeList(filename string, linkText string, list *html.Node) {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	listItem := html.Node{
		Type: html.ElementNode,
		Data: "li",
	}

	linkNode := html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{Key: "href", Val: fmt.Sprintf("/#/%s", name)},
			{Key: "data-nav", Val: "true"},
		},
	}

	linkTextNode := html.Node{
		Type: html.TextNode,
		Data: linkText,
	}

	linkNode.AppendChild(&linkTextNode)
	listItem.AppendChild(&linkNode)
	list.AppendChild(&listItem)
}

func copyToBuild(embedPath string, buildDir string, outName string) error {
	data, err := StaticFiles.ReadFile(embedPath)
	if err != nil {
		return fmt.Errorf("error reading embedded file %s: %v", embedPath, err)
	}
	dstPath := filepath.Join(buildDir, outName)
	parentDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %v", parentDir, err)
	}
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("error creating file in build dir: %v", err)
	}
	defer dstFile.Close()
	_, err = dstFile.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to %s: %v", dstPath, err)
	}
	return nil
}
