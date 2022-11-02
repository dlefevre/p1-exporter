package metrics

import (
	"errors"
	"fmt"
)

type MetricBuilderError interface {
	Error() string
	Ignore() bool
}

type MetricBuilderErrorImpl struct {
	err    error
	ignore bool
}

func NewMetricBuilderError(msg string, ignore bool) MetricBuilderError {
	return &MetricBuilderErrorImpl{
		err:    errors.New(msg),
		ignore: ignore,
	}
}

func (err *MetricBuilderErrorImpl) Error() string {
	return fmt.Sprintf("MetricBuilder: %s", err.err.Error())
}

func (err *MetricBuilderErrorImpl) Ignore() bool {
	return err.ignore
}
