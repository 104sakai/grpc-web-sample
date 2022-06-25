<h1 align="center">gRPC Web Application Templates for Next.js + golang</h1>

## 開発環境

brewとnpmをつかってインストールする

DockerとDockerComposeをつかうので、あわせてインストールする
```
% npm list
web@0.1.0 .../github.com/104sakai/grpc-web-sample/web
├── @tailwindcss/forms@0.5.2
├── @types/node@18.0.0
├── @types/react-dom@18.0.5
├── @types/react@18.0.14
├── autoprefixer@10.4.7
├── eslint-config-next@12.1.6
├── eslint@8.18.0
├── next@12.1.6
├── postcss@8.4.14
├── react-dom@18.2.0
├── react@18.2.0
├── tailwindcss@3.1.4
└── typescript@4.7.4

% yarn -v
1.22.19

% go version
go version go1.18.1 darwin/amd64

# golang-migrate
% migrate -version
v4.15.2

# protobuf
% protoc --version
libprotoc 3.19.4

# grpc
grpc_cli --version
grpc_cli
```
