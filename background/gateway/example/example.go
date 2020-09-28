package example

type ExampleContent struct {
	Content string
}

func GetExampleContent() []*ExampleContent {
	contentList := make([]*ExampleContent, 0)
	contentList = append(contentList, &ExampleContent{"hello world!"})
	contentList = append(contentList, &ExampleContent{"hello golang!"})
	contentList = append(contentList, &ExampleContent{"hello musehub!"})

	return contentList
}
