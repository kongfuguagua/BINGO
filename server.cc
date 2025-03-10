#include "hello.grpc.pb.h"
#include <string>
#include <grpcpp/grpcpp.h>

// rpc服务实现
class HelloServiceImpl : public HelloWorld::Service
{
    grpc::Status SayHello(grpc::ServerContext* context, const HelloRequest* req, HelloReply* rsp)
    {
        std::cout << "Request SayHello From " << context->peer() << std::endl;
        rsp->set_message("hello " + req->name() + "!");
        return grpc::Status::OK;
    }

    grpc::Status SayHelloAgain(grpc::ServerContext* context, const HelloRequest* req, HelloReply* rsp)
    {
        std::cout << "Request SayHelloAgain From " << context->peer() << std::endl;
        rsp->set_message("hello " + req->name() + " again!!");
        return grpc::Status::OK;
    }

};

// 启动运行
int main(int argc, char** argv)
{
    std::string address("localhost:5000");
    HelloServiceImpl service;

    grpc::ServerBuilder builder;
    builder.AddListeningPort(address, grpc::InsecureServerCredentials());
    builder.RegisterService(&service);

    std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
    std::cout << "Server Listening on port: " << address << std::endl;

    server->Wait();

    return 0;
}