#include <string.h>
#include <grpcpp/grpcpp.h>
#include <grpcpp/channel.h>
#include <grpcpp/client_context.h>
#include <grpcpp/create_channel.h>

#include "route_guide.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using routeguide::Feature;
using routeguide::Point;


class RouteGuideClient {
    public:
    RouteGuideClient(std::shared_ptr<Channel> chan) : stub_(routeguide::NewStub(chan)){}

    int sendRequest(google::protobuf::int32 lat, google::protobuf::int32 long) {
        Point p;
        p.set_latitude(lat);
        p.set_longitude(long);

        Feature feature;
        ClientContext context;

        Status status = stub_->sendRequest(&context, p, &feature);

        if (status.ok()) {
            return reply.result();
        }
    }
}


void run() {
    std::string address("localhost:10000");

    RouteGuideClient client(grpc::CreateChannel(
        address,
        grpc::InsecureCredentials()
    ));

    auto response = client.GetFeature()
    
    
}


int main() {
    run();
    return 0;

}