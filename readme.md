

# Api

## Struct of message
```
type Request struct {
	Token string `json:"token"`
	To string `json:"to"`
	Title string `json:"title"`
	Channel string `json:"channel"`
	Text string `json:"text"`
	Image string `json:"image"`
	Color string `json:"color"`
}
```

## Types

### If "channel" defined:
```
{
  "token": "0x...ff",
  "channel": "portal",
  "text": "Hello"
}
```
Message in channel:  
<img src="https://vicuesoft.sfo2.digitaloceanspaces.com/images/basicbot.PNG">

### If "channel" not defined
```{
  "token": "0.....f",
  "to": "GattoDiLauro",
  "text": "Hello"
}
```
Message in private dialogue with bot:  
<img src="https://vicuesoft.sfo2.digitaloceanspaces.com/images/basicbot.PNG">

### If "color" or "image" defined

<img src="https://vicuesoft.sfo2.digitaloceanspaces.com/images/mediumbot.png">

If channel defined: message will be sent into channel (Higher priority than "to" field)
If channel not defined, but "to" defined: message will be sent personally to user