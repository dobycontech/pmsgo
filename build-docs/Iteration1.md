## Iteration 1.

### run process
touch /test/ps.json
{
    "name": "psname1",
    "bin":"./litehttpserver",
    "timeoutseconds":-1
}
go run bin/main.go --file="./test/jsonpath"
init->{
run by arg file
Load Json into struct process
Run process
}->exit

ok-crit {
    * the json is loaded into process
    * the process is running message
    * the process is really running and the process exited
}

### watch process

Watch process changes
go run . --file="./test/ps.json" --monit="psname1"
init->{
    run by arg monit
    Load os.process into process or nil
    print process status
}<-exit

ok-crit {
    * The json is loaded
    * don't run the process
    * show the status of the last process executed
}

### stop process
Stop running process
go run . --stop="psname1"
init->{
    run by arg stop
    Load os.process into process
    process stop
    run watch process
}

ok-crit{
    * json is loaded
    * process is stopped
    * show monit
}
