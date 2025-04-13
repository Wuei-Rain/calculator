This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## 项目特色

- 基于Next.js框架构建的高性能计算器应用
- 使用React 19最新特性
- 采用TypeScript进行类型安全开发
- 集成ConnectRPC实现前后端通信
- 支持Jest单元测试

## 技术栈

- 前端: Next.js 15, React 19, TypeScript
- 后端: Go, ConnectRPC
- 测试: Jest, Testing Library

## 快速开始

首先，启动开发服务器:

```bash
npm run dev
```

打开 [http://localhost:3000](http://localhost:3000) 在浏览器中查看结果。

您可以通过修改 `app/page.tsx` 来开始编辑页面。文件保存后页面会自动更新。

## 测试指南

### 前端测试

前端使用Jest和Testing Library进行单元测试。测试文件位于`__tests__`目录下。

运行前端测试:
```bash
npm test
```

### 后端测试

后端使用Go标准测试包进行单元测试。测试文件以`_test.go`结尾。

运行后端测试:
```bash
cd server

go test -v ./...
```

## 数据交互说明

### ConnectRPC协议

本项目使用ConnectRPC实现前后端通信，基于gRPC协议的高性能RPC框架。

### 接口定义

接口定义文件位于`proto/calculator.proto`，包含以下服务定义：

```protobuf
service CalculatorService {
  rpc Calculate (CalculateRequest) returns (CalculateResponse) {}
}

message CalculateRequest {
  string expression = 1;
}

message CalculateResponse {
  double result = 1;
  string error = 2;
}
```

### 请求示例

前端请求示例(TypeScript):
```typescript
const response = await client.calculate({ expression: "1+2*3" });
```

后端响应示例(Go):
```go
return &calculatorv1.CalculateResponse{
  Result: 7,
}, nil
```

### 错误处理

- 表达式错误返回`error`字段
- HTTP状态码400表示请求格式错误
- HTTP状态码500表示服务器内部错误

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
