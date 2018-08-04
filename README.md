# Goland Ginkgo Reporter

Ginkgo does not output results in a format which can be converted into JSON 
(via [`go tool test2json`](https://golang.org/cmd/test2json/), which is used by
Goland to get a list of output.

This reporter will trick test2json into outputting the Ginkgo specs similar to
`go test` output.

## Usage

In your suite replace `RunSpecs(t, "Integration Suite")` with the following

```go

golandReporter := golandreporter.NewReporter()
RunSpecsWithCustomReporters(t, "Integration Suite", []Reporter{golangReporter})

```
