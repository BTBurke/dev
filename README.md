dev
===

Dev is a simple command-line tool to make developing Golang apps in Docker containers easy.

### How it works:

- Starts containers to isolate your development environment and code
- Uses (http://goconvey.co)[go-convey] to set up automated testing on save
- Exposes the go-convey web interface so you can monitor your test status
- Creates convenience commands for installing dependencies and getting to a shell in your development container

### Why use this?

Previously, I used Vagrant for isolated development environments that mimic production infrastructure.  Since switching to containerized applications, I wanted an easy way to isolate development code and take advantage of using containers for other infrastructure, just like it runs in production.

## Quickstart

```bash
# Create a basic fig.yml environment for your dev environment and a 
# data-only container for your code

>> dev init 

# Start your development infrastructure

>> dev up

# Want to get a shell in your dev container?

>> dev shell

# Get to the test status web interface?

>> dev web

# See the output of your tests in the terminal?

>> dev test
```