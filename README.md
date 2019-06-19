# 🖥️ desktop-app


## Running Development Apps
```bash
go run main.go 

## OUTPUT
# package main

# import "github.com/zserge/webview"

# func main() {
# 	webview.Open("My Desktop App", "https://medium.com/@david.richard.holtz/how-facebook-is-going-to-change-crypto-forever-155498c22b0", 900, 900, true)
# }
```

Now run the desktop app
```bash
go run app.go
```


# Building Production Apps

Lets build the CLI app as a binary
```bash
go build main.go
```


Now we make a new desktop app
```bash
./main -url="https://medium.com/@david.richard.holtz/wifi-cracking-ii-with-tutorial-5159b1ca385" -title="WIFI CRACKING" -height=1000 -width=1000
```

A new file name app.go is generated

```bash
go build app.go 
```

now we have a compiled desktop app!

```bash
./app
```

# A single line commnad

```bash 
go build main.go && ./main -url="https://medium.com/@david.richard.holtz/wifi-cracking-ii-with-tutorial-5159b1ca385" -title="WIFI CRACKING" -height=1000 -width=1000 && go build app.go && ./app
```