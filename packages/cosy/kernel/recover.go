package kernel

import (
    "github.com/uozi-tech/cosy/logger"
    "runtime"
)

// recovery recover from panic
func recovery() {
    if err := recover(); err != nil {
        buf := make([]byte, 1024)
        runtime.Stack(buf, false)
        logger.Errorf("%s\n%s", err, buf)
    }
}
