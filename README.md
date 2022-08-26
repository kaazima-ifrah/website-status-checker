# website-status-checker

### :memo: &nbsp; Problem Statement:

To build a service that checks the status of websites continuously and expose APIs to add websites and retrieve statuses.

### :gear: &nbsp; Sample cURL commands for the APIs:

#### /add-websites (POST)

- Accepts the list of websites in the request body as an array
```
curl -v -d '[{"Url":"https://www.google.com"},{"Url":"https://www.amazon.com"},{"Url":"https://www.airbnb.com"}]' "127.0.0.1:3000/add-websites"
```

#### /view-websites-status (GET)

- Without query parameter - Returns the websites along with their corresponding status
```
curl -v "127.0.0.1:3000/view-websites-status"
```

- With query parameter - Returns the status of the website passed as a query parameter
```
curl -v "127.0.0.1:3000/view-websites-status?name=https://www.google.com"
```
