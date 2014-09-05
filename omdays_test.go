package omdays

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"appengine/aetest"
)

func TestTask(t *testing.T) {
	assert := assert.New(t)
	t1 := &Task{Summary: "some summary", Done: "Todo"}
	assert.Equal(t1.Summary, "some summary")
	assert.Equal(t1.Done, "Todo")
}

func TestDefaultTaskList(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	k1 := defaultTaskList(c)
	k2 := defaultTaskList(c)
	if !k1.Equal(k2) {
		t.Error("Expected keys not to be equal, got: k1 == k2")
	}
}
