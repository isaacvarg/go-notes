- bufio allows you to handle longer user inputs
  - reader := bufio.NewReader(os.Stdin)
  - Stdin is simply your command line
  - then read by
  - reader.ReadString('\n')
  - this tells to stop reading at line break
  - notice that this is using '' and not ""
  - this is a special value type in go called rune
- encoding/json package
  - json.Marshal() 
  - accepts structs
- struct tags is like adding metadata that any code or methods can access
  - this is useful in the above example where
  - Marshal only works on values that are exported in the struct
  - so if you have a property like Name, it would be in the json file as "Name": 
  - whereas the convention is lower case
  - you add struct tags by defining them in the struct after the type with backticks

```go

type Note struct {
  Title string `json:"title"`
}
```

- you can of course add any metadata you want (e.g., not just json but xxyyzz)
- but the code obviously must parse the meta data, and in this case the json.Marshal 
- respects the json metadata
