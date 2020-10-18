# haldiayatra

go mod init github.com/vrindavanstudio/haldiayatra
go get go.mongodb.org/mongo-driver/mongo
go get -u github.com/cosmtrek/air



References:
https://docs.mongodb.com/drivers/go               -->mongodb drivers refers to mongodb connection and usage

https://github.com/golang/go/wiki/Modules         --> go modules usage
https://github.com/mongodb/mongo-go-driver#usage  --> mongodb connection and usage
https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f --> golang singleton implementation
Note: we can also make db global variable and use Init in main fn to initialize at the start of the program and everyone can use that global var instance
Another approach is thread safe sync.Once

https://github.com/cosmtrek/air  Air for gin live reload
https://medium.com/wesionary-team/jwt-authentication-in-golang-with-gin-63dbc0816d55 for Jwt with gin


Issues:
panic: error parsing uri: lookup cluster0.ocshf.mongodb.net on 127.0.0.53:53: cannot unmarshal DNS message
Seen in Ubuntu 18.04LTS 

edited /etc/resolv.conf
nameserver 8.8.8.8 

it worked but it may have some persistence issue afer reboot may be overwritten refer this soln
https://stackoverflow.com/questions/55660134/cant-connect-to-mongo-cloud-mongodb-database-in-golang-on-ubuntu

