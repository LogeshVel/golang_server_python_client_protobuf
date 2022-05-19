from pb_files import server_client_pb2 as pb_client
import requests
from google.protobuf.text_format import Parse, MessageToBytes

def request_n_parse_response(request_url, request_method, pb_obj, query_param = None, payload=None):
    """
    method: to make the http request to the given url with the given http method and parse the response from the server to the pb message
    """
    available_http_methods = ["get", "post", "put", "delete"]
    if request_method.lower() not in available_http_methods:
        print("Please Provide the Correct HTTP request methods. Available methods ",available_http_methods)
        return
    if request_method.lower() == "get":
        if query_param is None:
            response = requests.post(url = request_url)
            print(response.headers)
            parse_response(response,pb_obj)
            return
        response = requests.get(url = request_url, params=query_param)
        parse_response(response,pb_obj)
    elif request_method.lower() == "post":
        if payload is None:
            response = requests.post(url = request_url)
            print(response.headers)
            parse_response(response,pb_obj)
            return
        
        response = requests.post(url = request_url, data=payload)
        print(response.headers)
        parse_response(response,pb_obj)

def parse_response(http_response_obj, pb_obj):
    print("EchoResponse content ", http_response_obj.content)
    print("EchoResponse content type : ", type(http_response_obj.content))
    pb_obj.ParseFromString(http_response_obj.content)
    print(pb_obj)


# request root page
root_url = "http://localhost:8090"
echo_url = root_url + "/echo"
# request_n_parse_response(root_url,"get",pb_client.RootPage())
# request_n_parse_response(echo_url,"get",pb_client.EchoResponse())
# request_n_parse_response(echo_url,"get",pb_client.EchoResponse(), query_param={"query-key-1":"query-value-1-by-PYTHON"})
# request_n_parse_response(echo_url,"post",pb_client.EchoResponse())
request_n_parse_response(echo_url,"post",pb_client.EchoResponse(), payload={"Payloadd":"This is my payload"})

# response = requests.get(url = root_url)

# echo_response = requests.get(url = echo_url)
# echo_response_content = echo_response.content


# echo_url = root_url + "/echo"
# echo_response = requests.post(url = echo_url)
# echo_response_content = echo_response.content
