# First step
You need to create bot in Discord App. Read manual in the internet, i don't care.  
Bot must be added into server.
# Interface

```cassandraql
curl -H "Content-type: application/json" -d "@D:\data.json" localhost:5000
```
where @data.json is 
```cassandraql
{
	"token": "0x00-0xff",
	"to": "devops",
	"text":"Hello!"
}
```

#### Known issues
```cassandraql
panic: runtime error: index out of range [0] with length 0
```
...Happend when token is wrong or connection to Discord failed.