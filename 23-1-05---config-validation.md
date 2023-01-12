
### Go Struct Validation - the idiomatic way  

Validations, in general, are one of the most common problems across all programming languages.  
Before I start explaining how to do it idiomatically in golang... 

let me briefly explain why I dived into it in the first place.

While working on a project with tonnes of microservices, 
there was a need to maintain build time configurations for each of them.

This meant validating those configurations as well. Otherwise the microservices would have been at the risk of runtime failures!

So let's say we have the below configuration in `config.json`
```json
{
  "server_url": "https://some-server-url",
  "app_port": 8080
}
```

This can then be encapsulated in a golang struct: 

```go
type Config struct {
    ServerUrl string `json:"server_url"`
    AppPort   int    `json:"app_port"`
}
```

We would load the config.json to the struct:

```go
func LoadConfig() (*Config, error) {
	configFile, err := os.Open("configuration/config.json")
	if err != nil {
		log.Fatal("Could not open config file : ", err.Error())
	}

	decoder := json.NewDecoder(configFile)
	config := Config{}

	decodeErr := decoder.Decode(&config)
	if decodeErr != nil {
		log.Fatal("Could not decode config file : ", decodeErr.Error())
	}
	
	if !Validate(config) {
		return nil, errors.New("Invalid config !")
	}
	return &config, nil
}
```

Now, let's say we want to validate:
1. Server URL and port are non-empty 
2. The value of port is numeric and is between 8080 and 8085

Now the straight-forward way would be :

Write `if-else` conditional code to validate each of them. Something like:
```go
func Validate(config Config) bool {
	if config.ServerUrl == "" {
		return false
	}
	if config.AppPort == 0 {
		return false
	}
	if config.AppPort >= 8080 || config.AppPort <= 8085 {
		return false
	}
	return true
}
```    
This can get pretty messy with more such fields and validations. Yikes! 

There's got to be a better way (Otherwise this blog wouldn't exist ! :D ) 

A cleaner, idiomatic way to do this is to use **struct validations** <3

```go
type Config struct {
    ServerUrl string `json:"server_url" validate:"required"`
    AppPort   int    `json:"app_port" validate:"required,numeric,gte=8080,lte=8085"`
}    
```
Declarative and closer to the struct definition. Isn't that beautiful!

Now our validate function would look something like this:

```go
func Validate(config Config) bool {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		fmt.Println("Invalid config !", err.Error())
		return false
	}
	return true
}
```

Furthermore, we can have more meaningful errors to point out failing validations:
```go
func Validate(config Config) bool {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		fmt.Println("Invalid config !")
		for _, validationErr := range err.(validator.ValidationErrors) {
			fmt.Println(validationErr.StructNamespace() + " violated " + validationErr.Tag() + " validation.")
		}
		return false
	}
	return true
}    
```

Head over to the awesome [validator library (https://github.com/go-playground/validator#baked-in-validations)] for more such tags... 


//ADD once the second blog is ready

"But Pankhudi, what if I need to add validations of my own?" - you ask.

To this I answer head over to... [Custom validation blog link]