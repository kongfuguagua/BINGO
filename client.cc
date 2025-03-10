#include "hello.grpc.pb.h"
#include <string>
#include <iostream>
#include <sstream>
#include <grpcpp/grpcpp.h>

// 客户端
class HelloClient
{
public:
    HelloClient(std::shared_ptr<grpc::Channel> channel) : _stub(HelloWorld::NewStub(channel))
    {}

    std::string SayHello(const std::string& name)
    {
        HelloRequest req;
        HelloReply rsp;
        req.set_name(name);
        grpc::ClientContext context;
        grpc::Status status = _stub->SayHello(&context, req, &rsp);
        if(status.ok())
        {
            return rsp.message();
        }
        else
        {
             std::ostringstream out;
             out << status.error_code() << " : " << status.error_message();
             return out.str();
        }
    }

    std::string SayHelloAgain(const std::string& name)
    {
        HelloRequest req;
        HelloReply rsp;
        req.set_name(name);
        grpc::ClientContext context;
        grpc::Status status = _stub->SayHelloAgain(&context, req, &rsp);
        if(status.ok())
        {
            return rsp.message();
        }
        else
        {
            std::ostringstream out;
            out << status.error_code() << " : " << status.error_message();
            return out.str();
        }
    }

private:
    std::unique_ptr<HelloWorld::Stub> _stub;
};

// 调用rpc
int main(int argc, char** argv)
{
    std::string address("0.0.0.0:5000");
    HelloClient client(grpc::CreateChannel(address, grpc::InsecureChannelCredentials()));
    std::string name("bazel");
    std::cout << client.SayHello(name) << std::endl;
    std::cout << client.SayHelloAgain(name) << std::endl;
    return 0;
}