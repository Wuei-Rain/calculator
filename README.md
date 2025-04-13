# 🧮 全栈计算器项目
一个现代化的全栈计算器应用，结合了Go的ConnectRPC与Next.js的App Router，实现高效的前后端交互，是一个前后端分离项目。

## ✨ 特性

- **高效通信**: 使用ConnectRPC替代传统REST
- **性能优化**: Next.js App Router实现按需加载
- **模块化设计**: 易于扩展新运算功能
- **全面测试**: 完善的单元测试覆盖

## 🛠️ 技术栈

- **前端**: Next.js (App Router), TypeScript
- **后端**: Go, ConnectRPC
- **协议**: Protocol Buffers

## 📂 项目结构

```
calculator/
├── client/                  # 前端Next.js项目
│   ├── app/                 # App Router核心目录
│   │   ├── page.tsx         # 主页面组件
│   │   └── layout.tsx       # 全局布局
│   ├── src/                 # 源代码
│   │   ├── components/      # 公共组件
│   │   └── utils/          # 工具函数
│   └── __tests__/           # 前端测试
│       └── page.test.tsx    # 页面测试
├── proto/                   # Protocol Buffers定义
│   └── calculator.proto     # gRPC服务接口定义
└── server/                  # 后端Go服务
    ├── calculatorv1/        # 生成的gRPC代码
    ├── calculatorv1connect/ # Connect适配器
    ├── main.go              # 服务入口
    └── server.go            # 服务实现
```

## 🚀 快速开始

### 前端
在client的目录下运行
```
npm install
```
如果出现 peer 依赖冲突，请尝试添加 --legacy-peer-deps 参数：
```
npm install --legacy-peer-deps
```
```
npm run dev
```
前端服务会默认启动在：```http://localhost:3000```


### 后端
在 server/ 目录下，确保 go.mod 文件中包含所有依赖（如 ConnectRPC、rs/cors 等）。如果未安装 rs/cors，请运行：
```
go get github.com/rs/cors
go mod tidy
go run .
```
服务会监听在端口 8080，并通过 CORS 中间件允许来自： ```http://localhost:3000 ```的请求

## 测试指南

### 前端测试

前端使用Jest和Testing Library编写了测试用例。测试文件位于client/tests/page.test.tsx目录下。

运行前端测试:
```
npm run test 或者 npx jest
```
### 后端测试

后端使用 Go 的 testing 包编写了单元测试（例如在 server/ 目录下的 _test.go 文件）。

运行后端测试:
```
go test -v 
```
## 数据交互说明

### ConnectRPC协议

本项目使用ConnectRPC实现前后端通信，基于gRPC协议的高性能RPC框架。

### 接口定义

接口定义文件位于proto/calculator.proto，包含以下服务定义：

```
service CalculatorService {
  rpc Calculate(CalculatorRequest) returns (CalculatorResponse);
}

message CalculatorRequest {
  double number1 = 1;
  double number2 = 2;
  string operator = 3;
}

message CalculatorResponse {
  double result = 1;
}




