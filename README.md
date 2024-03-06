Steps

## Port forward your one of the mongos 
`kubectl port-forward -n demo pod/mgs-mongos-0 27017`

## View DB creds
`kubectl view-secret -n mongodb mongo-test-auth -a`

## Export envs
export MONGODB_USERNAME=root
export MONGODB_PASSWORD=<pass from above view-secret command>

## RUN
`go run main.go`

## Share 
All the logs are stored in `./log` directory. Share it with us. 
