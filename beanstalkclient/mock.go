package beanstalkclient

import (
	"errors"
	"time"

	"github.com/google/go-cmp/cmp"
)

type MockClient struct {
	Name        string
	expectedPut *expectedPut
}

type expectedPut struct {
	body string
	id   uint64
}

type MockConn struct {
}

func NewMock(name string) Client {
	return &MockClient{Name: name}
}

func NewMockWillPut(name string, body string, id uint64) Client {
	c := &MockClient{Name: name}
	c.ExpectPut(body, id)
	return c
}

func (c *MockClient) Put(body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error) {
	if c.expectedPut == nil {
		return 0, errors.New("unexpected call to Put")
	}

	if diff := cmp.Diff(c.expectedPut.body, string(body)); diff != "" {
		return 0, errors.New("unexpected body: " + diff)
	}

	return c.expectedPut.id, nil
}

func (c *MockClient) PeekReady() (id uint64, body []byte, err error) {
	// Mock implementation
	return 0, nil, errors.New("not implemented")
}

func (c *MockClient) PeekDelayed() (id uint64, body []byte, err error) {
	// Mock implementation
	return 0, nil, errors.New("not implemented")
}

func (c *MockClient) PeekBuried() (id uint64, body []byte, err error) {
	// Mock implementation
	return 0, nil, errors.New("not implemented")
}

func (c *MockClient) Kick(bound int) (n int, err error) {
	// Mock implementation
	return 0, errors.New("not implemented")
}

func (c *MockClient) Stats() (map[string]string, error) {
	// Mock implementation
	return nil, errors.New("not implemented")
}

func (c *MockClient) Pause(d time.Duration) error {
	// Mock implementation
	return errors.New("not implemented")
}

func (c *MockClient) ExpectPut(body string, id uint64) {
	c.expectedPut = &expectedPut{body: body, id: id}
}
