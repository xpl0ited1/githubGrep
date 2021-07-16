# GitHub Grep v0.1
![](https://img.shields.io/maintenance/yes/2021?style=plastic) 
![](https://img.shields.io/github/languages/top/xpl0ited1/githubGrep?style=plastic)
![](https://img.shields.io/github/issues/xpl0ited1/githubGrep?style=plastic)
![](https://img.shields.io/github/license/xpl0ited1/githubGrep?style=plastic)

<hr/>

### Description

This tool will allow you to search for any code on a specific repository or organization and even a user

### Install

```
go get github.com/xpl0ited1/githubGrep
```

### Usage

```
$ githubGrep -search GITHUB -user xpl0ited1 -lang Go

  -content
        display content of code
  -lang string
        programming language
  -org string
        organization to look at
  -page int
        page number, only if results are more than 100 (default 1)
  -repo string
        repo to look at
  -search string
        code to search
  -user string
        user to look at


```

### Examples

![img.png](img.png)

![img_1.png](img_1.png)