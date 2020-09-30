package example

import "fmt"

type ExampleContent struct {
	Content string
}

func GetExampleContent(name string) []*ExampleContent {
	contentList := make([]*ExampleContent, 0)
	contentList = append(contentList, &ExampleContent{"hello world!"})
	contentList = append(contentList, &ExampleContent{"hello golang!"})
	contentList = append(contentList, &ExampleContent{"hello musehub!"})
	contentList = append(contentList, &ExampleContent{fmt.Sprint("from: ", name)})

	return contentList
}
