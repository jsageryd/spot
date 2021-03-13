# Spot
Provides an API to let clients generate PDFs from [Graphviz
dot](http://www.graphviz.org/) files.

## Installation
```
go install github.com/jsageryd/spot@latest
```

## Usage
### Server
Spot has no config. Just start it.

```
$ spot
```

### Client
Spot only cares that the request method is POST and will blindly assume that the
body is dot data. Here is an example request using
[HTTPie](https://github.com/jkbrzt/httpie).

```
$ echo 'digraph G { Hello -> World }' | http post :4649 > graph.pdf
```

# Licence
Copyright (c) 2015 Johan Sageryd <j@1616.se>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
