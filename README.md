# Shutup docker compose! 

When running a docker-composed project in Jenkins where the environment is shared you may end up bashing into each other's exposed ports. 

Using `ports` locally is useful as you can run your app and look at it in the browser, but you dont want to do that in Jenkins when you're just running a build

This tool will replace all `ports` declarations with `expose` in a really naive way.
 
`go run main.go example.yaml > somefile`