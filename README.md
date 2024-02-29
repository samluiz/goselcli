# Go Select CLI (goselcli)

Goselcli is a dead simple go library designed to easily build simple CLI's. It's purpose is not to be used to build complex command line interfaces.

# Installation

```bash
go get github.com/samluiz/goselcli@latest
```

# Usage

##### Import the menu:
```go
import "github.com/samluiz/goselcli/menu"
```

##### Create a new menu:
```go
m := menu.NewMenu("Select an option!")
```

##### Add options to the menu:
```go
m.AddOption("Option 1", "1")
m.AddOption("Option 2", "2")
m.AddOption("Option 3", "3")
m.AddOption("Option 4", "4")
m.AddOption("Option 5", "5")
```

##### Display the menu and retrieve the choice:
```go
choice := m.Display()

println("You chose:", choice)
```

##### Output:
<img width="160" alt="image" src="https://github.com/samluiz/goselcli/assets/97702597/c6f94cf9-d4e9-44cd-b930-c3b7efdc584c">


