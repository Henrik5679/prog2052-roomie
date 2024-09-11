# Intended Repository Structure

This file contains a cool graphic representation of how we want to
structure our repository. This is just a template, but we shall be nice
little devs and follow it to a degree which is reasonable and logical.

[https://go.dev/doc/modules/layout](https://go.dev/doc/modules/layout)

[https://go.dev/doc/modules/layout#server-project](https://go.dev/doc/modules/layout#server-project)

## Graphic Representation

```
// The Repo
-- root
   |   // Dockerfile for deployment anywhere
   +-- Dockerfile
   +-- Perhaps compose.yml for docker compose
   +-- .gitignore
   |
   |   // Secrets is for Docker to do confidential stuff.
   |   // Strictly in .gitignore; should **never** be included in a commit
   +-- .secrets
   |   +-- database key and other privileged files
   |
   |   // Basic go structure files right at root
   +-- go.mod
   +-- go.sum
   |
   |   // Some bash scripts for easy building/execution
   +-- execute.sh
   |
   |   // README.md which shows up on the repository front page
   |   // Basic information about the program, how to build, 
   |   // well wishes for our supervisor...
   +-- README.md
   | 
   |   // cmd is Go convention for executables
   +-- cmd
   |   +-- README.md // which explains a little about the different versions
   |   +-- full
   |   |   +-- main.go
   |   +-- stub (Don't know if this is necessary)
   |       +-- main.go
   | 
   |   // The front-end portion of our repo
   |   // (directory name is negotiable)
   +-- presentation
   |   |   // This README should include information pertinent to developers
   |   |   // such as what each module does, where to find certain information,
   |   |   // where to put what when developing further, et.c.
   |   +-- README.md
   |   +-- index.html/templ
   |   +-- reusable.css
   |   +-- reusable.js
   |   +-- consts.js
   |   |
   |   |   // Images. Don't know how we want to do THIS 
   |   |   // Maybe binary data such as pictures will
   |   |   // take up a lot of space in the repository...
   |   +-- shared-images
   |   |   +-- pic1.jpg
   |   |   +-- pic2.jpg
   |   |
   |   +-- Bogus1
   |   |   +-- bogus1.html
   |   |   +-- bogus1.css
   |   |   +-- bogus1.js
   |   |
   |   +-- Bogus2
   |       +-- bogus2.html
   |       +-- bogus2.css
   |       +-- bogus2.js
   |
   |   // The back-end portion of our repo
   +-- internal
       |   // README with information pertinent to developers
       |   // for this part of the program
       +-- README.md
       +-- consts and such
       +-- structs and such
       +-- shared functions and such
       |
       +-- Bogus1
       |   +-- bogus1.go
       |   +-- bogus1_test.go
       |
       +-- Bogus2
       |   +-- bogus2.go
       |   +-- bogus2_test.go
       |
       +-- Diagnostics
           +-- diagnostics.go
           +-- diagnostics_test.go
```
