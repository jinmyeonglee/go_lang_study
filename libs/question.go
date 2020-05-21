package main

type Question struct {
	question string

	choices []string

	selected int
}

func new_question() *Question {
	q = Question{}

}
