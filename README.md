# Go Throttle - a simple sliding window rate limiter

This module implements a Sliding Window throttle.

# INSTALL

    go get github.com/jbmcgill/go-throttle

# USAGE

    import github.com/jbmcgill/go-throttle
    // create throttle to allow max 100 invocations per 100ms
    throttle := &throttle.Throttle{PeriodicityMs: 100, Limit: 100}
    for i := 0; i < 200; i++ {
        throttle.Invoke(func() { _ = 1+1 })
    }


