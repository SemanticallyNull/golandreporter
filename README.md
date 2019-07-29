# Goland Ginkgo Reporter

Ginkgo does not output results in a format which can be converted into JSON 
(via [`go tool test2json`](https://golang.org/cmd/test2json/), which is used by
Goland to get a list of output.

This reporter will trick test2json into outputting the Ginkgo specs similar to
`go test` output.

![Ginkgo output in Goland's "Run" window](https://gist.githubusercontent.com/SemanticallyNull/19215f014f3ef0db3c3cd0b46da4d929/raw/18801f7f19a26ea19ba48a1c95d5e068396f7dd7/image1.png)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSemanticallyNull%2Fgolandreporter.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FSemanticallyNull%2Fgolandreporter?ref=badge_shield)

## Usage

In your suite replace `RunSpecs(t, "Integration Suite")` with the following:

```go
golandReporter := golandreporter.NewGolandReporter()
RunSpecsWithCustomReporters(t, "Integration Suite", []Reporter{golangReporter})
```

If you want to retain normal Ginkgo formatting when using it from the CLI the
best option is to use an environment variable in your Run Configuration, and
use it like this:

```go
RunSpecsWithCustomReporters(t, "Integration Suite", []Reporter{golandreporter.NewAutoGolandReporter})
```



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSemanticallyNull%2Fgolandreporter.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSemanticallyNull%2Fgolandreporter?ref=badge_large)