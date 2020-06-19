package router

import (
	"control/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func GetServer() *negroni.Negroni {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.HomeHandler)
	
	// service GET
	r.Handle("/services/{id}/state", controller.ResponseHandler(controller.GetServiceRunStatus)).Methods("GET")
	r.Handle("/services/{id}/stime", controller.ResponseHandler(controller.GetServiceRunTime)).Methods("GET")
	r.Handle("/services/{id}/rtime", controller.ResponseHandler(controller.GetServiceResponseTime)).Methods("GET")
	r.Handle("/services/{id}/number", controller.ResponseHandler(controller.GetServiceNumbers)).Methods("GET")
	r.Handle("/services/{id}/stotal", controller.ResponseHandler(controller.GetRequestSuccessTotal)).Methods("GET")
	r.Handle("/services/{id}/ftotal", controller.ResponseHandler(controller.GetRequestFailTotal)).Methods("GET")
	r.Handle("/services/{id}/available", controller.ResponseHandler(controller.GetServiceAvailableTime)).Methods("GET")
	r.Handle("/services/{id}/unavailable", controller.ResponseHandler(controller.GetServiceUnavailableTime)).Methods("GET")
	
	// service PATCH
	r.Handle("/services/{id}/state/{is_running}", controller.ResponseHandler(controller.PatchServiceRunStatus)).Methods("PATCH")
	
	// service POST
	r.Handle("/services/logUp", controller.ResponseHandler(controller.ServiceLogUp)).Methods("POST")
	// 待实现
	r.Handle("/services/{id}/rtime/{response_time}", controller.ResponseHandler(controller.PostServiceResponseTime)).Methods("POST")
	r.Handle("/services/{id}/number/{pod_number}", controller.ResponseHandler(controller.PostServiceNumbers)).Methods("POST")
	r.Handle("/services/{id}/stotal/{request_success_total}", controller.ResponseHandler(controller.PostRequestSuccessTotal)).Methods("POST")
	r.Handle("/services/{id}/ftotal/{request_fail_total}", controller.ResponseHandler(controller.PostRequestFailTotal)).Methods("POST")
	r.Handle("/services/{id}/available/{service_time_available}", controller.ResponseHandler(controller.PostServiceAvailableTime)).Methods("POST")
	r.Handle("/services/{id}/unavailable/{service_time_unavailable}", controller.ResponseHandler(controller.PostServiceUnavailableTime)).Methods("POST")

	// algorithm GET
	r.Handle("/algorithms/{id}/state", controller.ResponseHandler(controller.GetAlgorithmRunStatus)).Methods("GET")
	r.Handle("/algorithms/{id}/reliability", controller.ResponseHandler(controller.GetReliabilityIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/availability", controller.ResponseHandler(controller.GetAvailabilityIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/stability", controller.ResponseHandler(controller.GetStabilityIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/cost", controller.ResponseHandler(controller.GetCostIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/elasticity", controller.ResponseHandler(controller.GetElasticityIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/oscillation", controller.ResponseHandler(controller.GetOscillationIndex)).Methods("GET")
	r.Handle("/algorithms/{id}/sla", controller.ResponseHandler(controller.GetSlaSatisfactionIndex)).Methods("GET")

	// algorithm PATCH
	r.Handle("/algorithms/{id}/state/{is_running}", controller.ResponseHandler(controller.PatchAlgorithmRunStatus)).Methods("PATCH")

	// algorithm POST
	r.Handle("/algorithms/logUp", controller.ResponseHandler(controller.AlgorithmLogUp)).Methods("POST")
	// 待实现
	r.Handle("/algorithms/{id}/reliability/{reliability}", controller.ResponseHandler(controller.PostReliabilityIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/availability/{availability}", controller.ResponseHandler(controller.PostAvailabilityIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/stability/{stability}", controller.ResponseHandler(controller.PostStabilityIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/cost/{cost}", controller.ResponseHandler(controller.PostCostIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/elasticity/{elasticity}", controller.ResponseHandler(controller.PostElasticityIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/oscillation/{oscillation}", controller.ResponseHandler(controller.PostOscillationIndex)).Methods("POST")
	r.Handle("/algorithms/{id}/sla/{sla_satisfaction}", controller.ResponseHandler(controller.PostSlaSatisfactionIndex)).Methods("POST")

	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
