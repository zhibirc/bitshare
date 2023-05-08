package controllers

import "net/http"

type HealthController struct {
}

func (cnt *HealthController) GetHealthStatus(res http.ResponseWriter, req *http.Request) {
	// TODO: implement
}
