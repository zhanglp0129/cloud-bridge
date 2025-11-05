# cloud-bridge
## 项目简介

这是一个用Go语言开发的桌面客户端，作为云盘和用户本地文件系统的桥梁。项目包含两个部分：

1. **mock-server**: 模拟云端认证和配置下发服务
2. **client-bridge**: 核心客户端，负责登录、挂载云存储，并提供本地API服务

## 环境依赖

### 必需依赖
- **Go 1.24.9**: https://golang.google.cn/dl/
- **rclone**: 文件同步工具

### 平台特定依赖

#### Windows
- **WinFSP**: 用于FUSE文件系统支持
  - 下载: https://github.com/winfsp/winfsp/releases
  - 安装后需要重启

#### macOS
- **macFUSE**: 用于FUSE文件系统支持
  - 安装: `brew install macfuse`
  - 或下载: https://osxfuse.github.io/

#### Linux (Ubuntu/Debian)
- **FUSE**: 通常已预装，如未安装：
``` bash
sudo apt update
sudo apt install fuse
```
### 安装rclone

```bash
# 使用包管理器或从官网下载
macOS
brew install rclone
# Linux (Ubuntu/Debian)
sudo apt install rclone
```
或从官网下载: https://rclone.org/downloads/
#### 验证安装
```bash
rclone version
```
## 运行步骤

### 1. 启动Mock Server
```bash
# 运行服务器
go run ./mock-server
```
服务器默认运行在 `http://0.0.0.0:8080`

### 2. 运行Client Bridge
```bash
go run ./client-bridge
```

### 3. 配置选项
在配置文件`./config.yaml`中填写配置
```yaml
server_url: "http://127.0.0.1:8080"
username: user
password: pass
mount_path: "X:" # 挂载路径
```

## 测试方法

### 1. 验证挂载状态

#### 检查API状态
```bash
# 检查健康状态
curl http://localhost:9527/health
# 检查挂载状态
curl http://localhost:9527/status
```
#### 手动验证文件同步
1. 在模拟的远程存储中创建文件：
```bash
echo "Hello from cloud" > /tmp/mock-remote-storage/test.txt
```
2. 检查挂载点是否出现文件：
```bash
ls ~/MyCloudDrive/
cat ~/MyCloudDrive/test.txt
```
3. 在挂载点创建文件，检查是否同步到远程：
```bash
# 在挂载点创建文件
echo "Hello from local" > ~/MyCloudDrive/local-test.txt
# 检查远程存储
cat /tmp/mock-remote-storage/local-test.txt
```
### 2. 测试强制同步
```bash
# 触发强制同步
curl -X POST http://localhost:9527/trigger-sync
```
### 3. 信号处理测试
发送停止信号，验证优雅关闭：
```bash
# 查找进程ID
ps aux | grep client-bridge
# 发送停止信号
kill -TERM <PID>
```

## License
MIT 