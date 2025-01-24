
## PT-BR

# Go-reusable-api

Este README está disponível em vários idiomas:  
- [English](../README.md)  
- [Português](README.pt.md)  


## Visão Geral
Go-reusable-api é uma API altamente eficaz, reutilizável e fácil de entender, construída com Golang, Gin e Gorm. Esta API foi projetada para oferecer um conjunto de recursos robusto, mantendo um desempenho e segurança otimizados.

### Funcionalidades
- **Suporte a Cache**: Utiliza Redis para cache.
- **Base de Dados**: PostgreSQL.
- **Limitação de Taxa**: Limitador de taxa embutido para prevenir spam e quedas do serviço.
- **Suporte a Docker**: Inclui Dockerfile para containerização.
- **Monitoramento**: Integração com Prometheus para monitoramento de desempenho.
- **Autenticação**: Token Bearer e senhas hash para segurança.

## Primeiros Passos

### Pré-requisitos
- Golang
- Docker
- PostgreSQL
- Redis

### Executando com Docker
Execute os seguintes comandos para construir e executar o container Docker:

```bash
docker build -t go-reusable-api .
docker run -p 8080:8080 go-reusable-api
```

## Configuração

### `env.go`
Este arquivo deve conter uma função chamada `SetSalt` que retorna o salt em tipo byte.

```go
func SetSalt() []byte {
    seu_salt := "Hey"
    return []byte(seu_salt)
}
```

Também deve conter uma função chamada `envGo`, para que você possa definir as informações do seu banco de dados.

```go 
func envGo() {
    os.Setenv("database", "host=example-db server user=fakeuser password=fakepassword dbname=fakedb port=1234 sslmode=require")
}
```

## Desempenho
Para garantir o mais alto nível de desempenho, a API possui:
- Otimização de cache através de Redis.
- Índices de base de dados e consultas otimizadas em PostgreSQL.
- Base de código limpa e eficiente.

## Segurança
- Senhas são hash para armazenamento seguro.
- Token Bearer é usado para autorização do sistema.

## Monitoramento com Prometheus
A API inclui uma configuração Prometheus dentro do container Docker para monitoramento de desempenho em tempo real.

## Requisições de API

Exemplos de como fazer requisições à API podem ser encontrados na pasta `rest`. Isso deve orientá-lo sobre como fazer requisições adequadas à API.

### Parâmetros Obrigatórios
- `Email`
- `Senha`
- `Token Bearer`

Lembre-se de incluir seu e-mail, senha e token de portador apropriado ao fazer chamadas de API.

## Novas features:
- [ ] **Adicionar documentação Swagger com swaggo**
- [ ] **Incluir govalidator**
- [ ] **Load-Balancer**
- [ ] **Adicionar Feature Flags**
