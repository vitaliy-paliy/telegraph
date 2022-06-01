![](art/telegraph.png)

```Telegraph``` is a simple messaging server written in Go. It utilizes [GraphQL](https://github.com/99designs/gqlgen), a query language for APIs, to handle user requests. Additionally, it makes use of libraries such as [Casbin](https://github.com/casbin/casbin), [Gorm](https://gorm.io/index.html), and [Twilio](https://github.com/twilio/twilio-go). 

### How to Run?

1. Obtain Twilio credentials and put them in the ./.env file.
2. Run ```open -a Docker``` to open Docker. 
3. Run ```sh start.sh``` to start the server. 
