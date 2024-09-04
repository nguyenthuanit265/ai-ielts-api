package utils

// Constant datetime format
const (
	DATE_TIME_FORMAT_YYYY_MM_DDT_HH_MM_SS string = "2006-01-02T15:04:05"
	DATE_TIME_FORMAT_YYYY_MM_DD           string = "2006-01-02"
	DATE_TIME_FORMAT_YYYY_MM_DD_HH_mm_SS  string = "2006-01-02 15:04:05"
	DATE_TIME_FORMAT_YYYYMMDDTHHmmSS      string = "20060102150405" // yyyyMMddHHmmss
	DATE_TIME_FORMAT_DD_MM_YYYY           string = "02/01/2006"
	DATE_TIME_FORMAT_YYYYMMDD             string = "20060102"
)

const (
	SQL_DATE_TIME_FORMAT_YYYY_MM_DD_HH_MM_SS string = "yyyy-mm-dd hh24-mi-ss"
	SQL_DATE_TIME_FORMAT_YYYYMMDDHH24MMSS    string = "yyyymmddhh24miss"
)

const (
	CurrentUser string = "currentUser"
)

// Constant method
const (
	MethodPost   string = "POST"
	MethodGet    string = "GET"
	MethodPut    string = "PUT"
	MethodDelete string = "Delete"
)

// Constant role
const (
	RoleCodeHeadDepartment string = "role_head_department"
	RoleCodeManager        string = "role_manager"
)

// Constant permission
const (
	PermissionProductView   string = "product.view"
	PermissionProductCreate string = "product.create"
	PermissionProductUpdate string = "product.update"
	PermissionProductDelete string = "product.delete"
	PermissionProductExport string = "product.export"
)

const (
	PermissionCategoryView   string = "category.view"
	PermissionCategoryCreate string = "category.create"
	PermissionCategoryUpdate string = "category.update"
	PermissionCategoryDelete string = "category.delete"
)

const (
	PermissionManufacturerView   string = "manufacturer.view"
	PermissionManufacturerCreate string = "manufacturer.create"
	PermissionManufacturerUpdate string = "manufacturer.update"
	PermissionManufacturerDelete string = "manufacturer.delete"
)

const (
	ExtensionCsv   string = ".csv"
	ExtensionExcel string = ".xlsx"
)

// Model ChatGPT
const (
	Gpt4TurboPreview  = "gpt-4-turbo-preview"
	Gpt4VisionPreview = "gpt-4-vision-preview"
	Gpt4              = "gpt-4"
	Gpt3Dot5          = "gpt-3.5"
	Gpt3Dot5Turbo     = "gpt-3.5-turbo"
)

const (
	GptCompletions = "https://api.openai.com/v1/chat/completions"
)

const (
	PROMT_WRITING_SUBMIT_1 = "Evaluate the following IELTS writing submission. Provide feedback on the user's response, suggest improvements, and offer an example of an ideal answer."
	PROMT_WRITING_SUBMIT_2 = "Evaluate the following IELTS writing submission. Provide an example of an ideal answer."
	PROMT_WRITING_SUBMIT_3 = `
			I need help editing a response for an IELTS General Training Task 1 question. 
			
			%s

			Could you help by:
			- Correcting any grammatical or spelling mistakes?
			- Suggesting improvements for better coherence and cohesion, perhaps with more effective linking words or transitions?
			- Enhancing the vocabulary to be more varied and precise, particularly around expressing apologies and suggesting solutions?
			- Ensuring the tone remains appropriately formal and polite throughout?
			- Checking that the response fully addresses the task, including all necessary explanations and suggestions?
			
			I’m also interested in any specific sentences or phrases that could be reworded for a stronger impact or clearer expression. Thank you!

`

	PROMT_SPEAKING_SUBMIT_1 = "Evaluate the following IELTS speaking submission. Provide feedback on the user's response, suggest improvements, and offer an example of an ideal answer."
	PROMT_SPEAKING_SUBMIT_2 = "Evaluate the following IELTS speaking submission. Provide an example of an ideal answer."
)
