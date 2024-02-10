package services

import (
	"sync"
)

var syncDataServiceOnce sync.Once
var syncDataService SyncDataService

type SyncDataService interface {
	Sync()
}

type SyncDataServiceImpl struct {
}

func NewSyncDataService() SyncDataService {
	syncDataServiceOnce.Do(func() {
		syncDataService = &SyncDataServiceImpl{}
	})
	return syncDataService
}

func (service *SyncDataServiceImpl) Sync() {
}
