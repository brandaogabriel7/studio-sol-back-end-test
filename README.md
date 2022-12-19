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
  value: Int!
}

type Query {
  verify(password: String!, rules: [Rule!]!): Verify!
}
```

## Configurar ambiente para testes

Eu resolvi utilizar o [ginkgo](https://github.com/onsi/ginkgo) e o [gomega](https://github.com/onsi/gomega) porque eu gosto da estrutura que eles fornecem para escrita de testes, acho que os testes ficam mais legíveis e organizados, além de mais fáceis de escrever. E, como eu desenvolvo com TDD, usar essas bibliotecas me traz mais produtividade para escrever muitos testes.

> Essas bibliotecas possuem uma sintaxe baseada no BDD, o que vai faicilitar muito o entendimento do que o código faz pela leitura dos testes.

Então, antes de começar o desenvolvimento, eu configurei o projeto e o pipeline para utilizarem a **ginkgo CLI**.

## Implementação das validações

### Testes de integração como base

Meu primeiro passo foi escrever testes de integração que testassem alguns casos da query `verify`, inclusive o caso fornecido na descrição da prova. Depois que eu tivesse os testes de integração falhando, eu seguiria o ciclo do TDD com testes de unidade até que a feature estivesse completa e os testes de integração também passassem.

Escrevi testes de integração parametrizados com as mesmas regras do exemplo da prova (minSize, minSpecialChars, noRepeted, minDigit). Coloquei alguns casos de testes apenas com essas regras para começar. Depois que elas estivessem implementadas eu acrescentaria mais casos de testes de integração para cobrir as outras regras.

### Arquitetura

Uma das minhas primeiras ideias para implementação das validações era usar o padrão [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility). No entanto, também tive a ideia de utilizar o [Strategy Pattern](https://refactoring.guru/design-patterns/strategy) e chamar cada strategy na ordem que aparecesse na coleção de regras da requisição.

A segunda opção me pareceu mais simples, além de me permitir retornar as regras que não fossem cumpridas na mesma ordem que elas viesse na requisição.

### MinSizeValidationStrategy

Tendo decidido que usaria o ***Strategy Pattern***, comecei a implementar os strategies, um por um, sempre seguindo o ciclo do TDD.

Iniciei pela validação que me pareceu mais simples, a `minSize`.

Essa validação consiste em retornar **inválido** para senhas *menores* que o valor fornecido para o `minSize` e retornar **válido** para senhas *maiores* ou de tamanho *igual* ao valor fornecido para `minSize`.

### MinDigitValidationStrategy

As validações de `minDigit` e `minSpecialChars` são relativamente simples, também. As duas podem ser resolvidas facilmente usando **Regex**. Comecei pela de dígitos porque a expressão regular é mais simples.

A validação consiste em encontrar todos as ocorrências de um dígito numérico dentro da senha, retornar *inválido* se o número de ocorrências for *menor* que o valor `minDigit` ou retornar *válido* se o número de ocorrências for *maior* ou igual ao valor `minDigit`.

A expressão regular é: `\d`

### MinSpecialChars

A lógica do `minSpecialChars` é a mesma do minDigit, mas a expressão regular é: `[!@#$%^&*()\-+\\\/{}\[\]]`

### NoRepeted

A regra `noRepeted` foge mais da lógica das outras regras. A solução que eu pensei foi a seguinte, comprimir todos os caracteres repetidos consecutivos da senha em um só e comparar o tamanho da senha comprimida com a senha original.

Se a senha comprimida e a senha original forem iguais, a regra passou. Se elas forem diferentes, a regra não passou.

Exemplos:
- Sucesso: A senha *"abacate123"*, depois de comprimida, continua *"abacate123"*.

- Falha: A senha *"Opaaa73"*, depois de comprimida, vira *"Opa73"* (diferente da original).

### PasswordValidationService

O `PasswordValidationService` vai ser o serviço responsável por chamar as strategies em ordem e retornar a validação completa para o `resolver` da query **verify**.

Ele recebe um map que atrela os nomes das regras de validação às suas respectivas *estratégias*.

### Resolver implementado com algumas regras

Com o serviço de validação implementado, eu fiz a injeção no resolver e coloquei tudo para funcionar.

Nesse ponto, todos os testes estavam passando (de integração e de unidade) e eu testei alguns casos manualmente pelo playground da aplicação e tudo funcionou.

Até o momento apenas as regras **minSize**, **minDigit**, **minSpecialChars** e **noRepeted** haviam sido implementadas, mas a estrutura já estava preparada para receber as regras restantes facilmente.
