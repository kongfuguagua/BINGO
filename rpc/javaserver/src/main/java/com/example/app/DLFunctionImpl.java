package com.example.app;

import io.grpc.stub.StreamObserver;
import dl.Dl.DLapp;
import dl.Dl.DLMetadata;
import dl.Dl.DLSpec;
import dl.Dl.DLModel;
import dl.Dl.DLDataOBJ;
import dl.Dl.DLGetRequestById;
import dl.Dl.DLGetResponseById;
import dl.Dl.DLCreateRequest;
import dl.Dl.DLCreateResponse;
import dl.DLfunctionGrpc.DLfunctionImplBase;

public class DLFunctionImpl extends DLfunctionImplBase {
    @Override
    public void getDLById(DLGetRequestById request, StreamObserver<DLGetResponseById> responseObserver) {
        System.out.println("Received getDLById request for ID: " + request.getId());
        
        // 创建一个测试响应
        DLapp.Builder dlAppBuilder = DLapp.newBuilder();
        
        // 设置元数据
        DLMetadata.Builder metadataBuilder = DLMetadata.newBuilder()
            .setId(request.getId())
            .setNamespace("test-namespace")
            .setDLName("test-dl");
        
        // 设置规格
        DLSpec.Builder specBuilder = DLSpec.newBuilder();
        
        // 设置模型
        DLModel.Builder modelBuilder = DLModel.newBuilder()
            .setStatus(true)
            .setPath("/path/to/model");
        specBuilder.setModel(modelBuilder);
        
        // 设置数据对象
        DLDataOBJ.Builder dataObjBuilder = DLDataOBJ.newBuilder()
            .setStatus(true)
            .setDatabase("test-database");
        specBuilder.setDataObj(dataObjBuilder);
        
        // 构建完整的DLapp
        dlAppBuilder.setMetadata(metadataBuilder)
                   .setSpec(specBuilder);
        
        // 创建响应
        DLGetResponseById response = DLGetResponseById.newBuilder()
            .setDlApp(dlAppBuilder)
            .build();
        
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void createDL(DLCreateRequest request, StreamObserver<DLCreateResponse> responseObserver) {
        System.out.println("Received createDL request");
        
        // 生成一个唯一ID
        String id = String.valueOf(System.currentTimeMillis());
        
        // 使用请求中的信息创建新的DLapp
        DLapp.Builder dlAppBuilder = DLapp.newBuilder();
        
        // 设置元数据
        DLMetadata.Builder metadataBuilder = DLMetadata.newBuilder()
            .setId(id)
            .setNamespace(request.getSpec().getNamespace())
            .setDLName(request.getSpec().getDLName());
        
        // 设置规格
        DLSpec.Builder specBuilder = DLSpec.newBuilder();
        
        // 设置默认的模型和数据对象
        DLModel.Builder modelBuilder = DLModel.newBuilder()
            .setStatus(true)
            .setPath("/default/model/path");
        specBuilder.setModel(modelBuilder);
        
        DLDataOBJ.Builder dataObjBuilder = DLDataOBJ.newBuilder()
            .setStatus(true)
            .setDatabase("default-database");
        specBuilder.setDataObj(dataObjBuilder);
        
        // 构建完整的DLapp
        dlAppBuilder.setMetadata(metadataBuilder)
                   .setSpec(specBuilder);
        
        // 创建响应
        DLCreateResponse response = DLCreateResponse.newBuilder()
            .setDlApp(dlAppBuilder)
            .build();
        
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void initDL(DLapp request, StreamObserver<DLapp> responseObserver) {
        System.out.println("Received initDL request for DL: " + request.getMetadata().getDLName());
        
        // 创建响应，基于输入请求进行修改
        DLapp.Builder responseBuilder = request.toBuilder();
        
        // 更新模型状态
        DLSpec.Builder specBuilder = responseBuilder.getSpec().toBuilder();
        specBuilder.setModel(specBuilder.getModel().toBuilder().setStatus(true));
        specBuilder.setDataObj(specBuilder.getDataObj().toBuilder().setStatus(true));
        responseBuilder.setSpec(specBuilder);
        
        DLapp response = responseBuilder.build();
        
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
} 