package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type AlgorithmMetric struct {
	ID              int   `gorm:"primary_key;column:id" json:"id"`
	Reliability     int64 `column:"reliability" json:"reliability"`
	Availability    int64 `column:"availability" json:"availability"`
	Stability       int64 `column:"stability" json:"stability"`
	Cost            int64 `column:"cost" json:"cost"`
	Elasticity      int64 `column:"elasticity" json:"elasticity"`
	Oscillation     int64 `column:"oscillation" json:"oscillation"`
	SLASatisfaction int64 `column:"sla_satisfaction" json:"sla_satisfaction"`
	// 由于取消掉workload表，要增加一个WorkloadName字段记录负载名称
	// WorkloadName    string `column:"workload_name" json:"workload_name"`
}

func (AlgorithmMetric) TableName() string {
	return "algorithm_metrics"
}

type AlgorithmService struct {
	ID                 int    `gorm:"primary_key;column:id" json:"id"`
	AlgorithmName      string `column:"algorithm_name" json:"algorithm_name"`
	AlgorithmAddress   string `column:"algorithm_address" json:"algorithm_address"`
	AlgorithmNamespace string `column:"algorithm_namespace" json:"algorithm_namespace"`
	IsRunning          int    `column:"is_running" json:"is_running"`
}

func (AlgorithmService) TableName() string {
	return "algorithm_service"
}

type ApplicationMetric struct {
	ID                       int   `gorm:"primary_key;column:id" json:"id"`
	ResponseTime             int64 `column:"response_time" json:"response_time"`
	PodNumber                int64 `column:"pod_number" json:"pod_number"`
	RequestSuccessTotal      int64 `column:"request_success_total" json:"request_success_total"`
	RequestFailTotal         int64 `column:"request_fail_total" json:"request_fail_total"`
	ServiceTimeAvailable     int64 `column:"service_time_available" json:"service_time_available"`
	ServiceTimeUnavailable   int64 `column:"service_time_unavailable" json:"service_time_unavailable"`
	// 要增加一个StartTime字段记录数据抓取的开始时间
    // StartTime                  int64 `column:"start_time" json:"start_time"`

	// 以及每个指标的历史存储字段（string类型）
	// ResponseTimeRec            string `column:"response_time_rec" json:"response_time_rec"`
	// PodNumberRec               string `column:"pod_number_rec" json:"pod_number_rec"`
	// RequestSuccessTotalRec     string `column:"request_success_total_rec" json:"request_success_total_rec"`
	// RequestFailTotalRec        string `column:"request_fail_total_rec" json:"request_fail_total_rec"`
	// ServiceTimeAvailableRec    string `column:"service_time_available_rec" json:"service_time_available_rec"`
	// ServiceTimeUnavailableRec  string `column:"service_time_unavailable_rec" json:"service_time_unavailable_rec"`
}

func (ApplicationMetric) TableName() string {
	return "application_metrics"
}

type ApplicationService struct {
	ID               int    `gorm:"primary_key;column:id" json:"id"`
	ServiceName      string `gorm:"column:service_name" json:"service_name"`
	ServiceAddress   string `gorm:"column:service_address" json:"service_address"`
	ServiceNamespace string `gorm:"column:service_namespace" json:"service_namespace"`
	IsRunning        int    `gorm:"column:is_running; default:0" json:"is_running"`
	ServiceTime      int64  `gorm:"column:service_time; default:0" json:"service_time"`
}

func (ApplicationService) TableName() string {
	return "application_service"
}

type User struct {
	ID       int    `gorm:"primary_key;column:id"`
	Username string `column:"username"`
	Password string `column:"password"`
}

func (User) TableName() string {
	return "user"
}

type Workload struct {
	ID                int    `gorm:"primary_key;column:id"`
	WorkloadName      string `column:"workload_name"`
	WorkloadDuration  int64  `column:"workload_duration"`
	WorkloadType      string `column:"workload_type"`
	WorkloadStartTime int64  `column:"workload_start_time"`
	WorkloadEndTime   int64  `column:"workload_end_time"`
	WorkloadRoute     string `column:"workload_route"`
}

func (Workload) TableName() string {
	return "workload"
}
