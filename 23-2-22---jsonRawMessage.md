
### json.RawMessage

1. When to use ? 
   when you want to ignore OR defer unmarshalling of certain part of your json


## Message Producer
```go
    	type SubMessage struct {
            Name  string `json:"name"`
            Place int    `json:"place"`
        }

        type Message struct {
            Type       string     `json:"type"`
            SubMessage SubMessage `json:"sub_message"`
        }

        message := Message{
		Type:       "blah",
		SubMessage: SubMessage{Name: "Pankhudi", Place: 2},
	}

	bytes, err := json.Marshal(message)
	err = websocketConn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		fmt.Println("Error writing to ws connection !", err.Error())
	}

```

## Message Consumer 
```go
    type Message struct {
        Type string          `json:"type"`
        X    json.RawMessage `json:"sub_message"`
    }

    type SubMessage struct {
        Name  string `json:"name"`
        Place int    `json:"place"`
    }

    message := Message{}
    err = json.Unmarshal(dataBytes, &message)
    if err != nil {
    fmt.Println("Error unmarshalling...", err.Error())
    }
    
    fmt.Println("Level 1 Unmarshalling...", "\n", "message.Type : ", message.Type, "\n", "message.X : ", message.X)
    
    m := SubMessage{}
    err = json.Unmarshal(message.X, &m)
    if err != nil {
    fmt.Println("Error ", err.Error())
    }

    fmt.Println("Level 2 Unmarshalling...", "\n", "m.Name : ", m.Name, "\n", "m.Place : ", m.Place)
```