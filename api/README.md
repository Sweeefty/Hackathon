# Init the project 
```sh-session
    go mod init Hackathon/api
    go get github.com/mattn/go-sqlite3
    go mod tidy
```

# Start the api 
```sh-session
    go run main.go
```

# Connect to the API 

- First you need to have a admin account to get a Token (for the dev all account are admin)

- Create a Token with the endpoint ```/createToken``` with email and password as parameter like on this exexmple: ```/createToken?email=gurvan&password=gurvan```

- After this you recive a Token in response and you set this token in your requests header in ```Authorization=YourToken```

- Well played you now use the api