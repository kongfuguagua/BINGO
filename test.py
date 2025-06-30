import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
try:
    s.bind(("192.168.3.61", 5051))
    print("绑定成功")
except OSError as e:
    print(f"绑定失败: {e}")  # 具体错误会直接打印
finally:
    s.close()