package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
)


func Test(username string) (User, error) {
	u := User{}
	//res := GormDB.Where("username = ? ", username).Find(&u)
	res := GormDB.First(&u, "username = ?", username)
	return u, res.Error
}

func GetApplicationServiceByID(id int) (ApplicationService, error) {
	service := ApplicationService{}
	res := GormDB.First(&service, id)
	return service, res.Error
}

func GetApplicationMetricByID(id int) (ApplicationMetric, error) {
	metric := ApplicationMetric{}
	res := GormDB.First(&metric, id)
	return metric, res.Error
}

func GetAlgorithmServiceByID(id int) (AlgorithmService, error) {
	service := AlgorithmService{}
	res := GormDB.First(&service, id)
	return service, res.Error
}

func GetAlgorithmMetricByID(id int) (AlgorithmMetric, error) {
	metric := AlgorithmMetric{}
	res := GormDB.First(&metric, id)
	return metric, res.Error
}

func CreateApplicationService(service ApplicationService) (int, error) {
	res := GormDB.Create(&service)
	return service.ID, res.Error
}

func CreateAlgorithmService(service AlgorithmService) (int, error) {
	res := GormDB.Create(&service)
	return service.ID, res.Error
}

func UpdateApplicationServiceStatus(id int, status int) error {
	service := ApplicationService{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("is_running", status).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAlgorithmRunStatus(id int, status int) error {
	service := AlgorithmService{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("is_running", status).Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceMetrics(id int, info ApplicationMetric) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Updates(info).Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceResponseTime(id int, responseTime int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("response_time", responseTime).Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceNumbers(id int, podNumber int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("pod_number", podNumber).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequestSuccessTotal(id int, requestSuccessTotal int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("request_success_total", requestSuccessTotal).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequestFailTotal(id int, requestFailTotal int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("request_fail_total", requestFailTotal).Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceAvailableTime(id int, serviceTimeAvailable int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("service_time_available", serviceTimeAvailable).Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceUnavailableTime(id int, serviceTimeUnavailable int64) error {
	if GormDB == nil {
		return errors.New("No Database")
	}
	service := ApplicationMetric{}
	if err := GormDB.First(&service, id).Error; err != nil {
		return err
	}
	if err := GormDB.Model(&service).Update("service_time_unavailable", serviceTimeUnavailable).Error; err != nil {
		return err
	}
	return nil
}
