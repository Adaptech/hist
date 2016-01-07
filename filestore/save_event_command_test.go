package filestore

import (
	"testing"

	"github.com/nu7hatch/gouuid"
	"github.com/robertreppel/hist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStoringEventWithoutData(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-test-data"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("and an event without any data", func() {
			aggregateType := "GravelAggregate"
			id, _ := uuid.NewV4()
			aggregateID := id.String()
			var data []byte
			Convey("when save event is called", func() {
				err = eventStore.Save(aggregateType, aggregateID, "EventType", data)
				Convey("then an error occurs.", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestStoringNewEvent(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-test-data"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("and an event which doesn't exist yet", func() {
			aggregateType := "TestSaveEventAggregate"
			id, _ := uuid.NewV4()
			aggregateID := id.String()
			data := []byte("Here's a test event.")
			Convey("when save event is called", func() {
				eventStore.Save(aggregateType, aggregateID, "EventType", data)
				Convey("then an aggregate  with the event exists.", func() {
					events, err := eventStore.Get(aggregateType, aggregateID)
					if err != nil {
						panic(err)
					}
					So(len(events) > 0, ShouldBeTrue)
				})
			})
		})
	})
}
