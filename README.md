**Intro:**

main.go - contains logic to process transactions limited by total time 50ms, 60ms, 90ms, 1000ms
api_latencies is added as map

**Question:**
Write a `prioritize` function body, which returns a subset of `transactions` that have the maximum total USD.

Run following command to process transactions.csv
```
go run main.go
```

**Question:** 
What's is the max USD value that can be processed in 1 second?

Answer: Total amount: 11852.99

**Question:**

Modify the `prioritize` function to also accept the `totalTime` in milliseconds (default=1000ms).
Your implementation should correctly prioritize based on the `totalTime` argument.


**Question:**
What is the max USD value that can be processed in 50ms, 60ms, 90ms?

    Process transactions limited by total time: 50 ms
    Total amount: 999.17

    Process transactions limited by total time: 60 ms
    Total amount: 999.17

    Process transactions limited by total time: 90 ms
    Total amount: 999.17