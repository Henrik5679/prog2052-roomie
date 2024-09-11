# Guide for Comments, Commits, and Branching

## Go Function Docstrings

We're using Doxygen-style function comments. It's a good template, and if
needed, we can use Doxygen to make a quick documentation file.

```go
// Function description; what the results of the function are.
//
// @param   paramName          parameter description; what the function expects to find
// @return  returnType         what users can expect the function to return
// @see     mod.functionName   relevant functions used
```

### Otherwise Documentation in Go

For that, refer to this:

[https://tip.golang.org/doc/comment](https://tip.golang.org/doc/comment)


## JS Function Docstrings

Based on example from this post:
https://stackoverflow.com/questions/34205666/utilizing-docstrings

```js
// Function description; what the results of the function are.
//
// @param  {paramtype}  paramname   Parameter description.
// @return {returntype} returnname  Return description
// @see                 functionname
```


## Commits

Commits to follow the "Conventional Commits" convention. The original is
meant to service an automated system of monitoring commits. Read more here:

[https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/)

We don't need to follow it exactly, but we shold follow the basic idea of
using a "type" keyword up front, the short description, a greater body, and
a footer for info on issues and such. Like this:

### Example

```
feat: Added an automatic sort to the list

Added function to sort list, and implemented it so it displays as sorted
on the front page.

Updates: #12
```

The first line appears as a headline for Git, and it's recommended to keep
this under 60 or so characters.

Footer informs if you're creating, updating, or closing certain issues.
"#12" and such makes it so you're referring to the issue by number as it's
stored on GitLab.

Merges and similar commit types are exempt from this convention.


### Relevant Keywords

Please stick to these:
- feat:  means you're adding or updating a feature.
- fix:   means you've fixed a bug.
- docs:  means you've added documentation of some sort, including comments.
- chore: means anything else. Such as changing indentaion, whitespace,
         refactoring, moved files, removed artifacts, et.c.


## Branching

```
   ╭────────╮  // main is sacred. main should only be merged
   │  main  │  // with dev if and only *if* dev is 100% working.
   ╰────────╯
       ^
       |          // dev is where we merge to and fro while working on our
      dev         // branches. Do not push directly to dev.
   /  |  |  \
 b0  b1  b2  b3   // Branch out your personal branched from dev, and then
    / |   |  |    // branch out from those again. Branch and merge often,
 b4  b5  b6  b7   // and do it to solve issues you've claimed.
```


## Versioning

```
v1.20.3
```

Versioning in the above style:
- First slot representing greater version number; 0 while we're in beta, 1
  once we have MVP, et.c.
- Seconds slot representing an update; an implementation of greater
  features and such. An update to this happens when we merge to main.
- Third slot representing a hotfix. A bug fix, smaller feature
  implementation, smaller visual changes, et.c.
