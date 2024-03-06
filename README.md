Steps

## Port forward your one of the mongos 
`kubectl port-forward -n demo pod/mgs-mongos-0 27017`

## View DB creds
`kubectl view-secret -n mongodb mongo-test-auth -a`

## Export envs
```
export MONGODB_USERNAME=root
export MONGODB_PASSWORD=<pass from above view-secret command>
```

## RUN
- If you have linux/amd64 architecture: `./bin-linux-amd64`
- If you have darwin/arm64 architecture: `./bin-darwin-arm64`
- If you prefer go directly : `go run main.go`

## Share 
All the logs are stored in `./log` directory. Share it with us. 


```
db.Activity.stats()
```
- Are tables sharded or not?
- Are sharded balanced?
- Check the disk used size (in GB, not %) to confirm that shards are balanced.
- https://www.mongodb.com/docs/manual/core/sharding-choose-a-shard-key/

