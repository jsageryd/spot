# Spot
Provides an API to let clients generate PDFs from [Graphviz
dot](http://www.graphviz.org/) files.

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
