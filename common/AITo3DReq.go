package common

// ReqAITo3DJob request parameter for creating SkyNet 3d job
type ReqAITo3DJob struct {
	JobName         string `json:"job_name"`                        //job name
	Computation     string `json:"computation" example:"AI-to-3D (Tripo)"` //computation type
	Prompt string `json:"prompt" example:"prompt"`
	Type string `json:"type" example:"text2Model,img2Model"` //type,  text2Model or img2Model
	ImageURL string `json:"image_url"` // The ImageURL parameter will only take effect when the Type is set to img2Model. Only JPEG and PNG formats are supported
}

type ReqQueryAITo3DJobDetail struct {
	JobID string `json:"job_id"`
}

type ReqQueryAITo3DTaskModelURL struct {
	JobID string `json:"job_id"`
	TaskID string `json:"task_id"`
}

type ReqRefineModel struct {
	JobID string `json:"job_id"`
	TaskID string `json:"task_id"`
}

type ReqAnimateModel struct {
	JobID string `json:"job_id"`
	TaskID string `json:"task_id"`
	Animation string `json:"animation" example:"preset:run , preset:walk , preset:dive"`
}

type ReqConvertModel struct {
	JobID string `json:"job_id"`
	TaskID string `json:"task_id"`
	Format string `json:"format" example:"USDZ,FBX,OBJ,STL"`
}