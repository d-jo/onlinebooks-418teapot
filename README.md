# OnlineBooks

## Development
First, install Go at `https://golang.org/doc/install`. This process depends on your operating system

I strongly recommend VSCode for writing Go. It is a free IDE from Microsoft. After installing VSCode, you should install the Go extension using the four squares in the left riboon for better integration.

Before continuing, run the commands `go version` and `go env` to confirm the installation.

Once Go is installed, clone this repository. 

How is the project structured? We are using Go modules. All the project dependencies are stored in the `go.mod` file. Use the command `go mod download all` to automatically download and install the dependencies. You should do this before writing any code.

Before running the code, you should rename creds_sample.json to creds.json

The config.json file stores configuration information and SQL,
 the creds.json file stores credentials. Look at structs.go to see what fields are available. The json:"" lets go know how to read and write json

Setting up the SQL user :

`CREATE USER 'onlinebooks'@'localhost' IDENTIFIED BY 'onlinepassword';`

`GRANT ALL PRIVILEGES ON * . * TO  'onlinebooks'@'localhost';`

`FLUSH PRIVILEGES;`

`CREATE DATABASE onlinebooks`


(you can change the username/password if you want, make sure it is correct in the creds.json)

After creating the SQL user, edit the creds.json file with the correct username and password. If the creds.json file doesnt exist, copy the sample and rename it. Add you username and password.

How to run the code: Using the terminal at the bottom of VSCode, you can do `go run github.com/onlinebooks-418teapot` to start the application.

Workflow:
GitHub branch workflow
https://guides.github.com/introduction/flow/


