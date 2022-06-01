![](art/telegraph.png)

```Telegraph``` is a simple messaging server written in Go. It utilizes [GraphQL](https://github.com/99designs/gqlgen), a query language for APIs, to handle user requests. Additionally, it makes use of libraries such as [Casbin](https://github.com/casbin/casbin), [Gorm](https://gorm.io/index.html), and [Twilio](https://github.com/twilio/twilio-go). 

### How to Run?

1. Obtain Twilio credentials and put them in the ./.env file.
2. Run ```open -a Docker``` to open Docker. 
3. Run ```sh start.sh``` to start the server. 

### License

```
MIT License

Copyright (c) 2022 Vitaliy Paliy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
