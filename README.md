# goreasons

A library to provide sane error handling for Go.

## Motivation

Why use this library for error handling? This library provides the mechanism to
report fatal and non-fatal errors, in a way that allows calculations to
optimistically continue.

For example, imagine the code below:

```golang
type Job func(ctx context.Context) error

func Run(ctx context.Context) error {
    var jobs []Job

    // ...

    for _, job := range jobs {
        if err := job(ctx); err != nil {
            return err
        }
    }

    return nil
}
```

While this is the idiomatic way to report errors, it limits error reporting to
a binary success / failure.

Imagine the case where we want to run these jobs, but we expect a certain level
of failure. We want to then ignore individual job errors.

However, if we simply ignore errors then we will fail to report when all jobs
fail. If all jobs are failing, it could be that the environment has been
incorrectly configured, and we should be alerting on that.

goreasons allows reporting "non-fatal" errors. These errors will not instantly
cause execution to fail, however if later execution cannot continue due to
faults in the configured system, the incremental non-fatal errors will be
included in the error trace of things to report. This prevents "non-fatal"
errors from being simply ignored.

## Usage

To simplify use of this error handling, `reasons.Reasons` can be used as an
`error` instance. For example:

```golang
import (
    "github.com/Crystalix007/goreasons/reasons"
    "github.com/Crystalix007/goreasons/reason"
)

type Job func(ctx context.Context) error

func Run(ctx context.Context) error {
    var jobs []Job
    var reasons reasons.Reasons

    // ...

    for _, job := range jobs {
        if err := job(ctx); err != nil {
            reasons.NonFatal("job failed", err, reason.Priority(reason.LOW))
        } else {
            return nil
        }
    }

    return reasons
}
```

If the error is fatal, it can be reported as such:

```golang
func DoSomething() error {
    var reasons reasons.Reasons
    // ...
    reasons.NonFatal("non-critical dependency", err)
    // ...

    if err := somethingCritical(); err != nil {
        return reasons.Fatal("critical dependency", err)
    }

    return nil
}
```

This will also report the non-critical failure that led up to the critical
failure occurring.
