#

#  Bank Account Api

> Project to simulate bank account api

---

## API Design

É adotado o padrão de design de APIs REST da google. Para mais informações como este padrão funciona leia a seguinte [doc](https://cloud.google.com/apis/design) do google.

## Configuracoes do Projeto

---

## Como Rodar

- First run

Antes de tudo, rode este comando para evitar problemas durante a instalação das dependências da aplicação:
```
git config --global url."git@github.com:your-user".insteadOf "https://github.com/your-user"
go env -w GOPRIVATE="github.com/your-user/*"
```

Então, você pode executar as dependências rodando:

```
make setup/local
```

- Run API

```
make run/api
```

- Run Container

```
make run/container
```
- Unit Test
```
make test/local/unit
```