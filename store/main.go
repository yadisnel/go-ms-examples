package main

import (
	"time"

	"github.com/yadisnel/go-ms/v2/logger"
	"github.com/yadisnel/go-ms/v2/service"
	"github.com/yadisnel/go-ms/v2/service/mucp"
	"github.com/yadisnel/go-ms/v2/store"
	"github.com/yadisnel/go-ms/v2/store/memory"
)

func main() {
	service := mucp.NewService(
		service.Name("go.micro.service.store-test"),
		service.Version("latest"),
	)

	service.Init()

	st := memory.NewStore()

	// write to the store
	record := &store.Record{Key: "foo", Value: []byte("bar")}
	if err := st.Write(record); err != nil {
		logger.Fatalf("Error writing to the store: %v", err)
	}
	logger.Infof("Wrote value to the store")

	// read from the store
	go func() {
		ticker := time.NewTicker(time.Second * 15)

		for {
			<-ticker.C

			recs, err := st.Read("foo")
			if err != nil {
				logger.Errorf("Error reading from store: %v", err)
				continue
			}
			if len(recs) == 0 {
				logger.Errorf("No results returned from store")
				continue
			}
			logger.Infof("Got a value from the store: %v", string(recs[0].Value))
		}
	}()

	// run the service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
