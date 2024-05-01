package beanstalkworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// MockLogger A custom logger that must implement Info() Infof(), Error() and Errorf() to
// implement CustomLogger
type MockLogger struct {
}

func (l *MockLogger) Info(v ...interface{}) {
	log.Print("Info: ", fmt.Sprint(v...))
}

func (l *MockLogger) Infof(format string, v ...interface{}) {
	format = "Infof: " + format
	log.Print(fmt.Sprintf(format, v...))
}

func (l *MockLogger) Error(v ...interface{}) {
	log.Fatal("Error: ", fmt.Sprint(v...))
}

func (l *MockLogger) Errorf(format string, v ...interface{}) {
	format = "Errorf: " + format
	log.Fatalf(fmt.Sprintf(format, v...))
}

type MockWorker struct {
	tubeSubs   map[string]func(*MockJob)
	tubeJobs   map[string]*MockJob
	numWorkers int
	log        *Logger
}

func NewMockWorker() WorkerClient {
	return &MockWorker{
		tubeSubs: make(map[string]func(*MockJob)),
	}
}

func NewMockWillDeleteWorker(tube, jobStr string) WorkerClient {
	return &MockWorker{
		tubeSubs: make(map[string]func(*MockJob)),
		tubeJobs: map[string]*MockJob{
			tube: NewWillDeleteMockJob(jobStr),
		},
		log: NewDefaultLogger(),
	}
}

func NewMockWillTouchWorker(tube, jobStr string) WorkerClient {
	return &MockWorker{
		tubeSubs: make(map[string]func(*MockJob)),
		tubeJobs: map[string]*MockJob{
			tube: NewWillTouchMockJob(jobStr),
		},
		log: NewDefaultLogger(),
	}
}

func NewMockWillReleaseWorker(tube, jobStr string) WorkerClient {
	return &MockWorker{
		tubeSubs: make(map[string]func(*MockJob)),
		tubeJobs: map[string]*MockJob{
			tube: NewWillReleaseMockJob(jobStr),
		},
		log: NewDefaultLogger(),
	}
}

func (w *MockWorker) SetNumWorkers(numWorkers int) {
	w.numWorkers = numWorkers
}

func (w *MockWorker) SetLogger(cl CustomLogger) {
	w.log.Info = cl.Info
	w.log.Error = cl.Error
	w.log.Errorf = cl.Errorf
	w.log.Infof = cl.Infof
}

func (w *MockWorker) Subscribe(tube string, cb Handler) {
	w.tubeSubs[tube] = func(job *MockJob) {
		jobVal := reflect.ValueOf(job)
		cbFunc := reflect.ValueOf(cb)
		cbType := reflect.TypeOf(cb)
		if cbType.Kind() != reflect.Func {
			panic("Handler needs to be a func")
		}

		dataType := cbType.In(1)
		dataPtr := reflect.New(dataType)

		if err := json.Unmarshal(*job.body, dataPtr.Interface()); err != nil {
			job.LogError("Error decoding JSON for job: ", err, ", '", string(*job.body))
			return
		}

		cbFunc.Call([]reflect.Value{jobVal, reflect.Indirect(dataPtr)})
	}
}

func (w *MockWorker) Run(ctx context.Context) {
	if w.numWorkers <= 0 {
		w.log.Error("No Worker threads defined, cannot proceed.")
		return

	}

	if len(w.tubeSubs) <= 0 {
		w.log.Error("No job subscriptions defined, cannot proceed.")
		return
	}

	for tube, job := range w.tubeJobs {
		w.log.Infof("Processing mock job from tube %s", tube)
		if cb, ok := w.tubeSubs[tube]; ok {
			cb(job)
			ctx.Done()
		}
	}
}

func (w *MockWorker) SetUnmarshalErrorAction(action string) {
	// Not implemented
}
