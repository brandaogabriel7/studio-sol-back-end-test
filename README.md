# studio-sol-back-end-test

## Introdução

O problema consiste em validar uma senha com base nas regras fornecidas na requisição, depois, retornar se a senha é válida e, se não for, listar quais regras aquela senha fere.

## Processo

Abaixo está descrito meu processo, desde a decisão das tecnologias e padrões, passando pelo raciocínio dos problemas até a solução final.

## Tecnologias e padrões

Eu decidi usar `golang` para resolver o desafio, visto que a linguagem está na stack da **Studio Sol** e que eu usei **.NET** no meu outro desafio e não conseguiram testar. Também decidi usar GraphQL porque vale mais pontos no processo.

Para facilitar a execução da aplicação e deploy para algum ambiente, resolvi criar uma `Dockerfile` para gerar uma imagem de container e um pipeline para garantir que tudo funciona, gerar a imagem e fazer o deploy automaticamente.

### Lista completa:

- golang
- gqlgen (para facilitar o uso do GraphQL em go)
- ginkgo e gomega (para testes)
- Docker
- Github e Github Actions
- Heroku e Cloudflare (para deploy da aplicação)

## Gerar projeto inicial

Meu primeiro passo foi criar o projeto usando o [gqlgen](https://github.com/99designs/gqlgen).

Depois que o projeto foi gerado, parti para a criação da `Dockerfile`.

Por último, com o projeto criado e a `Dockerfile` pronta, criei o pipeline com *Github Actions* para automatizar o deploy da aplicação.

Quando a aplicação já estava funcionando no Heroku e o pipeline estava configurado, eu adicionei um domínio customizado usando o **Cloudflare**.

As rotas ficaram:
- Playground: https://studio-sol-back-end-test.gabrielbrandao.net
- Endpoint GraphQL: https://studio-sol-back-end-test.gabrielbrandao.net/graphql

## Criar schema da aplicação

O próximo passo foi editar o arquivo `schema.graphqls` para configurar a query `verify` e seus tipos. Depois, o script `generate` da **gqlgen** atualizou os códigos com base no novo esquema graphql.

O schema ficou assim:
```gql
type Verify {
  verify: Boolean!
  noMatch: [String!]!
}

input Rule {
  rule: String!
  value: Int
}

type Query {
  verify(password: String!, rules: [Rule!]!): Verify!
}
```

## Configurar ambiente para testes

Antes de começar o desenvolvimento, eu decidi configurar o projeto para utilizar [ginkgo](https://github.com/onsi/ginkgo) e [gomega](https://github.com/onsi/gomega) para os testes e atualizar o pipeline para rodar os testes com a ginkgo CLI.