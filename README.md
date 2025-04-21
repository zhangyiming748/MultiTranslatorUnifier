# MultiTranslatorUnifier

一个强大的多翻译器统一服务，支持多个翻译源（Google、Bing、DeepLX），具有智能缓存和故障转移机制。

## 功能特性

- **多翻译源支持**
  - Google Translate
  - Microsoft Bing Translator
  - DeepLX API

- **智能缓存系统**
  - 自动缓存翻译结果
  - 快速检索历史翻译
  - 减少重复请求

- **高可用性设计**
  - 多翻译源并行请求
  - 自动故障转移
  - 超时自动重试机制

- **跨平台支持**
  - Linux系统支持Google和Bing翻译
  - 全平台支持DeepLX API

## 环境要求

- Go 1.24
- MySQL数据库（用于存储翻译历史）
- Linux环境需要安装`translate-shell`

## 快速开始

### 安装

```bash
# 克隆项目
git clone https://github.com/zhangyiming748/MultiTranslatorUnifier.git

# 进入项目目录
cd MultiTranslatorUnifier

# 安装依赖
go mod download
```

### 配置

1. **环境变量设置**

```bash
# Google翻译代理设置（可选）
export PROXY="your-proxy-address"

# DeepLX API密钥
export LINUXDO="your-api-key"
```

2. **数据库配置**

确保MySQL服务已启动，并正确配置了数据库连接信息。

### 使用方法

```go
import "github.com/zhangyiming748/MultiTranslatorUnifier/logic"

func main() {
    source := "Hello World"
    from, result := logic.Trans(source)
    fmt.Printf("Translation source: %s\nResult: %s\n", from, result)
}
```

## 技术架构

### 核心组件

- **Trans模块**: 统一翻译入口，负责调度和管理翻译流程
- **Storage模块**: 处理翻译历史的存储和检索
- **Translate-Shell模块**: 封装Google和Bing翻译接口
- **LinuxDo模块**: 集成DeepLX API服务

### 工作流程

1. 接收翻译请求
2. 检查翻译缓存
3. 并行请求可用翻译源
4. 采用最快返回的有效结果
5. 缓存新的翻译结果

## 性能优化

- 使用goroutine实现并行翻译请求
- sync.Once确保只采用最快的翻译结果
- 智能缓存机制避免重复翻译
- 30秒超时保护，自动重试机制

## 错误处理

- 翻译源故障自动切换
- 超时自动重试
- 详细的错误日志记录

## Docker支持

项目提供了Docker支持，可以通过以下命令快速部署：

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d
```

## 许可证

本项目基于MIT许可证开源。

## 贡献指南

欢迎提交Issue和Pull Request来帮助改进项目。在提交代码前，请确保：

1. 代码符合Go的编码规范
2. 添加了必要的测试用例
3. 所有测试用例通过
4. 更新了相关文档