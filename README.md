## Behavioral Code Analysis

### Change frequency
```
git log --format=format: --name-only | egrep -v '^$' | sort | uniq -c \ 
 | sort -r | head -20
```

### Commits by Author
```
git shortlog -sn
```

## Questionnaire
A questionnaire is modeled as a linked list of questions, each containing a list of possible answers. 

```
Q1 -> Q2 -> Q3

Q1: A1, A2, A3, A4
```

Questions could depend on each other. Dependencies overrule the original order.

```
Q2 -> depends on Q1,A1
```

Thus Q2 is only a valid success of Q1 if it was answered with A1. State is maintained on the clients behalf. Thus in order to get the next valid successor of Q1 the client is required to pass all previous answers to the server:

```
GET /questions/Q1/successor
Request Body:
{
    "previous_answers": [
        {
            "question": "Q1",
            "answers": ["A1"]
        }
    ]
}
```

This endpoint returns Q2 only if Q1 was answered by A1. Otherwise Q3 is returned.

### Endpoints 
```
GET /questions

GET /questions/{id}

POST /answers/{id}
Request Body:
{
    "answers": ["A1"]
}

GET /questions/Q1/successor
Request Body:
{
    "previous_answers": [
        {
            "question": "Q1",
            "answers": ["A1"]
        }
    ]
}
```
