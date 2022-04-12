package retry

import (
	"github.com/glebteterin/go-xero"
	"time"
)

func Do(f func() error, maxSingleWait, maxTotalWait time.Duration, onRetry func(err *xero.ErrorResponse)) error {

	totalWait := time.Duration(0)

	for {
		err := f()
		if err != nil {
			if resp, ok := err.(*xero.ErrorResponse); ok {
				if resp.Response == nil || resp.Response.StatusCode != 429 {
					return err
				}
				if resp.Limits.RetryAfter() > maxSingleWait {
					return err
				}
				if resp.Limits.RetryAfter() + totalWait > maxTotalWait {
					return err
				}

				if onRetry != nil {
					onRetry(resp)
				}

				time.Sleep(resp.Limits.RetryAfter())
			} else {
				return err
			}
		} else {
			return nil
		}
	}
}
