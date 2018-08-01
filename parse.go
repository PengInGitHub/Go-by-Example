package main

type Problem struct {
	Question string
	Answer   string
}

func Parse() []Problem {
	var problems []Problem
	for _, line := range GetLines() {
		p := Problem{
			line[0],
			line[1],
		}
		problems = append(problems, p)
	}

	return problems
}

func GetLines() [][]string {
	pls := [][]string{
		{"Best language?", "PHP"},
		{"First language?", "Java"},
		{"Favourit? ", "Go"},
	}
	return pls
}
