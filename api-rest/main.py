import requests
URL = 'http://localhost:8083/api/names'

# data = requests.get(URL, headers = {"Content-Type": "application/json"})


# data = requests.get(f'{URL}/4', headers = {"Content-Type": "application/json"})


dataJSON = {"name": "Wellington", "age": 36}
data = requests.post(URL, headers = {"Content-Type": "application/json"}, json = dataJSON)

# data = requests.delete(f'{URL}/2', headers = {"Content-Type":"application/json"})



# data = requests.patch(f'{URL}/4'
#     ,headers =  {"Content-type": "application/json"}
#     ,json = {'name': 'Wilson Costa', 'age': 35} 
# )

print(data.json())

