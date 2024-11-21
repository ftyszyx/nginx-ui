package kernel

var async []func()

var syncs []func()

// Boot the kernel
func Boot() {
    defer recovery()

    for _, v := range async {
        go v()
    }

    for _, v := range syncs {
        v()
    }
}

// RegisterAsyncFunc Register async functions, this function should be called before kernel boot.
func RegisterAsyncFunc(f ...func()) {
    async = append(async, f...)
}

// RegisterSyncsFunc Register syncs functions, this function should be called before kernel boot.
func RegisterSyncsFunc(f ...func()) {
    syncs = append(syncs, f...)
}
