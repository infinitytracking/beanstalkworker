package beanstalkclient

import (
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

type Client interface {
	Put(body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error)
	PeekReady() (id uint64, body []byte, err error)
	PeekDelayed() (id uint64, body []byte, err error)
	PeekBuried() (id uint64, body []byte, err error)
	Kick(bound int) (n int, err error)
	Stats() (map[string]string, error)
	Pause(d time.Duration) error
}

type beanstalkClient struct {
	tube *beanstalk.Tube
}

func New(c *beanstalk.Conn, name string) Client {
	return &beanstalkClient{
		tube: &beanstalk.Tube{Conn: c, Name: name},
	}
}

func (t *beanstalkClient) Put(body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error) {
	return t.tube.Put(body, pri, delay, ttr)
}

func (t *beanstalkClient) PeekReady() (id uint64, body []byte, err error) {
	return t.tube.PeekReady()
}

func (t *beanstalkClient) PeekDelayed() (id uint64, body []byte, err error) {
	return t.tube.PeekDelayed()
}

func (t *beanstalkClient) PeekBuried() (id uint64, body []byte, err error) {
	return t.tube.PeekBuried()
}

func (t *beanstalkClient) Kick(bound int) (n int, err error) {
	return t.tube.Kick(bound)
}

func (t *beanstalkClient) Stats() (map[string]string, error) {
	return t.tube.Stats()
}

func (t *beanstalkClient) Pause(d time.Duration) error {
	return t.tube.Pause(d)
}
