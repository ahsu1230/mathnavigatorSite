# Handling JSON in Golang

Review [JSON](https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/02_protocols.md#what-is-json)

# What is JSON Serializing

[Serializing](https://en.wikipedia.org/wiki/Serialization) is the process of translating data structures into a storable format. In this case, we are serializing Golang structs into JSON. This is important because JSON is the standardized Javascript format for data, so we need to serialize structs into JSON before sending them to the frontend.

Golang provides JSON serializing in the `encoding/json` package ([docs](https://golang.org/pkg/encoding/json/)). The function

`func Marshal(v interface{}) ([]byte, error)`

serializes a Golang interface into a slice of bytes, which can then be converted into JSON using `bytes.NewBuffer()`.

# What is JSON Deserializing

JSON deserializing is the opposite. When frontend sends data to the backend in the form of JSON, we need to deserialize the data into a struct that backend can then work with. Golang also provides JSON deserializing through the function

`func Unmarshal(data []byte, v interface{}) error`,

which takes in a slice of bytes and inserts them into the `v` parameter.

# Golang Tags

In Golang, field declarations in a struct may be followed by an optional [string literal tag](https://golang.org/ref/spec#Struct_types). When serializing structs into JSON, we can specify the name of the field in JSON. 

Here is an example:
```go
type Example struct {
	Id        uint      `json:"id"`
	Empty     string    `json:"empty,omitempty"`
	CreatedAt time.Time `json:"-"`
}
```

In this struct, when we serialize, the `Id` field will have a JSON field of `id`. The `omitempty` means that if `Empty` is empty, it will not show up in the JSON. `-` means that this field will be omitted from the JSON no matter what.

# Uses in Our Codebase

In our codebase, [gin](https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/backend/04_gin.md) handles the JSON serializing and deserialization for us. However, for testing, we need to use `Marshal` and `Unmarshal`.

[Here](https://github.com/ahsu1230/mathnavigatorSite/blob/master/orion/pkg/tests_integration/setup_test.go#L106) is an example of a helper function that serializes structs into JSON:

```go
func createJsonBody(v interface{}) io.Reader {
	marshal, _ := json.Marshal(v)
	return bytes.NewBuffer(marshal)
}
```