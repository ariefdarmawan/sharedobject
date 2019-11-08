package sharedobject_test

import (
	"testing"

	"github.com/ariefdarmawan/sharedobject"

	"github.com/eaciit/toolkit"

	"github.com/smartystreets/goconvey/convey"
)

type obj struct {
	ID    string
	Value int
}

func newObj() *obj {
	return &obj{
		ID:    toolkit.RandomString(20),
		Value: toolkit.RandInt(1000),
	}
}

var (
	count = 1000
)

func TestShared(t *testing.T) {
	convey.Convey("prepare sharing obj", t, func() {
		sources := make([]obj, count)

		for i := 0; i < count; i++ {
			o := newObj()
			sources[i] = *o
		}
		so := sharedobject.NewSharedData()
		for _, v := range sources {
			so.Set(v.ID, v)
		}

		convey.So(so.Count(), convey.ShouldEqual, count)

		convey.Convey("get data", func() {
			for i := 0; i < 10; i++ {
				idx := toolkit.RandInt(count) - 1
				owant := sources[idx]
				oget := so.Get(owant.ID, obj{}).(obj)

				convey.So(owant.ID, convey.ShouldEqual, oget.ID)
				convey.So(owant.Value, convey.ShouldEqual, oget.Value)
			}

			convey.Convey("remove data", func() {
				idx := toolkit.RandInt(count) - 1
				odel := sources[idx]

				so.Remove(odel.ID)
				oget := so.Get(odel.ID, obj{}).(obj)
				convey.So(oget.ID, convey.ShouldBeBlank)
				convey.So(so.Count(), convey.ShouldEqual, count-1)
			})
		})
	})
}
