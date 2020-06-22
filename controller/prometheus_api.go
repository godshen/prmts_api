package controller

import (
	"context"
	"control/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	_timeFormat                   = "2006-01-02 15:04:05"
	_commonTimeout                = 10 * time.Second
	_lastedTime                   = 3600
	_prometheusUrl                = "http://139.9.57.167:9090/api/v1/query"
	_demandPodApi                 = "instance_predict_v1"
	_throughputApi                = "sum(rate(istio_requests_total{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage'}[30s]))"
	_responseTimeApi              = "sum(delta(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage'}[15s]))/sum(delta(istio_request_duration_seconds_count{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage'}[15s])) * 1000"
	_supplyPodApi                 = "count(sum(rate(container_cpu_usage_seconds_total{image!='',namespace='jx-test',pod_name=~'cproductpage.*'}[10s])) by (pod_name, namespace))"
	_requestSuccessTotalApi       = "sum(istio_requests_total{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'2.*'})"
	_requestFailTotalApi          = "sum(istio_requests_total{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'5.*'})"
	_serviceTimeFailTotalApi      = "sum(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'5.*'})"
	_serviceTimeAvailableTotalApi = "sum(istio_request_duration_seconds_sum{destination_workload_namespace='jx-test',reporter='destination',destination_workload='cproductpage',response_code=~'2.*'})"
)

var prometheusClient *http.Client

func NewPrometheusClient() {
	prometheusClient = &http.Client{}
}

func doRequest(apiUrl string, qTime int64, qApi string) (msg string, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		msg = fmt.Sprintf("NewRequest Failed, URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
		return
	}
	q := req.URL.Query()
	q.Add("time", strconv.FormatInt(qTime, 10))
	q.Add("query", qApi)
	req.URL.RawQuery = q.Encode()
	ctx, cancel := context.WithTimeout(context.TODO(), _commonTimeout)
	defer cancel()
	req = req.WithContext(ctx)
	var response *http.Response
	response, err = prometheusClient.Do(req)
	bodyByte := make([]byte, 1024)
	total := 0
	if err == nil {
		total, _ = response.Body.Read(bodyByte)
		fmt.Printf("%s\n", string(bodyByte))
	} else {
		fmt.Printf("%+v\n", err)
	}
	resBody := &model.ApiDataModel{}
	err = json.Unmarshal(bodyByte[:total], resBody)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", resBody)
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

func fetchData(apiStr string, startTime, lastedTime int64) (valueRet int64) {
	valueRet = 0
	for i := int64(0); i < lastedTime; i += 5 {
		queryTime := startTime + i
		queryTimeStr := time.Unix(queryTime, 0)
		fmt.Printf("%+v\n", queryTimeStr.String())
		_, _ = doRequest(_prometheusUrl, queryTime, apiStr)
		res := make(map[string]struct{})
		// if "result" in res && len(res["result"]) > 0 && "value" in res["result"][0] {
		if hasResult(res["result"]) {
			// v := res["result"][0]["value"]
			// sv := strconv.Itoa(v[1])
			sv := ""
			if sv == "NaN" {
				// print("0", file=pout)
			} else {
				// print(sv, file=pout)
			}

		} else {
			// print("0", file=pout)
		}

	}
	return
}

func FetchDataAll(id int) (serviceRes model.ApplicationMetric) {
	serviceRes = model.ApplicationMetric{
		ID: id,
	}
	startNow := time.Now().Unix()
	// fetchData(_throughputApi, startNow, _lastedTime)

	ResponseTime := fetchData(_responseTimeApi, startNow, _lastedTime)
	serviceRes.ResponseTime = ResponseTime

	PodNumber := fetchData(_supplyPodApi, startNow, _lastedTime)
	serviceRes.PodNumber = PodNumber

	// fetchData(_demandPodApi, startNow, _lastedTime)

	RequestSuccessTotal := fetchData(_requestSuccessTotalApi, startNow, _lastedTime)
	serviceRes.RequestSuccessTotal = RequestSuccessTotal

	RequestFailTotal := fetchData(_requestFailTotalApi, startNow, _lastedTime)
	serviceRes.RequestFailTotal = RequestFailTotal

	ServiceTimeUnavailable := fetchData(_serviceTimeFailTotalApi, startNow, _lastedTime)
	serviceRes.ServiceTimeUnavailable = ServiceTimeUnavailable

	ServiceTimeAvailable := fetchData(_serviceTimeAvailableTotalApi, startNow, _lastedTime)
	serviceRes.ServiceTimeAvailable = ServiceTimeAvailable

	return
}

func hasResult(v interface{}) bool {
	return true
}
