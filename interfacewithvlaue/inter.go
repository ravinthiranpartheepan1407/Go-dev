package main

import "fmt"

type Exam interface {
	evaluateExam()
}

type Writer string
type Evaluater string

func (w Writer) evaluateExam() {
	fmt.Println("Write the exam")
}

func (e Evaluater) evaluateExam() {
	fmt.Println("Evaluate the exam paper")
}

func evaluateExamPapers(evaluate []Exam) {
	fmt.Println("Evaluate exam")
	for i := 0; i < len(evaluate); i++ {
		eval := evaluate[i]
		fmt.Printf("Process: (%v)", eval)
		eval.evaluateExam()
	}
}

func main() {
	evals := []Exam{Writer("Ravinthiran"), Evaluater("Suren")}
	evaluateExamPapers(evals)
}
