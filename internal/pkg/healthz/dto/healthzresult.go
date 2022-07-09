package dto

type HealthzResult struct {
	Status  string               `json:"status" example:"ok" enums:"ok,error"`
	Details HealthzResultDetails `json:"details"`
}

type HealthzResultDetails struct {
	Db HealthzResultDetail `json:"db"`
}

type HealthzResultDetail struct {
	Status string `json:"status" example:"ok" enums:"ok,error"`
}
