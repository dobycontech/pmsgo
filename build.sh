echo "TEST 1:"
go run bin/main.go --file="./test/ps.json"
echo "TEST 2:"
go run bin/main.go --file="./test/ps.json" --monit
echo "TEST 3:"
go run bin/main.go --file="./test/ps.json" --stop
echo "TEST 4:"
go run bin/main.go --file="./test/ps.json" --monit
echo "TEST 5:"
go run bin/main.go --file="./test/ps.json" --stop