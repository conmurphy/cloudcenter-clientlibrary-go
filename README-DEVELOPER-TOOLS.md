# Quick Start - Creation from JSON file 

For some situations it may be easier to have the configuration represented as JSON rather than conifguring individually. In this scenario you can either build the JSON file yourself or monitor the API POST call for the JSON data sent to CloudCenter. This can be achieved using the browsers built in developer tools. 

If using the developer tools process you will first need to create the resource required in the GUI. This will allow you to find the correct JSON structure and reuse this structure in any future calls.

### 1. Login to CloudCenter and navigate to the resource you are trying to create. In this example we will create a new user

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/add_user.jpg)

### 2. Open the developer tools window within your browser and navigate to the network tab

### 3. Fill in the required details

### 4. Select `Save` and you should see a success banner indicating the user was successfully created

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/successfull_creation.jpg)

### 5. Look through the developer tools window to find the call made to CloudCenter. In this example it is `users/`. 

### 6. You should see in the right hand widow that the `Request Method` is `POST`

### 7. Scroll to the bottom and look at the `Request Payload`. You can select `View source` to obtain the data in JSON format

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/developer_tools.jpg)

### 8. Save the JSON into a file 

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/json.jpg)

### 9. You can now parse this JSON file with the library
