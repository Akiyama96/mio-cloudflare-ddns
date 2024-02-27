package services

import (
	"log"
	"sync"
	"time"
)

func StartServices() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go startAutoOverwriteService(wg)

	wg.Wait()
}

func startAutoOverwriteService(wg *sync.WaitGroup) {
	defer wg.Done()

	autoOverwriteServerSetting, err := readAutoOverwriteServerConfig()
	if err != nil {
		log.Println(err)
		return
	}

	ticker := time.NewTicker(time.Duration(autoOverwriteServerSetting.ServerSetting.IntervalsMin) * time.Minute)
	defer ticker.Stop()
	for {
		err = autoOverwrite(autoOverwriteServerSetting)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Update dns success!")
		}

		select {
		case <-ticker.C:
			continue
		}
	}
}
