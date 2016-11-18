Simple script to show and open your gist files in terminal.

### Install

```golang
go get "github.com/lccezinha/gistfinder"
```

#### Disclaimer

> You need a github token with **repo** permission allowed. More info [here](https://help.github.com/articles/creating-an-access-token-for-command-line-use/)

### How to use

After install script, just run the command `gistfinder` in your terminal and then select the number of Gist file you wanna open in browser:

#### Options

You can use options when run the `gistfinder` command with options, for now only `-p` option is available:

```
gistfinder -p=false
```

More informations about options:

```
gistfinder -h

Usage of gistfinder:
  -p	Use false value to do not list private gists. (default true)
```

![gistfinder](http://i.imgur.com/5MRD31N.png)
