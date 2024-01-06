package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

const (
	HeaderFilePath = "./Blog/header.html"
	FooterFilePath = "./Blog/footer.html"
)

func convertMarkdownToHTML(mdPath string) (string, error) {
	mdContent, err := os.ReadFile(mdPath)
	if err != nil {
		return "", err
	}

	// Replace [[name]] with [name.html](name.html)
	mdContent = replaceCustomLinks(mdContent)

	htmlContent := blackfriday.Run(mdContent)

	// Read content of header.html
	headerContent, err := os.ReadFile(HeaderFilePath)
	if err != nil {
		return "", err
	}

	// Read content of footer.html
	footerContent, err := os.ReadFile(FooterFilePath)
	if err != nil {
		return "", err
	}

	// Extract the file name without extension
	fileName := strings.TrimSuffix(filepath.Base(mdPath), filepath.Ext(mdPath))

	// Replace the placeholder "{{TITLE}}" with the file name in the header content
	headerContent = bytes.ReplaceAll(headerContent, []byte("{{TITLE}}"), []byte(fileName))

	// Combine header, footer, and HTML content
	finalHTMLContent := append(headerContent, htmlContent...)
	finalHTMLContent = append(finalHTMLContent, footerContent...)

	return string(finalHTMLContent), nil
}

// func replaceCustomLinks(content []byte) []byte {
// 	re := regexp.MustCompile(`\[\[([^\]]+)\]\]`)
// 	return re.ReplaceAllFunc(content, func(match []byte) []byte {
// 		linkText := string(match[2 : len(match)-2]) // Extracting the content inside [[ ]]
// 		link := fmt.Sprintf("[%s.html](%s.html)", linkText, linkText)
// 		return []byte(link)
// 	})
// }

func saveHTMLFile(mdPath, htmlContent, blogDir string) error {
	htmlFileName := strings.TrimSuffix(filepath.Base(mdPath), filepath.Ext(mdPath)) + ".html"
	htmlPath := filepath.Join(blogDir, htmlFileName)
	return os.WriteFile(htmlPath, []byte(htmlContent), 0644)
}

func main() {
	blogDir := "Blog" // Name of the folder to store HTML files
	if err := os.Mkdir(blogDir, 0755); err != nil && !os.IsExist(err) {
		fmt.Printf("Error creating Blog directory: %s\n", err)
		os.Exit(1)
	}

	directoryLinks := make(map[string][]string)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %s\n", path, err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == ".md" {
			mdPath := path
			htmlContent, err := convertMarkdownToHTML(mdPath)
			if err != nil {
				fmt.Printf("Error converting %s to HTML: %s\n", mdPath, err)
				return nil
			}

			err = saveHTMLFile(mdPath, htmlContent, blogDir)
			if err != nil {
				fmt.Printf("Error saving HTML file for %s: %s\n", mdPath, err)
				return nil
			}

			dir := filepath.Dir(mdPath)
			base := filepath.Base(mdPath)
			directoryLinks[dir] = append(directoryLinks[dir], base)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directories: %s\n", err)
		os.Exit(1)
	}

	createIndexHTML(directoryLinks, blogDir)
	// Change the working directory to "Blog"
	err = os.Chdir(blogDir)
	if err != nil {
		fmt.Printf("Error changing directory to %s: %s\n", blogDir, err)
		promptToExit()
	}
	// Run Netlify deploy command
	err = runNetlifyDeployCommand()
	if err != nil {
		fmt.Printf("Error running Netlify deploy command: %s\n", err)
		promptToExit()
	}

	promptToExit()
}

func promptToExit() {
	fmt.Println("Press Enter to exit.")
	fmt.Scanln()
}

func createIndexHTML(directoryLinks map[string][]string, blogDir string) {
	// Read content of header.html for index file
	headerContent, err := os.ReadFile(HeaderFilePath)
	if err != nil {
		fmt.Printf("Error reading %s: %s\n", HeaderFilePath, err)
		return
	}

	// Read content of footer.html for index file
	footerContent, err := os.ReadFile(FooterFilePath)
	if err != nil {
		fmt.Printf("Error reading %s: %s\n", FooterFilePath, err)
		return
	}
	// Replace the placeholder "{{TITLE}}" with the file name in the header content
	headerContent = bytes.ReplaceAll(headerContent, []byte("{{TITLE}}"), []byte("BLOG"))

	// Combine header content with index-specific content
	indexContent := append(headerContent, []byte("<pre>by Marko Veselinovic</pre>")...)

	for dir, files := range directoryLinks {
		indexContent = append(indexContent, []byte("<h2 class='accordion-item' onclick='toggleFunction(this)'>"+dir+"<span class='arrow'>▲</span></h2><ul class='accordion-content'>")...)
		for _, file := range files {
			htmlFileName := file[:len(file)-3] + ".html"
			link := filepath.Join(htmlFileName)
			indexContent = append(indexContent, []byte("<li><a href='"+link+"'>"+file+"</a></li>")...)
		}
		indexContent = append(indexContent, []byte("</ul>")...)
	}

	indexContent = append(indexContent, footerContent...)
	//indexContent = append(indexContent, []byte("</body></html>")...)

	err = os.WriteFile(filepath.Join(blogDir, "index.html"), indexContent, 0644)
	if err != nil {
		fmt.Printf("Error creating index.html: %s\n", err)
	}
}
