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