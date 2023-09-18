package main

import (
	"go.uber.org/zap"
	"sync"
)

type DBConnector struct {
	connectionUrl *string
	once          sync.Once // 특정 함수를 한 번만 수행해야 할 때 sync.Once가 사용되어짐
}

var logger *zap.Logger
var singletonDBConnector *DBConnector

func init() {
	logger, _ = zap.NewDevelopment()
	logger.Info("1. logger 객체는 전체 패키지 중 init 함수 내부에서 한 번만 생성되어지기 때문에, 이 코드는 싱글톤 패턴의 좋은 예시가 됩니다.")
}

func (d *DBConnector) initializeUrlOnlyOnce() *string {
	if d.connectionUrl == nil {
		d.once.Do(func() {
			logger.Info("db url이 초기화 되지 않았을 경우 한 번만 초기화 합니다.")
			mysqlUrl := "gbike.mysql.url"
			d.connectionUrl = &mysqlUrl
		})
	} else {
		logger.Info("db url이 이미 초기화되었습니다.")
	}
	return d.connectionUrl
}

func (d *DBConnector) ShowDBUrl() {
	logger.Info(*d.connectionUrl)
}

func main() {
	logger.Info("2. init 함수에서 생성된 logger 객체는 main 함수 어디에서든 사용 가능합니다.")
	dbConnector := DBConnector{}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			dbConnector.initializeUrlOnlyOnce()
			done <- struct{}{}
		}()
	}

	// 모든 고루틴이 완료될 때까지 대기
	for i := 0; i < 1000; i++ {
		<-done
	}

	dbConnector.ShowDBUrl()
}
