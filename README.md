### Introdution
Auto commit your git repository by cron. You can deploy by the `repo.json`

### Usage
1. deploy config file 

`repo.json`
```json
{
    "repos": [
        {
            "path": "path to your git repository",
            "message": "commit message",
            "autoCommit": true,
            "at": "23:30:00"    
        }
    ]
}
```

2. run 
```
1. make build
2. make run 
```
