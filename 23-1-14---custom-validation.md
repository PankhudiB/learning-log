
### Go Custom Struct Validation - the idiomatic way  

In the previous blog, we had gone over struct validations in golang using [validator library](.................)
Continuing on, in this blog, we'll cover custom validations.  

Let's say for the below struct:
```go
type Config struct {
    ServerUrl string
    AppPort   int
}
```

We could validate our struct with `validations` provided by the library like: 
```go
import (
    "fmt"
    "github.com/go-playground/validator"
    "strings"
)

func Validate(config Config) bool {
    validate := validator.New()
    validationErr := validate.Struct(config)
    if validationErr != nil {
        return false	
    }
    return true
}
```

Now onto the topic of _custom validations_. 

Say, we want to check whether `ServerUrl` field is an HTTPS URL
The straight forward way would be to introduce a function:  

```go
func customValidator(config Config) bool {
	value := config.ServerUrl
	if !strings.HasPrefix(value, "https://") {
		return false
	}
	return true
}
```
and remember to invoke it in our `Validate()` func [PUT EMOJI OF REMEMBRANCE]

Or we could do it the idiomatic way and leverage [validator library](https://github.com/go-playground/validator)

The custom validation function would then look like: 
```go
import (
    "github.com/go-playground/validator"
    "strings"
)

func CustomValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if !strings.HasPrefix(value, "https://") {
		return false
	}
	return true
}
```
"Why this weird function signature?" you ask - hold on to that question for a short while... 

The next step is to register our shiny new rule with the [validator library](https://github.com/go-playground/validator).
```go
import (
    "fmt"
    "github.com/go-playground/validator"
    "strings"
)

func Validate(config Config) bool {
    validate := validator.New()

    // Registering the new rule "is_https" ...
    err := validate.RegisterValidation("is_https", CustomValidation)
    if err != nil {
        fmt.Println("Error registering custom validation :", err.Error())
    }

    validationErr := validate.Struct(config)
    if validationErr != nil {
        return false	
    }
    return true
}
```

Now that our rule is registered let us attach the `tag` our struct's field:
```go
type Config struct {
	// Careful, the tag name should be identical to the one registered
    ServerUrl string `validate:"is_https"`
    AppPort   int 
}
```

And Voilà! we are done! 

Now, back to the question of this weird signature of `CustomValidation()` function -

Observe `validate.RegisterValidation()` takes in: 
1. A `tag` of how you want to refer the validation in your struct.
2. A `validator.Func` - essentially a function of type:
`func(fl FieldLevel) bool` [(reference source code)](https://github.com/go-playground/validator/blob/master/baked_in.go#L28) 


Even validations at struct level - ones that involve multiple fields of the same struct - are easily achievable with similar function signature: `func(fl StructLevel) bool`   

This is how the final code looks like: https://github.com/PankhudiB/go-config-validation/blob/main/configuration/custom_validation.go



