## golang_server_python_client_protobuf

#### Simple Golang HTTP Server and Python HTTP Client. Both communicates using Protocol Buffer

Main outcome of this repo is to use the Protocol Buffers (Protobuf) in the REST API.

The HTTP Server is written in Golang and the HTTP Client is our Python. This two communicates using Protobuf.(Atleast the Server will always respond with the Protobuf message serialized to bytes)

The Client which receivers the bytes then Deserialize to the Protobuf message.

Folder Structure

```
.
|
|-http_client
|    |_client.py
|    |_pb_files
|        |_server_client_pb2.py
|-http_server
|    |_main.go
|-proto_files  
|    |_server_client.proto
|-README.md  
|_src
  |_httpserver
        |_golang
            |_protobuf
                |_server_client.pb.go
                
```

Details of the Files and Folders:

  - http_client - contains the Python HTTP client file and also the compiled proto file output in the pb_files folders

  - http_server - contains the Golang HTTP server file

  - proto_files - has the proto files

  - src - directory for the Golang PB file compiled by protoc

Prerequisites:

  - protoc installed
  
  - protoc-go-gen Plugin installed

  - protobuf - for both the Python and Golang


PATHS:
  
  - Append the Project's root directory path to the GOPATH and also PYTHONPATH if required


Usage:

  - To start the HTTP server, navigate to the _http_server_ folder and execute

    ```
      go run main.go
    ```
    
  - Now, the HTTP server is up and running in the **localhost:8090**. This URL has been provided to the Python Client file. To run the client, navigate to the _http_client_ folder and execute

    ```
      python client.py
             or
      python3 client.py
    ```
    
    
    