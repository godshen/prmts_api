package controller

import (
	"control/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	user, _ := model.Test("student")
	log.Println(user)
}

func ServiceLogUp(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	r.ParseForm()
	service := model.ApplicationService{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return false, "Can not read body"
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println(err)
		return false, "Invalid json message"
	}
	//log.Println(service)
	id, err := model.CreateApplicationService(service)
	if err != nil {
		log.Println(err)
		return false, "Create service fail"
	}
	return true, id
}

func GetServiceRunStatus(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	// 根据ApplicationService中的ServiceAddress、ServiceNamespace、ServiceName信息
	// 调用k8s的实例扩缩api【获取】应用服务状态
	//（0表示未就绪、1表示运行中）
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	service, err := model.GetApplicationServiceByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(service)
	return true, service.IsRunning
}

func PatchServiceRunStatus(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	// 根据ApplicationService中的ServiceAddress、ServiceNamespace、ServiceName信息
	// 调用k8s的实例扩缩api【更新】应用服务状态
	//（0表示未就绪、1表示运行中）
	vars := mux.Vars(r)
	serviceID, status := vars["id"], vars["is_running"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		return false, err
	}
	err = model.UpdateApplicationServiceStatus(serviceIDInt, statusInt)
	if err != nil {
		return false, err
	}
	log.Println("Status update to:", status)
	return true, nil
}

func GetServiceRunTime(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	// 用个计时器记录服务运行时间
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	service, err := model.GetApplicationServiceByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(service)
	return true, service.ServiceTime
}

func GetServiceResponseTime(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.ResponseTime
}

func GetServiceNumbers(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.PodNumber
}

func GetRequestSuccessTotal(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.RequestSuccessTotal
}

func GetRequestFailTotal(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.RequestFailTotal
}

func GetServiceAvailableTime(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.ServiceTimeAvailable
}

func GetServiceUnavailableTime(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	appID := vars["id"]
	appIDInt, err := strconv.Atoi(appID)
	if err != nil {
		return false, err
	}
	appMetric, err := model.GetApplicationMetricByID(appIDInt)
	if err != nil {
		return false, err
	}
	log.Println(appMetric)
	return true, appMetric.ServiceTimeUnavailable
}

func AlgorithmLogUp(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	r.ParseForm()
	service := model.AlgorithmService{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return false, "Can not read body"
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println(err)
		return false, "Invalid json message"
	}
	//log.Println(service)
	id, err := model.CreateAlgorithmService(service)
	if err != nil {
		log.Println(err)
		return false, "Create service fail"
	}
	return true, id
}

func GetAlgorithmRunStatus(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	// 根据AlgorithmService中的AlgorithmAddress、AlgorithmNamespace、AlgorithmName信息
	// 调用k8s的实例扩缩api【获取】应用服务状态
	//（0表示未就绪、1表示运行中）
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	service, err := model.GetAlgorithmServiceByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(service)
	return true, service.IsRunning
}

func PatchAlgorithmRunStatus(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	// 根据AlgorithmService中的AlgorithmAddress、AlgorithmNamespace、AlgorithmName信息
	// 调用k8s的实例扩缩api【更新】应用服务状态
	//（0表示未就绪、1表示运行中）
	vars := mux.Vars(r)
	serviceID, status := vars["id"], vars["is_running"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		return false, err
	}
	err = model.UpdateAlgorithmRunStatus(serviceIDInt, statusInt)
	if err != nil {
		return false, err
	}
	log.Println("Status update to:", status)
	return true, nil
}

func GetReliabilityIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Reliability
}

func GetAvailabilityIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Availability
}

func GetStabilityIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Stability
}

func GetCostIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Cost
}

func GetElasticityIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Elasticity
}

func GetOscillationIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.Oscillation
}

func GetSlaSatisfactionIndex(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID := vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	metric, err := model.GetAlgorithmMetricByID(serviceIDInt)
	if err != nil {
		return false, err
	}
	log.Println(metric)
	return true, metric.SLASatisfaction
}

func UpdateServiceMetrics(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID:= vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	
	r.ParseForm()
	service := model.ApplicationMetric{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return false, "Can not read body"
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println(err)
		return false, "Invalid json message"
	}
	log.Println(service)

	err = model.UpdateServiceMetrics(serviceIDInt, service)
	if err != nil {
		log.Println(err)
		return false, "Update service fail"
	}
	return true, nil
}

func UpdateAlgorithmMetrics(w http.ResponseWriter, r *http.Request) (bool, interface{}) {
	vars := mux.Vars(r)
	serviceID:= vars["id"]
	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		return false, err
	}
	
	r.ParseForm()
	service := model.AlgorithmMetric{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return false, "Can not read body"
	}
	if err := json.Unmarshal(body, &service); err != nil {
		log.Println(err)
		return false, "Invalid json message"
	}
	log.Println(service)

	err = model.UpdateAlgorithmMetrics(serviceIDInt, service)
	if err != nil {
		log.Println(err)
		return false, "Update service fail"
	}
	return true, nil
}
// service类函数待实现

// PostServiceResponseTime函数
// 调用prometheus查询api

// PostServiceNumbers函数
// 调用prometheus查询api

// PostRequestSuccessTotal函数
// 调用prometheus查询api

// PostRequestFailTotal函数
// 调用prometheus查询api

// PostServiceAvailableTime函数
// 调用prometheus查询api

// PostServiceUnavailableTime函数
// 调用prometheus查询api

// algorithm类函数待实现

// PostReliabilityIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetRequestFailTotal函数,GetRequestSuccessTotal函数获取的时序数据
// 公式略（已有）
// 输出Reliability计算结果

// PostAvailabilityIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceAvailableTime函数,GetServiceUnavailableTime函数获取的时序数据
// 公式略（已有）
// 输出Availability计算结果

// PostStabilityIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceResponseTime函数获取的时序数据
// 公式略（已有）
// 输出Availability计算结果

// PostCostIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceNumbers函数获取的时序数据
// 公式略（已有）
// 输出Cost计算结果

// PostElasticityIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceNumbers函数获取的时序数据
// 公式略（已有）
// 输出Elasticity计算结果

// PostOscillationIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceNumbers函数获取的时序数据
// 公式略（已有）
// 输出Oscillation计算结果

// PostSlaSatisfactionIndex函数
// 把获取的时序数据取出来用于计算
// 输入GetServiceResponseTime函数获取的时序数据
// 公式略（已有）
// 输出SlaSatisfaction计算结果
