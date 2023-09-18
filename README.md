# fetch-challenge

## Intro

If you are reading this you are probably an engineer working at fetch :)

First of all thanks for taking the time to review my code, I hope you find my solution is up to your standards, and if not at least I learned a lot and had fun.

## How to run the project

I've been told that you already have go installed in your computer but if for some reason that is not the case you could install the corresponding version according to your OS in [The official golang page](https://go.dev/dl/).

after you have downloaded go in your machine you could create a folder anywhere you want, and open your terminal in that folder

    cd \your\folder\path

then clone my project using this command while inside the folder you created.

    git clone https://github.com/RubenVillalpando/fetch-challenge.git

after the repo finishes cloning you change into the directory of the code

    cd fetch-challenge\src

and from there you could either run the project by running the command:

    go run .

to test the endpoints with your own receipts, you could create ones using the structs from the `schemas.go` file on the src folder or with this example

    curl --location '127.0.0.1:8080/receipts/process' \
    --header 'Content-Type: application/json' \
    --data '{
    "retailer": "M&MCornerMarket",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
        {
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        }
    ],
    "total": "9.00"
    }'

Once you have the receipt in memory you could get the points from said receipt using the id you get back from the response

    curl --location '127.0.0.1:8080/receipts/{{your receipt id}}/points'

or run the tests written by me by using the command:

    go test .

which runs all the tests in the project, which are just the tests in the receipts_test.go file

## A little bit about me and the project

### Learned go for the project

I've been meaning to learn go for a while and saw the challenge as an opportunity to learn it since I've heard tons of good stuff about the language. Even though in terms of time, I was probably better off using a language I was already comfortable with, I decided to use go because in case I did good and get the job, I already know the language the company uses, and if not I at least learned something new!

### My journey through the project

First it took me a couple of hours to do some hello world programs in go and using gin framework to get a feel for backend development in golang. then over the course of many days, working just a few minutes or an hour at a time due to busy days I developed the project, adding all the time it took me would be to around 10-12 hours.

Also some things to note

- I noticed that an example for the points had a retailer name with spaces but the pattern used to validate the retailer name specified no spaces, so I stuck to the specification although I do not know if that was the case.

- Also I used chat gpt for creating the receipts for the tests but for everything else I stuck to documentation and my brain.
