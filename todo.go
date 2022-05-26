package main

type Todo struct {
	Text string
}

func todoFactory(text string) Todo {
	return Todo{
		Text: text,
	}
}
