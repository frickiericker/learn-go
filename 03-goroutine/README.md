Goroutines and channels.

- [Use Ticker][ticker] for rate-limiting requests
- Use channel for propagating errors from groutines to the main thread
- [Use WaitGroup][waitgroup] to wait for goroutines to finish
- Anonymous function captures variables by reference
- [Type assertion][typeassert]
- `return "", err` for `(string, error)` ([Example][stringerror])
- go fmt

[ticker]: http://qiita.com/jpshadowapps/items/a49c448d5e6b5f45f754
[waitgroup]: http://stackoverflow.com/a/18207832/5266681
[typeassert]: https://tour.golang.org/methods/15
[stringerror]: https://github.com/aws/aws-sdk-go/blob/db3e1e27b1ace4fc57be9c5cf7cea0566bd12034/aws/ec2metadata/api.go#L118
