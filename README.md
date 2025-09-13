# conversions-go

Conversions is a module to convert between types in Go, with support for
`reflect.Type` or a custom type indicator, custom conversion functions and
indirect conversions.

## Features

- Uses generics in `Conversions` for type safety
- Supports `reflect.Type` or a custom type indicator
- Built-in conversions for `reflect.Type`
- Error handling, conversions can safely error if values cannot be converted
- Custom conversion functions can be registered
- Indirect conversions, if a direct conversion is not found between two types
  a chain of conversions will be searched for
- Thread-safe
  
## Installation

```bash
go get github.com/levelfourab/conversions-go
```

## Using

If you want to use conversions with your own types and information you need to
create a `Conversions` using the `WithTypeExtractor` option. The type extractor
is in charge of taking a value and returning the type indicator for that value.

This type is then used during conversion. If a direct conversion is not found
between the two types, a chain of conversions will be searched for.

Example:

```go
type MediaType string

type StringWithMediaType struct {
  MediaType MediaType
  Value     string
}

conversions, err := conversions.New(
  conversions.WithTypeExtractor(func(v StringWithMediaType) (MediaType, error) {
    return v.MediaType, nil
  }),
)

conversions.AddConversion("text/html", "text/plain", func(v StringWithMediaType) (StringWithMediaType, error) {
  // Using a hypothetical html2text package
  return html2text.Convert(v.Value)
})

// Perform the conversion
asPlainText, err := conversions.Convert(StringWithMediaType{
  MediaType: "text/html",
  Value:     "<h1>Hello World</h1>",
}, "text/plain")
```

### With Go-types

The `types` sub-package is available for the most common use case, converting
between Go types. Part of this includes default conversions for many built-in
types such as `string`, `int` (and variants), `uint` (and variants), `float32`
and `float64`.

Example:

```go
conversions, err := types.New(
  types.WithDefaultConversions(),
)

// Convert a string to an int, using types.TypeInt for convenience
asInt, err := conversions.Convert("1", types.TypeInt)

// Or using reflect.Type directly
asInt, err := conversions.Convert("1", reflect.TypeOf(int(0)))
```

Aliases are available to not have to carry generic types around:

- `types.Conversions` is an alias for `conversions.Conversions[reflect.Type, any]`
- `types.Converter` is an alias for `conversions.Converter[any]`
- `types.Option` is an alias for `conversions.Option[reflect.Type, any]`

## License

conversions-go is licensed under the MIT License. See [LICENSE](LICENSE) for 
the full license text.
