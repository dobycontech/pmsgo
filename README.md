# pmsgo
Process management Services

## Iteration 1.

### run process
touch /test/ps.json
{
    "name": "psname1",
    "bin":"./litehttpserver",
    "timeoutseconds":-1
}
go run . --file="./test/jsonpath"
init->{
run by arg file
Load Json into struct process
Run process
}->exit

### watch process

Watch process changes
go run . --monit="psname1"
init->{
    run by arg monit
    Load os.process into process or nil
    print process status
}<-exit

### stop process
Stop running process
go run . --stop="psname1"
init->{
    run by arg stop
    Load os.process into process
    process stop
    run watch process
}

