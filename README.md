# dision_test

This is the repository containing the test project for entry level in Dision Tech ltd

The app demonstrate my skill in ultilizing goroutine for concurrency and worker processes.

The application is also served on ipv6 and only accept https requests with client supporting Tls1.3

The app contain a single API end point that take in datas for some calculation, then use worker processes to process it

The body of the request is like below:

GET: https://::1:1234/api/v1/test?a=number&b=number

The response content some calculation on the query params input
