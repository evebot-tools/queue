package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewCronJob(t *testing.T) {
	Convey("TestNewCronJob", t, func() {
		Convey("Should return a new cron job", func() {
			job := NewCronJob("test", true, 30)
			So(job.Action, ShouldEqual, "test")
			So(job.Enabled, ShouldEqual, true)
			So(job.TTL, ShouldEqual, 30)
			So(job.StringData, ShouldBeNil)
			So(job.IntData, ShouldBeNil)
		})
	})
	Convey("TestNewCronJobWithStringData", t, func() {
		Convey("Should return a new cron job with string data", func() {
			job := NewCronJobWithStringData("test", true, 30, "data")
			So(job.Action, ShouldEqual, "test")
			So(job.Enabled, ShouldEqual, true)
			So(job.TTL, ShouldEqual, 30)
			So(*job.StringData, ShouldEqual, "data")
		})
	})
	Convey("TestNewCronJobWithIntData", t, func() {
		Convey("Should return a new cron job with int data", func() {
			job := NewCronJobWithIntData("test", true, 30, 111)
			So(job.Action, ShouldEqual, "test")
			So(job.Enabled, ShouldEqual, true)
			So(job.TTL, ShouldEqual, 30)
			So(*job.IntData, ShouldEqual, 111)
		})
	})
}
