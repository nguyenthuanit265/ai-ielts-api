package models

type SpeakingSetIeltsPartRequest struct {
	CurrentPart  string `json:"current_part"`
	QuestionPool string `json:"question_pool"`
}

type SpeakingSetIeltsPartResponse struct {
	CurrentPart  string `json:"current_part"`
	QuestionPool string `json:"question_pool"`
}

type SpeakingQuestionResponse struct {
	IeltsQuestion string `json:"ielts_question"`
}

type SpeakingFormatAnswerRequest struct {
	Text string `json:"text"`
}

type SpeakingFormatAnswerResponse struct {
	FormattedText string `json:"formatted_text"`
}

type SpeakingSubmitAnswerRequest struct {
	UserAnswer    string `json:"user_answer"`
	IeltsQuestion string `json:"ielts_question"`
	TimeAndWps    string `json:"time_and_wps,omitempty"`
}

type SpeakingSubmitAnswerResponse struct {
	ExampleAnswer   string `json:"example_answer,omitempty"`
	UserAnswer      string `json:"user_answer,omitempty"`
	CorrectedAnswer string `json:"corrected_answer,omitempty"`
	SynonymsList    string `json:"synonyms_list,omitempty"`
	Feedback        string `json:"feedback,omitempty"`
}
