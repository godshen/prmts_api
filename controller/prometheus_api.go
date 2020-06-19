package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	_commonTimeout = 10 * time.Second
	_responseTimeApi = "sum(delta(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage'}[15s]))/sum(delta(istio_request_duration_seconds_count{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage'}[15s])) * 1000"
	_supplyPodApi = "count(sum(rate(container_cpu_usage_seconds_total{image!='',namespace='jx-test',pod_name=~'cproductpage.*'}[10s])) by (pod_name, namespace))"
	_requestSuccessTotalApi = "sum(istio_requests_total{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'2.*'})"
	_requestFailTotalApi = "sum(istio_requests_total{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'5.*'})"
	_serviceTimeFailTotalApi = "sum(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'5.*'})"
	_serviceTimeAvailableTotalApi = "sum(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'2.*'})"
) 

var prometheusClient *http.Client

func newPrometheusClient()  {
	prometheusClient = &http.Client{}
}

func doRequest(apiUrl string) (msg string, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		msg = fmt.Sprintf("NewRequest Failed, URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
		return
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx, cancel := context.WithTimeout(context.TODO(), _commonTimeout)
	defer cancel()
	req = req.WithContext(ctx)
	var response *http.Response
	response, err = prometheusClient.Do(req)
	if err != nil {
		msg = fmt.Sprintf("Http Get Failed: URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
		return
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		msg = fmt.Sprintf("Http Get Failed: URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
	}
	return
}