go build -o main main.go
./main &
SRV_PID=$!
echo "start server, pid $SRV_PID"

sleep 1s

echo "send make users requests"
ab -n 1000 -c 10 http://localhost:8080/?cnt=10000 >/dev/null 2>&1

echo "record heap profile"
wget -O trace.out http://localhost:6060/debug/pprof/heap >/dev/null 2>&1

echo "kill server, pid $SRV_PID"
kill $SRV_PID
sleep 1s
rm main

echo "done, analize profile with 'go tool pprof -alloc_space trace.out'"