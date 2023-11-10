# timg
A simple to use terminal image renderer.

- uses unicode characters over double the width for images

## TODO
[ ] determine a good pre-alloc size
[ ] transparent background of image
[ ] toggle between ascii only '  ' and '▀'
[ ] all the ansi, ansi256 and true color have a lot of shared code


## Examples

### Render image as is in it's original size
```go
var img image.Image

img = ... // load image

// render the image (turn it to a string)
// Note:  timg.Render will auto detect what color
//        profile the environment supports (using termenv)
println(timg.Render(img))
```

### Render image but make it fit a certain size
```go
var img image.Image

img = ... // load image

// resize image to fit in a 25x25 cell box
println(timg.Render(img, timg.FitTo(25, 25))
```

### Other examples
*see ./examples*
```
$ cd ./examples
$ go run ./resize/main.go /path/to/your/image
```


## References
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797?permalink_comment_id=4619910#gistcomment-4619910
- https://chrisyeh96.github.io/2020/03/28/terminal-colors.html

