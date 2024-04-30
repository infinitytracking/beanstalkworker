package beanstalkworker

import (
	"github.com/beanstalkd/go-beanstalk"
	"time"
)

type MockJob struct {
	id          uint64
	err         error
	body        *[]byte
	tube        string
	prio        uint32
	releases    uint32
	reserves    uint32
	timeouts    uint32
	conn        *beanstalk.Conn
	delay       time.Duration
	age         time.Duration
	returnPrio  uint32
	returnDelay time.Duration
	log         *MockLogger
	willDelete  bool
	willTouch   bool
	willRelease bool
}

func NewMockJob(body string) *MockJob {
	jobBody := []byte(body)
	return &MockJob{
		body: &jobBody,
	}
}

func NewWillDeleteMockJob(body string) *MockJob {
	jobBody := []byte(body)
	return &MockJob{
		body:       &jobBody,
		willDelete: true,
	}
}

func NewWillTouchMockJob(body string) *MockJob {
	jobBody := []byte(body)
	return &MockJob{
		body:      &jobBody,
		willTouch: true,
	}
}

func NewWillReleaseMockJob(body string) *MockJob {
	jobBody := []byte(body)
	return &MockJob{
		body:        &jobBody,
		willRelease: true,
	}
}

func (job *MockJob) Delete() {
	if job.willDelete == false {
		job.log.Error("Could not delete job: " + "unexpected call to Delete")
	}
}

func (job *MockJob) Touch() {
	if job.willTouch == false {
		job.log.Error("Could not touch job: " + "unexpected call to Touch")
	}
}

func (job *MockJob) Release() {
	if job.willRelease == false {
		job.log.Error("Could not release job: " + "unexpected call to Release")
	}
}

func (job *MockJob) LogError(a ...interface{}) {
	job.log.Error(a)
}

func (job *MockJob) LogInfo(a ...interface{}) {
	job.log.Info(a)
}

func (job *MockJob) GetAge() time.Duration {
	// Mock implementation
	return job.age
}

func (job *MockJob) GetPriority() uint32 {
	// Mock implementation
	return job.prio
}

func (job *MockJob) GetReleases() uint32 {
	// Mock implementation
	return job.releases
}

func (job *MockJob) GetReserves() uint32 {
	// Mock implementation
	return job.reserves
}

func (job *MockJob) GetTimeouts() uint32 {
	// Mock implementation
	return job.timeouts
}

func (job *MockJob) GetDelay() time.Duration {
	// Mock implementation
	return job.delay
}

func (job *MockJob) GetTube() string {
	// Mock implementation
	return job.tube
}

func (job *MockJob) GetConn() *beanstalk.Conn {
	return job.conn
}

func (job *MockJob) SetReturnPriority(prio uint32) {
	// Mock implementation
	job.returnPrio = prio
}

func (job *MockJob) SetReturnDelay(delay time.Duration) {
	// Mock implementation
	job.returnDelay = delay
}
