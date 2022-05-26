package todo

type Todo struct {
	Text string
}

func TodoFactory(text string) Todo {
	return Todo{
		Text: text,
	}
}
