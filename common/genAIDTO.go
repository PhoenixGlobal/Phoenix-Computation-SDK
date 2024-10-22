package common

// ReqGenImg genImage request parameters
type ReqGenImg struct {
	UserToken            string  `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
	Prompt               string  `json:"prompt"`                            //prompt, not empty
	NegativePromt        string  `json:"negative_prompt"`                   //tokens, not empty
	NSteps               int     `json:"n_steps" example:"20"`
	GuidanceScale        float64 `json:"guidance_scale" example:"8"` // float, 0 to 10
	Seed                 int     `json:"seed" example:"1"`           // int, >0
	PipeName             string  `json:"pipe_name" example:"txt2img"`
	LoraPathList         string  `json:"lora_path_list" example:"https://xxxx"`
	LoraWeightList       string  `json:"lora_weight_list" example:"0.8"`
	TextualInversionPath string  `json:"textual_inversion_path" example:"without"`
	LoraScale            float64 `json:"lora_scale" example:"0.8"`
	Size                 string  `json:"size" example:"512^*^512"`
	ClipSkip             int     `json:"clip_skip" example:"1"`
	InitImage            string  `json:"init_image" example:"https://xxx.jpg"`
	MaskImage            string  `json:"mask_image" example:"https://xxx.jpg"`
}

// ReqGenImgSDXL gen SDXL image request parameters
type ReqGenImgSDXL struct {
	Prompt          string  `json:"prompt" example:"a blue dog"`       //prompt, not empty
	NegativePromt   string  `json:"negative_prompt"`                   //tokens, not empty
	UserToken       string  `json:"user_token" example:"xxxxxxxxxxxx"` //user_token, not empty
	NSteps          int     `json:"n_steps" example:"20"`
	GuidanceRescale float64 `json:"guidance_rescale" example:"8"`  // float, 0 to 10
	HighNoiseFrac   float64 `json:"high_noise_frac" example:"0.8"` // float, 0 to 1
	Seed            int     `json:"seed" example:"1"`              // int, >0
}

// ReqGenImgFlux gen Flux image request parameters
type ReqGenImgFlux struct {
	UserToken string `json:"user_token" example:"xxxxxxxxxxxx"`   //user_token, not empty
	Prompt    string `json:"prompt" example:"a flying airplane:"` //prompt, not empty
}
