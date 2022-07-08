package dto

type HealthzResult struct {
	Status  string               `json:"status"`
	Details HealthzResultDetails `json:"details"`
}

type HealthzResultDetails struct {
	Db HealthzResultDetail `json:"db"`
}

type HealthzResultDetail struct {
	Status string `json:"status"`
}
