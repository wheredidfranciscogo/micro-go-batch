package test

import (
    "errors"
    "testing"
    "micro-go-batch/pkg"
)

func TestNewJob(t *testing.T) {
    data := "data data"
    job := pkg.NewJob(data) // use the NewJob function from pkg package
    if job.Data != data {
        t.Errorf("NewJob() = %s; expected %s", job.Data, data)
    }
}

func TestNewJobResult(t *testing.T) {
    data := "data data"
    errMsg := "error error"
    err := errors.New(errMsg) // using the error package
    jobResult := pkg.NewJobResult(data, err)
    if jobResult.Data != data {
        t.Errorf("NewJobResult().Data = %s; expected %s", jobResult.Data, data)
    }
    if jobResult.Error.Error() != errMsg { // use Error() method to get the error message
        t.Errorf("NewJobResult().Error = %s; expected %s", jobResult.Error, errMsg)
    }
}
