package event

import (
	"fmt"

	"github.com/ovh/cds/engine/api/cache"
	"github.com/ovh/cds/engine/log"
	"github.com/ovh/cds/sdk"
)

var hostname, cdsname string
var kafkaBroker Broker
var brokers []Broker

// Broker event typed
type Broker interface {
	initialize(options interface{}) (Broker, error)
	sendEvent(event *sdk.Event) error
	status() string
	close()
}

func getBroker(t string, option interface{}) (Broker, error) {
	switch t {
	case "kafka":
		k := &KafkaClient{}
		return k.initialize(option)
	}
	return nil, fmt.Errorf("Invalid Broker Type %s", t)
}

// Initialize initializes event system
func Initialize(k KafkaConfig, h, cds string) error {
	hostname = h
	cdsname = cds

	brokers = []Broker{}
	if k.Enabled {
		var errk error
		kafkaBroker, errk = getBroker("kafka", k)
		if errk != nil {
			return errk
		}
		brokers = append(brokers, kafkaBroker)
	}
	return nil
}

// DequeueEvent runs in a goroutine and dequeue event from cache
func DequeueEvent() {
	for {
		e := sdk.Event{}
		cache.Dequeue("events", &e)
		for _, b := range brokers {
			if err := b.sendEvent(&e); err != nil {
				log.Warning("Error while sending message: %s", err)
			}
		}
	}
}

// Close closes event system
func Close() {
	for _, b := range brokers {
		b.close()
	}
}

// Status returns Event status
func Status() string {
	o := ""
	for _, b := range brokers {
		o += b.status() + " "
	}

	if o == "" {
		o = "⚠ "
	}

	return o
}
