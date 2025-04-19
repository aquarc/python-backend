package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
)

func main() {
	input := `
<h1><u>Hello</u>, World!</h1>
<p>This is <u>an</u> example.</p>
<p>Another <u>example</u> here.</p>
<p><u>This</u> is <u>a</u> test.</p>
`

	conv := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
		),
	)
	// Perform simple string replacements for <u> and </u>
    markdown := strings.ReplaceAll(input, "<u>", "[UNDERLINE]")
	markdown = strings.ReplaceAll(markdown, "</u>", "[END]")

    markdown, err := conv.ConvertString(markdown)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(markdown)
}

