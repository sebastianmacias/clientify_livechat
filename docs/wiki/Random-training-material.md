# Random training material

- [Random training material](#random-training-material)
    - [I know how to do this with Node.js](#i-know-how-to-do-this-with-nodejs)
    - [Learning resources](#learning-resources)
    - [stdlib](#stdlib)
    - [Installing global tools](#installing-global-tools)
    - [Uppercase enforced acronyms](#uppercase-enforced-acronyms)
    - [`init` functions](#init-functions)
    - [Go modules](#go-modules)
    - [pgFormatter](#pgformatter)
    - [SQLBoiler](#sqlboiler)
    - [Testing](#testing)
    - [When to use `context` to store things?](#when-to-use-context-to-store-things)
    - [Slices](#slices)
    - [CLI](#cli)
    - [gogenerate](#gogenerate)
    - [goswagger](#goswagger)
    - [Error handling in go](#error-handling-in-go)

### I know how to do this with Node.js
https://github.com/miguelmota/golang-for-nodejs-developers

### Learning resources
* https://github.com/inancgumus/learngo
* https://quii.gitbook.io/learn-go-with-tests/
* https://www.manning.com/books/go-in-practice

### stdlib
> Get an overview of the stdlib of go itself first!
https://golang.org/pkg/#stdlib

> Go formatting directives
https://golang.org/pkg/fmt/

> panic means an unexpected condition arises in your Go program due to which the execution of your program is terminated
> Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
https://gobyexample.com/panic *Do not use panic in normal scenarios, eventually only for on-off cli utils*

### Installing global tools
> We might want to install some go-based tools that are not being imported, but are used as part of project’s development environment. The officially recommended approach is to add a `tools.go` file (the name doesn’t matter). 

See [How can I track tool dependencies for a module?](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
See [Tools as dependencies](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md)

For automated install for go tools refer to [Manage Go Tools via Go Modules](https://marcofranssen.nl/manage-go-tools-via-go-modules/)

### Uppercase enforced acronyms

> Words in names that are initialisms or acronyms (e.g. "URL" or "NATO") have a consistent case. For example, "URL" should appear as "URL" or "url" (as in "urlPony", or "URLPony"), never as "Url". As an example: ServeHTTP not ServeHttp. For identifiers with multiple initialized "words", use for example "xmlHTTPRequest" or "XMLHTTPRequest".
https://github.com/golang/go/wiki/CodeReviewComments#initialisms
https://stackoverflow.com/questions/52358247/what-are-golang-struct-field-naming-conventions
https://github.com/golang/lint/blob/master/lint.go#L770

### `init` functions

> `init` is [automatically] called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized. 

https://golang.org/doc/effective_go.html#init
https://medium.com/golangspec/init-functions-in-go-eac191b3860a

### Go modules
[Go modules](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31)
> On a more general note, the purpose of go mod tidy is to also add any dependencies needed for other combinations of OS, architecture, and build tags. Make sure to run this before every release.

[Go modules by example](https://github.com/go-modules-by-example/index/blob/master/README.md)

### pgFormatter
We use pgFormatter both in VSCode and in our `Makefile` to enforce a proper formatting to all `*.sql` files
https://marketplace.visualstudio.com/items?itemName=bradymholt.pgformatter
https://github.com/darold/pgFormatter

### SQLBoiler
About UUID support: https://github.com/volatiletech/sqlboiler/issues/58

> No global function variants
We can't use the global db variant, as we don't want to lose the ability to do parallel testing, see [function variations](https://github.com/volatiletech/sqlboiler#function-variations)
The idea is to use the `db` that's available on the `server` struct, which means you always need to pass the `server` or `db` to your functions (same as it should happen with `context` anyways)

> panic variants might be handy for one-off tasks but in general are dangerours. Do not use them in a regular service.

> hooks Won't hopefully be required, refrain from using them unless you have a very very good reason.

### Testing

> Black-Box Testing (packages with `_test` suffix) is also required to deal with cyclic dependency imports. 
* https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language/31443271#31443271
* https://github.com/golang/go/issues/25223

> `TestMain` would be executed once per package (and is not required any more by us)
* https://www.reddit.com/r/golang/comments/7x4czh/namespacing_integration_tests_with_testmain/
* https://stackoverflow.com/questions/51026376/golang-global-setup-for-all-tests-within-same-and-other-sub-modules
* https://groups.google.com/forum/#!topic/golang-nuts/SxEkZhWl3QA

> Regarding `t.*` log methods
https://golang.org/pkg/testing/#B.Error
* Fail marks the function as having failed but continues execution.
* FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine).
* Error is equivalent to Log followed by Fail.
* Fatal is equivalent to Log followed by FailNow.

> Regarding `testdata` directory
https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859  `Firstly, go build ignores directory named testdata.`  `Secondly, when go test runs, it sets current directory as package directory.`
* https://dave.cheney.net/2016/05/10/test-fixtures-in-go `Second, the Go tool will ignore any directory in your $GOPATH that starts with a period, an underscore, or matches the word testdata`

> stdout and stderr are hijacked while testing, you will only receive output on test failure or when running with the `-v` flag (enabled in VSCode when running tests from there)
* https://blog.blueberry.io/end-to-end-tests-in-go-and-postgresql-40bc4884ac42
* https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318

> Random test related resources
* https://ieftimov.com/post/testing-in-go-fixtures/

> Create a server instance inside each test — if expensive things lazy load, this won’t take much time at all, even for big components
https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html


### When to use `context` to store things?
> **Context.Value should inform, not control**
https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39

> Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
> That said, people do use this approach. And if your project consists of a sprawling set of packages – and using a global config is out of the question – it's quite an attractive proposition.
* https://www.alexedwards.net/blog/organising-database-access -- `Dependency injection` approch currently used, we do not want to use the `Request-scoped context` for this!

> The go vet tool checks that CancelFuncs are used on all control-flow paths.
> Programs that use Contexts should follow these rules to keep interfaces consistent across packages and enable static analysis tools to check context propagation
> The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
https://golang.org/pkg/context/

### Slices
https://github.com/golang/go/wiki/SliceTricks
FIFO / Queues -- https://stackoverflow.com/questions/23531891/how-do-i-succinctly-remove-the-first-element-from-a-slice-in-go/23532104

### CLI
https://golang.org/pkg/flag/ `scripts, parsing CLI flags`

### gogenerate
Sample `stringer` -- https://medium.com/@arjun.mahishi/golang-stringer-ad01db65e306

### goswagger
why are required string pointer types? https://github.com/go-swagger/go-swagger/issues/1188
```
Add it to the required array
Set MinLength: 1
Set x-nullable: false
```

https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9

How to add response descriptions: https://github.com/go-swagger/go-swagger/issues/738#issuecomment-273014052

### Error handling in go

> In Go, error handling is important. The language's design and conventions encourage you to explicitly check for errors where they occur (as distinct from the convention in other languages of throwing exceptions and sometimes catching them).
https://blog.golang.org/error-handling-and-go

> Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks. https://gobyexample.com/errors