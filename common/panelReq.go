package common

// ReqJobDetailByUser job detail request parameters
type ReqJobDetailByUser struct {
	JobID string `json:"jobID" example:"65573937761644746890524415817246033372115830624452900918578145256037962235724"` //job id, non-null
}

// ReqSubListByUser submission list by user
type ReqSubListByUser struct {
	*ReqPage
	JobID string `json:"jobID" example:"65573937761644746890524415817246033372115830624452900918578145256037962235724"` //job id, non-null
}

// ReqCreateJobByInput request parameter for creating input job
type ReqCreateJobByInput struct {
	JobName     string   `json:"jobName" example:"matrix addition job"`                                      //job name
	Parties     []string `json:"parties" example:"11111111@gmail.com,22222222@gmail.com,33333333@gmail.com"` //participants
	Computation string   `json:"computation" example:"matrix addition"`                                      //computation type. One of "privacy comparison", "matrix addition", and "matrix multiplication"
	Input       string   `json:"input" example:"[[18,34,18.2]]"`                                             //data used for computing
}

// ReqCreateJobByFile request parameter for creating upload job
type ReqCreateJobByFile struct {
	JobName     string `json:"jobName" example:"Millionaires' Problem"`   //job name
	Computation string `json:"computation" example:"logistic regression"` //computation type. One of "decision tree", "logistic regression", and "svm"
}

// ReqFillData fill data request parameter
type ReqFillData struct {
	JobID string `json:"jobID" example:"65573937761644746890524415817246033372115830624452900918578145256037962235724"` //job id, non-null
	Input string `json:"input" example:"239483279483298918390"`                                                         //input data, non-null
}

// ReqDeleteJobByUser request parameters
type ReqDeleteJobByUser struct {
	JobID string `json:"jobID" example:"65573937761644746890524415819246033372115830624452900918578145256037962235724"` //job id, non-null
}

// ReqDownloadDataset download dataset request parameters
type ReqDownloadDataset struct {
	JobID       string `josn:"jobID" example:"65573937761644746890524415817246033372115830624452900918578145256037962235724"` //job id, non-null
	DatasetType string `json:"datasetType" example:"feature,target,testing"`                                                  //dataset type，one of feature, target, testing
}

// ReqQueryFlops query flops
type ReqQueryFlops struct {
	JobName        string  `json:"jobName" example:"CNN Classification"` //job name
	TrainData      string  `json:"x_train"`
	TrainLable     string  `json:"y_train"`
	Epoch          uint    `json:"epoch" example:"2"`
	Model          string  `json:"model" example:"CNN,LSTM"`
	TrainShard     float64 `json:"train_shard" example:"0.8"`
	ScaleRatio     int     `json:"scale_ratio"`
	InputWidth     int     `json:"input_width"`
	LabelWidth     int     `json:"label_width"`
	Shift          int     `json:"shift"`
	SequenceStride int     `json:"sequence_stride"`
	SamplingRate   int     `json:"sampling_rate"`
}

// ReqCreateAIJob the request parameters of createAIJob
type ReqCreateAIJob struct {
	JobName        string  `json:"jobName" example:"CNN Classification"` //job name
	Computation    string  `json:"computation" example:"Distributed AI"` //computation type
	TrainData      string  `json:"x_train"`
	TrainLable     string  `json:"y_train"`
	BatchSize      uint    `json:"batchSize" example:"64"`
	Epoch          uint    `json:"epoch" example:"2"`
	Model          string  `json:"model" example:"CNN,LSTM"`
	TrainShard     float64 `json:"train_shard" example:"0.8"`
	CcdCost        float64 `json:"ccd_cost"` //ccd cost
	ScaleRatio     int     `json:"scale_ratio"`
	InputWidth     int     `json:"input_width"`
	LabelWidth     int     `json:"label_width"`
	Shift          int     `json:"shift"`
	SequenceStride int     `json:"sequence_stride"`
	SamplingRate   int     `json:"sampling_rate"`
}
