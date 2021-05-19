package cliClient

import (
	"encoding/json"
	"net/http"
)

type Metadata struct {
	CliVersion      string `json:"cliVersion"`
	Os              string `json:"os"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
}

type CreateEvaluationRequest struct {
	CliId    string    `json:"cliId"`
	Metadata *Metadata `json:"metadata"`
}

func (c *CliClient) CreateEvaluation(request *CreateEvaluationRequest) (int, error) {
	httpRes, err := c.httpClient.Request(http.MethodPost, "/cli/evaluation/create", request, nil)
	if err != nil {
		return 0, err
	}

	var res = &struct {
		EvaluationId int `json:"evaluationId"`
	}{}
	err = json.Unmarshal(httpRes.Body, &res)
	if err != nil {
		return 0, err
	}

	return res.EvaluationId, nil
}

type Match struct {
	FileName string `json:"fileName"`
	Path     string `json:"path"`
	Value    string `json:"value"`
}

type EvaluationResult struct {
	Passed  bool `json:"passed"`
	Results struct {
		Matches    []*Match `json:"matches"`
		Mismatches []*Match `json:"mismatches"`
	} `json:"results"`
	Rule struct {
		ID             int    `json:"defaultRuleId"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		FailSuggestion string `json:"failSuggestion"`
	} `json:"rule"`
}

type EvaluationResponse struct {
	Results []*EvaluationResult `json:"results"`
	Status  string              `json:"status"`
}

type FileConfiguration struct {
	FileName       string                   `json:"fileName"`
	Configurations []map[string]interface{} `json:"configurations"`
}

type EvaluationRequest struct {
	EvaluationId int                  `json:"evaluationId"`
	Files        []*FileConfiguration `json:"files"`
}

func (c *CliClient) RequestEvaluation(request *EvaluationRequest) (*EvaluationResponse, error) {
	res, err := c.httpClient.Request(http.MethodPost, "/cli/evaluate", request, nil)
	if err != nil {
		return &EvaluationResponse{}, err
	}

	var evaluationResponse = &EvaluationResponse{}
	err = json.Unmarshal(res.Body, &evaluationResponse)
	if err != nil {
		return &EvaluationResponse{}, err
	}

	return evaluationResponse, nil
}

type UpdateEvaluationValidationRequest struct {
	EvaluationId   int       `json:"evaluationId"`
	InvalidFiles   []*string `json:"failedFiles"`
	StopEvaluation bool
}

func (c *CliClient) UpdateEvaluationValidation(request *UpdateEvaluationValidationRequest) error {
	_, err := c.httpClient.Request(http.MethodPost, "/cli/evaluation/validation/k8s", request, nil)
	if err != nil {
		return err
	}

	return nil
}
