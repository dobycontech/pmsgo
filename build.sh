cd test/psname1
go build -o psname1 main.go
cd ..
cd ..
echo "TEST 1: Read json file and run ----------------"
go run bin/main.go --file="./test/ps.json"
echo "end--------------------------------------"
echo ""
echo "TEST 2: Monit the process running"
go run bin/main.go --file="./test/ps.json" --monit
echo "end--------------------------------------"
echo ""
echo "TEST 3: Stop the run of the file and show"
go run bin/main.go --file="./test/ps.json" --stop
echo "end--------------------------------------"
echo ""
echo "TEST 4: Monit with no files running"
go run bin/main.go --file="./test/ps.json" --monit
echo "end--------------------------------------"
echo ""
echo "TEST 5: stop again with no files running"
go run bin/main.go --file="./test/ps.json" --stop
echo "end--------------------------------------"
echo ""