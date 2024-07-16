# Go Serverless Template

Este repositório é um template para criar aplicações serverless utilizando Go. Ele inclui configurações essenciais e estruturas para facilitar o desenvolvimento de projetos serverless com AWS Lambda.

## Estrutura do Projeto

- `cmd/`: Contém os comandos principais da aplicação.
- `docker/`: Configurações para containerização com Docker.
- `serverless.yml`: Configuração do Serverless Framework.
- `Makefile`: Comandos para build e deploy.
- `go.mod` e `go.sum`: Gerenciamento de dependências Go.
- `codecove.sh`: Script para integração contínua.

## Pré-requisitos

- Go 1.18+
- Node.js e npm
- Serverless Framework

## Instalação

Clone o repositório:
```sh
git clone https://github.com/mateusmacedo/go-serverless-template.git
cd go-serverless-template
```

Instale as dependências:
```sh
npm install -G serverless
npm install
```

## Uso

### Desenvolvimento Local

Para rodar a aplicação localmente:
```sh
make up
make deploy
```

### Utilizando o template

Para chamar o endpoint http de hello apos o deploy a url do api gateway sera exibida no terminal, basta acessar a url e adicionar o path /hello/{name} para ver o retorno.

```sh
endpoint: http://localhost:4566/restapis/{randon-id}/local/_user_request_/hello/{name}functions:
  http-hello-broadcaster: go-serverless-template-local-http-hello-broadcaster (3.8 MB)
  hello: go-serverless-template-local-hello (2.9 MB)
  hello-secondary: go-serverless-template-local-hello-secondary (2.9 MB)
  health-check: go-serverless-template-local-health-check (2.9 MB)
Stack Outputs:
  ServerlessDeploymentBucketName: go-serverless-template-local-serverlessdeploymentbuck-01f8a99f
  HelloLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:000000000000:function:go-serverless-template-local-hello:2
  HelloDashsecondaryLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:000000000000:function:go-serverless-template-local-hello-secondary:2
  HealthDashcheckLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:000000000000:function:go-serverless-template-local-health-check:2
  HttpDashhelloDashbroadcasterLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:000000000000:function:go-serverless-template-local-http-hello-broadcaster:2
  ServiceEndpoint: https://{randon-id}.execute-api.localhost.localstack.cloud:4566/local
```

Endpoint de rest api do localstack:

```sh
curl -X GET http://localhost:4566/restapis/{randon-id}/local/_user_request_/hello/{name}
```

Endpoint do api gateway criado pelo localstack:

```sh
curl -X GET https://{randon-id}.execute-api.localhost.localstack.cloud:4566/local/hello/{name}
```

Resultado esperado:

```sh
{"message":"Hello, {name}"}
```

Nos logs do localstack sera exibido algo parecido com:

```sh
2024-07-16T07:32:05.596 DEBUG --- [2c92675bff18] l.s.l.i.version_manager    : [go-serverless-template-local-hello-d4907159-5266-413c-813b-2c92675bff18] 2024/07/16 07:32:05 Hello output: Hello mateus from primary
...
2024-07-16T07:32:05.642 DEBUG --- [d44ac9d1a5cd] l.s.l.i.version_manager    : [go-serverless-template-local-hello-secondary-a46ea4f4-e573-4165-b591-d44ac9d1a5cd] 2024/07/16 07:32:05 Hello output: Hello mateus from secondary
```

O localstack ira levantar um container para cada lambda criada, e ira exibir os logs de cada lambda no terminal.

hello-primary:

```sh
2024/07/16 07:32:05 Hello output: Hello {name} from primary
```

hello-secondary:

```sh
2024/07/16 07:32:05 Hello output: Hello {name} from secondary
```

## Contribuição

Contribuições são bem-vindas! Por favor, abra um issue ou envie um pull request.

## Licença

Este projeto está licenciado sob a licença MIT.

## Autor

- [Mateus Macedo](https://github.com/mateusmacedo)