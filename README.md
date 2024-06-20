# An interpreter for the Elliott programming language.

To run the interpreter from terminal
```
make run
```

### Testing
All test files are ran in bulk by running
To run tests
```
make test
```

If you want to run an individual direectory test you can run
```
go test ./parser
go test ./ast
go test ./evaluator
go test ./lexer
go test ./object
```
