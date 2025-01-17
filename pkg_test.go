package log_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eluv-io/apexlog-go"
	"github.com/eluv-io/apexlog-go/handlers/memory"
)

type Pet struct {
	Name string
	Age  int
}

func (p *Pet) Fields() log.Fields {
	return log.Fields{
		{Name: "name", Value: p.Name},
		{Name: "age", Value: p.Age},
	}
}

func TestInfo(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	log.Infof("logged in %s", "Tobi")

	e := h.Entries[0]
	assert.Equal(t, e.Message, "logged in Tobi")
	assert.Equal(t, e.Level, log.InfoLevel)
}

func TestFielder(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	pet := &Pet{"Tobi", 3}
	log.WithFields(pet).Info("add pet")

	e := h.Entries[0]
	assert.Equal(t, log.Fields{{Name: "name", Value: "Tobi"}, {Name: "age", Value: 3}}, e.Fields)
}

// Unstructured logging is supported, but not recommended since it is hard to query.
func Example_unstructured() {
	log.Infof("%s logged in", "Tobi")
}

// Structured logging is supported with fields, and is recommended over the formatted message variants.
func Example_structured() {
	log.WithField("user", "Tobo").Info("logged in")
}

// Errors are passed to WithError(), populating the "error" field.
func Example_errors() {
	err := errors.New("boom")
	log.WithError(err).Error("upload failed")
}

// Multiple fields can be set, via chaining, or WithFields().
func Example_multipleFields() {
	log.WithFields(log.Fields{
		{Name: "user", Value: "Tobi"},
		{Name: "file", Value: "sloth.png"},
		{Name: "type", Value: "image/png"},
	}).Info("upload")
}

// Watch can be used to simplify logging of start and completion events,
// for example an upload which may fail.
func Example_watch() {
	fn := func() (err error) {
		defer log.Watch("upload").Stop(&err)
		return
	}

	fn()
	return
}
