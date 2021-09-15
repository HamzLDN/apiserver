# openapi-client
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

This Python package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.PythonClientCodegen

## Requirements.

Python >= 3.6

## Installation & Usage
### pip install

If the python package is hosted on a repository, you can install directly using:

```sh
pip install git+https://github.com/GIT_USER_ID/GIT_REPO_ID.git
```
(you may need to run `pip` with root permission: `sudo pip install git+https://github.com/GIT_USER_ID/GIT_REPO_ID.git`)

Then import the package:
```python
import openapi_client
```

### Setuptools

Install via [Setuptools](http://pypi.python.org/pypi/setuptools).

```sh
python setup.py install --user
```
(or `sudo python setup.py install` to install the package for all users)

Then import the package:
```python
import openapi_client
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```python

import time
import openapi_client
from pprint import pprint
from openapi_client.api import user_api
from openapi_client.model.main_create_user_input import MainCreateUserInput
from openapi_client.model.main_create_user_output import MainCreateUserOutput
from openapi_client.model.main_get_users_output import MainGetUsersOutput
from openapi_client.model.main_update_user_body import MainUpdateUserBody
from openapi_client.model.main_user import MainUser
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)



# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = user_api.UserApi(api_client)
    body = MainCreateUserInput(
        name="name_example",
        nick_name="nick_name_example",
        phone="phone_example",
    ) # MainCreateUserInput | 

    try:
        # create user
        api_response = api_instance.func1(body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling UserApi->func1: %s\n" % e)
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*UserApi* | [**func1**](docs/UserApi.md#func1) | **POST** /api/user | create user
*UserApi* | [**func2**](docs/UserApi.md#func2) | **GET** /api/user | search/list users
*UserApi* | [**func3**](docs/UserApi.md#func3) | **GET** /api/user/{user-name} | get user
*UserApi* | [**func4**](docs/UserApi.md#func4) | **PUT** /api/user/{user-name} | update user
*UserApi* | [**func5**](docs/UserApi.md#func5) | **DELETE** /api/user/{user-name} | delete user


## Documentation For Models

 - [MainCreateUserInput](docs/MainCreateUserInput.md)
 - [MainCreateUserOutput](docs/MainCreateUserOutput.md)
 - [MainGetUsersOutput](docs/MainGetUsersOutput.md)
 - [MainUpdateUserBody](docs/MainUpdateUserBody.md)
 - [MainUser](docs/MainUser.md)


## Documentation For Authorization

 All endpoints do not require authorization.

## Author




## Notes for Large OpenAPI documents
If the OpenAPI document is large, imports in openapi_client.apis and openapi_client.models may fail with a
RecursionError indicating the maximum recursion limit has been exceeded. In that case, there are a couple of solutions:

Solution 1:
Use specific imports for apis and models like:
- `from openapi_client.api.default_api import DefaultApi`
- `from openapi_client.model.pet import Pet`

Solution 2:
Before importing the package, adjust the maximum recursion limit as shown below:
```
import sys
sys.setrecursionlimit(1500)
import openapi_client
from openapi_client.apis import *
from openapi_client.models import *
```
