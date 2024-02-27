# radis-with-go

1.Update Package Index:

sudo apt update

2.Install Redis Server:

sudo apt install redis-server

3. Start Redis Service:

sudo systemctl status redis

sudo systemctl start redis

sudo systemctl enable redis

4.Verify Installation:

redis-cli ping




Sample Output:

vidkrix@venky-worker:~/Desktop/work/radis-with-go$ go run main.go
Connected to Redis: PONG
Set key successfully
Value for key 'mykey': Hello, Redis!
Set key with expiry successfully
Key 'mykey_with_expiry' has expired
Closed Redis client
vidkrix@venky-worker:~/Desktop/work/radis-with-go$
