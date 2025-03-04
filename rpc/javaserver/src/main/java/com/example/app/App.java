package com.example.app;

import io.etcd.jetcd.ByteSequence;
import io.etcd.jetcd.Client;
import io.etcd.jetcd.lease.LeaseGrantResponse;
import io.etcd.jetcd.options.PutOption;

import java.nio.charset.StandardCharsets;
import java.util.concurrent.ExecutionException;

/**
 * 示例Java应用程序
 */
public class App {
    public static void main(String[] args) {
        System.out.println("你好，世界！");
        
        try {
            // 注册服务到etcd
            registerService();
            
            // 保持服务运行
            while (true) {
                System.out.println("服务运行中...");
                Thread.sleep(5000); // 每5秒打印一次
            }
        } catch (Exception e) {
            System.err.println("服务运行出错: " + e.getMessage());
            e.printStackTrace();
        }
    }

    /**
     * 将服务注册到etcd
     */
    private static void registerService() throws ExecutionException, InterruptedException {
        // 创建etcd客户端
        Client etcdClient = Client.builder().endpoints("http://127.0.0.1:2379").build();
        
        // 服务信息
        String serverIp = "127.0.0.1";
        int serverPort = 8080;
        
        // 创建租约（600秒，与Python示例相同）
        LeaseGrantResponse leaseResponse = etcdClient.getLeaseClient().grant(600).get();
        long leaseId = leaseResponse.getID();
        
        // 注册键值
        String key = "dl.rpc/10"; // 与Python示例相同的键格式
        String value = serverIp + ":" + serverPort;
        
        // 将键值放入etcd，并绑定租约
        ByteSequence keyBytes = ByteSequence.from(key, StandardCharsets.UTF_8);
        ByteSequence valueBytes = ByteSequence.from(value, StandardCharsets.UTF_8);
        
        PutOption putOption = PutOption.newBuilder().withLeaseId(leaseId).build();
        etcdClient.getKVClient().put(keyBytes, valueBytes, putOption).get();
        
        System.out.println("服务已注册到etcd: " + key + " -> " + value);
        
        // 启动一个线程定期续约
        Thread keepAliveThread = new Thread(() -> {
            try {
                while (!Thread.currentThread().isInterrupted()) {
                    etcdClient.getLeaseClient().keepAliveOnce(leaseId).get();
                    System.out.println("租约续约成功: " + leaseId);
                    Thread.sleep(300000); // 每5分钟续约一次
                }
            } catch (Exception e) {
                System.err.println("租约续约失败: " + e.getMessage());
            }
        });
        keepAliveThread.setDaemon(true);
        keepAliveThread.start();
        
        // 添加关闭钩子，在应用程序关闭时关闭etcd客户端
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("应用程序关闭，正在清理资源...");
            try {
                // 撤销租约，这将自动删除与此租约关联的所有键
                etcdClient.getLeaseClient().revoke(leaseId).get();
                System.out.println("已撤销租约: " + leaseId);
            } catch (Exception e) {
                System.err.println("撤销租约失败: " + e.getMessage());
            } finally {
                // 关闭etcd客户端
                etcdClient.close();
                System.out.println("etcd客户端已关闭");
            }
        }));
    }

    /**
     * 简单的加法方法，用于演示测试
     * @param a 第一个数
     * @param b 第二个数
     * @return 两数之和
     */
    public static int add(int a, int b) {
        return a + b;
    }
} 