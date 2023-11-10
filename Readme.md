## Go Gin Starter

This is a starter boilerplate for the Gin framework in the Go programming language. This boilerplate serves as a foundational template for developing web applications using the Gin framework. It includes essential configurations, folder structures, and often-used components, providing developers with a kickstart for their Gin-based projects. With this boilerplate, you can quickly begin building web applications with Gin, saving you time and effort on initial setup and configuration tasks.

## Features
- `.env` file
- Live Reload
- Viper
- I18n
- Sentry
- Zap Logger

## Running the App

We utilize [wgo](https://github.com/bokwoon95/wgo) for seamless live reload support. 

To initiate live reloading, simply run the `wgo run main.go` command from the root directory of your project.

## Handle Errors

For custom error responses, refer to the routes in the `/api/errors-example.go` path. 

Additionally, for server-related errors, it's advisable to utilize the `panic` function to properly signal and manage the issue.

## Tasks
- [x] Integrate Viper
- [x] Docker
- [ ] Tests 
