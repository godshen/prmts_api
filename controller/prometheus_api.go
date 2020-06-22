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
	_lastedTime                   = 60
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

func doRequest(apiUrl string, qTime int64, qApi string) (resBody *model.ApiDataModel, err error) {
	resBody = &model.ApiDataModel{}
	var req *http.Request
	req, err = http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		msg := fmt.Sprintf("NewRequest Failed, URL:%s, err:%v", apiUrl, err)
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
	} else {
		fmt.Printf("%+v\n", err)
	}
	err = json.Unmarshal(bodyByte[:total], resBody)
	if err != nil {
		msg := fmt.Sprintf("Http Get Failed: URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
		return
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		msg := fmt.Sprintf("Http Get Failed: URL:%s, err:%v", apiUrl, err)
		log.Print(msg)
	}
	return
}

func getResValue(r *model.ApiDataModel) int64 {
	if r == nil {
		return 0
	}
	if r.Status != "success" {
		return 0
	}
	if len(r.Data.Result) == 0 {
		return 0
	}
	dataTmp := r.Data.Result[0].Value[1]
	dataStr := dataTmp.(string)
	dataRet, _ := strconv.ParseInt(dataStr, 10, 64)
	return dataRet
}

func fetchData(apiStr string, startTime, lastedTime int64) (valueRet int64) {
	valueRet = 0
	validCount := int64(0)
	for i := int64(0); i < lastedTime; i += 5 {
		queryTime := startTime + i
		dataRes, err := doRequest(_prometheusUrl, queryTime, apiStr)
		if err != nil {
			continue
		}
		if v := getResValue(dataRes); v != 0 {
			valueRet += v
			validCount++
			log.Printf("apiStr: get %d", v)
		} else {
			log.Printf("apiStr get no data")
		}

	}
	if validCount != 0 {
		valueRet = valueRet / validCount
	}
	return
}

func FetchDataAll(id int) (serviceRes model.ApplicationMetric) {
	serviceRes = model.ApplicationMetric{
		ID: id,
	}
	startNow := time.Now().Unix()
	// fetchData(_throughputApi, startNow, _lastedTime)

	retResponseTime := fetchData(_responseTimeApi, startNow, _lastedTime)
	serviceRes.ResponseTime = retResponseTime

	retPodNumber := fetchData(_supplyPodApi, startNow, _lastedTime)
	serviceRes.PodNumber = retPodNumber

	// fetchData(_demandPodApi, startNow, _lastedTime)

	retRequestSuccessTotal := fetchData(_requestSuccessTotalApi, startNow, _lastedTime)
	serviceRes.RequestSuccessTotal = retRequestSuccessTotal

	retRequestFailTotal := fetchData(_requestFailTotalApi, startNow, _lastedTime)
	serviceRes.RequestFailTotal = retRequestFailTotal

	retServiceTimeUnavailable := fetchData(_serviceTimeFailTotalApi, startNow, _lastedTime)
	serviceRes.ServiceTimeUnavailable = retServiceTimeUnavailable

	retServiceTimeAvailable := fetchData(_serviceTimeAvailableTotalApi, startNow, _lastedTime)
	serviceRes.ServiceTimeAvailable = retServiceTimeAvailable

	return
}
