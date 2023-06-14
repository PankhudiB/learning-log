## Generics in golang

- Type parameter

```go
func min[T any](a T, b T) T {

}
```

- Type inference -> to accept type aliases -> underlying type

```go

type sometype float64

func min[T float64](a T, b T) T {

}
becomes
func min[T ~float64](a T, b T) T {

}

PUT A TILDE 
```

- Type set --> to accept multiple type for func

```go
type mintypes interface {
    float64 | int
}

func min[T mintypes](a T, b T) T {

}
```