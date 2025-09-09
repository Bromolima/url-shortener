# Encurtador de URL em Go

Uma API simples que permite ao usuário converter URLs longas em links curtos.

- Encurta URLs em links curtos.
- Redireciona do link curto para a URL original.

## Como usar

### Pré-requisitos

- Go 1.24.6 ou superior
- Docker-compose
- Make (opcional, mas recomendado)

### Comandos

```bash
# Rodar os testes unitários
make test

# Iniciar o docker
make docker-up

# Criar as tabelas no banco de dados
make migrate

# Excecuta a API
make run
```

## Estrutura do Projeto

```
url-shortener/
├── config/             # Configurações gerais da aplicação
│   ├── database/       # Conexão e inicialização do banco
│   │   ├── migration/  # Scripts de migração
│   │   └── postgres.go # Configuração do PostgreSQL
│   └── infra/          # Configurações adicionais de infraestrutura
│
├── internal/           # Código principal da aplicação
│   ├── http/           # Camada HTTP
│   │   ├── handler/    # Controladores que recebem as requisições
│   │   ├── responses/  # Estruturas e padrões de resposta
│   │   └── routes/     # Definição das rotas da API
│   │
│   ├── injector/       # Injeção de dependências
│   ├── mocks/          # Mocks para testes
│   ├── model/          # Entidades
│   ├── repository/     # Persistência
│   ├── service/        # Regras de negócio
│   └── utils/          # Funções utilitárias
│
├── go.mod              # Gerenciamento de dependências
├── go.sum              # Hashes das dependências
├── main.go             # Ponto de entrada da aplicação
├── Makefile            # Comandos para excecução
```