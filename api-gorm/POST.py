# USE TO POST VALUES


import requests
print(

requests.post(
    url = 'http://localhost:8081/api/name'
    ,json = {"name": "Jessica Nascimento Costa", "age": 31}

    ).json()
) 