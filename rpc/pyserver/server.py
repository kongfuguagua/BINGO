# -*- coding: utf-8 -*-
import time
import grpc
import dl_pb2,dl_pb2_grpc
from  concurrent import futures
import etcd3

# 实现 proto 文件中定义的 SearchService
class DLserver(dl_pb2_grpc.DLfunctionServicer):
    # 实现 proto 文件中定义的 rpc 调用
    def createDL(self, request, context):
        # 创建DLModel对象
        dl_model = dl_pb2.DLModel(
            name=request.spec.modelName,
            path="",
            status=False,
            inputType="",
            outputType=""
        )

        # 创建DLDataOBJ对象
        dl_data_obj = dl_pb2.DLDataOBJ(
            Database="",
            name="",
            status=False,
            total=0,
            type=""
        )

        # 创建DLSpec对象
        dl_spec = dl_pb2.DLSpec(model=dl_model, dataObj=dl_data_obj)

        # 生成UUID
        dl_id = "test"

        # 创建DLMetadata对象
        dl_metadata = dl_pb2.DLMetadata(
            id=dl_id,
            Namespace=request.spec.Namespace,
            DLName=request.spec.DLName
        )

        # 创建DLapp对象
        dl_app = dl_pb2.DLapp(metadata=dl_metadata, spec=dl_spec)

        # 返回DLCreateResponse对象
        return dl_pb2.DLCreateResponse(dlApp=dl_app)

def serve():
    # 启动 rpc 服务，这里可定义最大接收和发送大小(单位M)，默认只有4M
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options=[
        ('grpc.max_send_message_length', 100 * 1024 * 1024),
        ('grpc.max_receive_message_length', 100 * 1024 * 1024)])
    
    dl_pb2_grpc.add_DLfunctionServicer_to_server(DLserver(), server)
    server.add_insecure_port('127.0.0.1:50051')
    server.start()
    # etcdClient=etcd3.client()
    # server_ip='127.0.0.1'
    # server_port=50051
    # etcdClient.put(key='dl.rpc/3424113',value=f'{server_ip}:{server_port}')
    register()
    while True:
        try:
            print("starting")
            time.sleep(5) # one day in seconds
        except Exception as e:
            print(e)

def register():
    etcdClient=etcd3.client()
    server_ip = '127.0.0.1'
    server_port = 50051
    lease=etcdClient.lease(600)
    key = f'dl.rpc/{10}'
    value = f'{server_ip}:{server_port}'
    etcdClient.put(key, value)

if __name__ == '__main__':
    serve()