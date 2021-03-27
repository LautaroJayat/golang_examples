## Unused variables

This is silly but cool stuff
`./main.go` has a problem, it has an unused variable. The compiler will throw an error because it knows the scope of your function and also knows if you are using all variables declared inside of those curly braces.
On the other hand, there is another unused variable. But this one is in the package scope.
If you check `./unused_but_chill.go` you will se a unused variable declared and exported variable outside any function.

This one wont cause any troubles, because go compiler cant know if that variable will be used in another package. The scope of that variable, and also it's role in the program can't be defined in a simple static analysis  of the package