from concurrent import futures
import grpc
import mnist_pb2
import mnist_pb2_grpc
import torch

# 假设我们有一个字典来存储加载的模型
loaded_models = {}

class DLfunctionServicer(mnist_pb2_grpc.DLfunctionServicer):
    def initDLModel(self, request, context):
        model_path = request.path

        # 检测模型是否已经加载
        if model_path in loaded_models:
            model = loaded_models[model_path]
            status = True
        else:
            # 尝试加载模型
            try:
                model = torch.load(model_path)
                status = True
                loaded_models[model_path] = model
            except Exception as e:
                context.abort(grpc.StatusCode.INTERNAL, str(e))
                status = False

        # 如果模型加载成功，获取输入和输出维度
        if status:
            # 创建一个随机张量作为输入
            example_input = torch.randn(1, *model.input_size)  # 假设模型有一个input_size属性
            # 获取模型的输出
            example_output = model(example_input)
            # 获取输入和输出的维度
            input_dim = example_input.shape
            output_dim = example_output.shape
        else:
            input_dim = "N/A"
            output_dim = "N/A"

        # 构建并返回DLModel对象
        return mnist_pb2.DLModel(
            name=model.__class__.__name__,
            path=model_path,
            status=status,
            inputType=input_dim,
            outputType=output_dim
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    mnist_pb2_grpc.add_DLfunctionServicer_to_server(DLfunctionServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
