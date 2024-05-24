# Optional

very simple optional for golang

## example

```go
import "github.com/damonchen/optional"

type Person struct {
	Email    Optional[string] `json:"email"`
	Username Optional[string] `json:"username"`
}
var person Person

data := []byte("{\"email\": \"xxxxx@gmail.com\"}")

fmt.Println(person.Username.IsNull())
fmt.Println(person.Email.IsNull(), person.Email.Value())
```