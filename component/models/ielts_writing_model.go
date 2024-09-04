package models

type WritingSubmitRequest struct {
	UserAnswer    string `json:"user_answer,omitempty"`
	IeltsQuestion string `json:"ielts_question,omitempty"`
}

type WritingSubmitResponse struct {
	ExampleAnswer       string `json:"example_answer,omitempty"`
	Scores              string `json:"scores,omitempty"`
	CorrectedAnswer     string `json:"corrected_answer,omitempty"`
	UserAnswer          string `json:"user_answer,omitempty"`
	CurrentQuestionType string `json:"current_question_type,omitempty"`
	Feedback            string `json:"feedback,omitempty"`
}

type WritingGetQuestionResponse struct {
	IeltsQuestion       string `json:"ielts_question,omitempty"`
	CurrentQuestionType string `json:"current_question_type,omitempty"`
}

type WritingGetQuestionTypeResponse struct {
	CurrentQuestionType string `json:"current_question_type,omitempty"`
}

type WritingSubmitRequestV2 struct {
	Prompt         string `json:"prompt"`
	UserSubmission struct {
		TaskType   string `json:"task_type"`
		Question   string `json:"question"`
		UserAnswer string `json:"user_answer"`
	} `json:"user_submission"`
	Parameters struct {
		FeedbackType         string `json:"feedback_type"`
		IncludeExampleAnswer bool   `json:"include_example_answer"`
	} `json:"parameters"`
}

type WritingSetIeltsTaskRequest struct {
	CurrentTask int `json:"current_task"`
}

type WritingSetIeltsTaskResponse struct {
	CurrentTask int `json:"current_task"`
}

type WritingQuestionTypeResponse struct {
	CurrentQuestionType string `json:"current_question_type"`
}

type WritingQuestionResponse struct {
	IeltsQuestion       string `json:"ielts_question"`
	CurrentQuestionType string `json:"current_question_type"`
}
