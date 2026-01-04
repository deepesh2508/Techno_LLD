package interfaces

type OpenAIClient interface {
	EvaluateAnswer(questionID int, userAnswer string) (score float64, feedback string, err error)
}
