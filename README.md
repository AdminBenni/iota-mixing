This linter is made to check if unrelated consts are being declared in the same block as an `iota`, which is usually used for defining enums in GO.

Example of `iota` mixing:
```go
package users

type Role uint64

const (
	UserFmtString = "user(%d)"
	
	Unknown Role = iota     // 1
	Client                  // 2
	Seller                  // 3
	Administrator           // 4
)
```

Defining enums with `iota` like this is very common. Though, looking at the example, we see that someone innocently thought they could reuse the same const block the `iota` was in to declare an unrelated constant which was also needed in the file. However, this is really dangerous, since `iota` is just a counter which counts from zero how many constants have already been defined within the block, so they have unintentionally altered the values of the enums. Notice how `Unknown` gets the value of `1` rather than `0`. This is especially dangerous if these enums have already been pushed to production.

This can be quite a nasty bug to figure out if you don't know `iota` behaves like this. The opinion of this linter is that mixing declarations in the same const block and altering the `iota` value like this is not intended in most circumstances. 
