# studio-sol-back-end-test

## Sumário
- [Introdução](#introdução)
- [Processo](#processo)
- [Tecnologias e padrões](#tecnologias-e-padrões)
    - [Lista completa](#lista-completa)
- [Gerar projeto inicial](#gerar-projeto-inicial)
- [Criar schema da aplicação](#criar-schema-da-aplicação)
- [Configurar projeto para testes](#configurar-projeto-para-testes)
- [Implementação das validações](#implementação-das-validações)
    - [Testes de integração como base](#testes-de-integração-como-base)
    - [Arquitetura](#arquitetura)
    - [MinSizeValidationStrategy](#minsizevalidationstrategy)
    - [MinDigitValidationStrategy](#mindigitvalidationstrategy)
    - [MinSpecialCharsValidationStrategy](#minspecialcharsvalidationstrategy)
    - [NoRepetedValidationStrategy](#norepetedvalidationstrategy)
    - [PasswordValidationService](#passwordvalidationservice)
    - [Injetar serviço no *resolver*](#injetar-serviço-no-resolver)
    - [Refatoração das regras com Regex](#refatoração-das-regras-com-regex)
    - [MinUppercaseStrategy e MinLowercaseStrategy](#minuppercasestrategy-e-minlowercasestrategy)
- [Como testar a aplicação](#como-testar-a-aplicação)
    - [Localmente](#localmente)
    - [Dockerfile](#dockerfile)
    - [Versão hospedada](#versão-hospedada)
- [Conclusão](#conclusão)
- [Alguns casos de teste](#alguns-casos-de-teste)

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

Por último, com o projeto criado e a [Dockerfile pronta](Dockerfile), criei o pipeline com *Github Actions* para automatizar o deploy da aplicação.

Quando a aplicação já estava funcionando no Heroku e o [pipeline estava configurado](.github/workflows/pipeline.yaml), eu adicionei um domínio customizado usando o **Cloudflare**.

As rotas ficaram:
- Playground: https://studio-sol-back-end-test.gabrielbrandao.net
- Endpoint GraphQL: https://studio-sol-back-end-test.gabrielbrandao.net/graphql

## Criar schema da aplicação

O próximo passo foi editar o arquivo `schema.graphqls` para configurar a query `verify` e seus tipos. Depois, o script `generate` da **gqlgen** atualizou os códigos com base no novo esquema graphql.

Veja o schema [aqui](graph/schema.graphqls).

## Configurar projeto para testes

Eu resolvi utilizar o [ginkgo](https://github.com/onsi/ginkgo) e o [gomega](https://github.com/onsi/gomega) porque eu gosto da estrutura que eles fornecem para escrita de testes, acho que os testes ficam mais legíveis e organizados, além de mais fáceis de escrever. E, como eu desenvolvo com TDD, usar essas bibliotecas me traz mais produtividade para escrever muitos testes.

> Essas bibliotecas possuem uma sintaxe baseada no BDD, o que vai faicilitar muito o entendimento do que o código faz pela leitura dos testes.

Então, antes de começar o desenvolvimento, eu configurei o projeto e o pipeline para utilizarem a **ginkgo CLI**.

## Implementação das validações

### Testes de integração como base

Meu primeiro passo foi escrever testes de integração que testassem alguns casos da query `verify`, inclusive o caso fornecido na descrição da prova. Depois que eu tivesse os testes de integração falhando, eu seguiria o ciclo do TDD com testes de unidade até que a feature estivesse completa e os testes de integração também passassem.

Escrevi [testes de integração](src/integration_tests/verify_test.go) parametrizados com as mesmas regras do exemplo da prova (minSize, minSpecialChars, noRepeted, minDigit). Coloquei alguns casos de testes apenas com essas regras para começar. Depois que elas estivessem implementadas eu acrescentaria mais casos de testes de integração para cobrir as outras regras.

### Arquitetura

Uma das minhas primeiras ideias para implementação das validações era usar o padrão [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility). No entanto, também tive a ideia de utilizar o [Strategy Pattern](https://refactoring.guru/design-patterns/strategy) e chamar cada strategy na ordem que aparecesse na coleção de regras da requisição.

A segunda opção me pareceu mais simples, além de me permitir retornar as regras que não fossem cumpridas na mesma ordem que elas viesse na requisição.

### MinSizeValidationStrategy

Tendo decidido que usaria o ***Strategy Pattern***, comecei a implementar os strategies, um por um, sempre seguindo o ciclo do TDD.

Iniciei pela validação que me pareceu mais simples, a `minSize`.

Essa validação consiste em retornar **inválido** para senhas *menores* que o valor fornecido para o `minSize` e retornar **válido** para senhas *maiores* ou de tamanho *igual* ao valor fornecido para `minSize`.

> [MinSizeValidationStrategy](src/strategies/validation/min_size.go)

### MinDigitValidationStrategy

As validações de `minDigit` e `minSpecialChars` são relativamente simples, também. As duas podem ser resolvidas facilmente usando **Regex**. Comecei pela de dígitos porque a expressão regular é mais simples.

A validação consiste em encontrar todos as ocorrências de um dígito numérico dentro da senha, retornar *inválido* se o número de ocorrências for *menor* que o valor `minDigit` ou retornar *válido* se o número de ocorrências for *maior* ou igual ao valor `minDigit`.

A expressão regular é: `\d`

> [MinDigitValidationStrategy](src/strategies/validation/min_digit.go)

### MinSpecialCharsValidationStrategy

A lógica do `minSpecialChars` é a mesma do minDigit, mas a expressão regular é: `[!@#$%^&*()\-+\\\/{}\[\]]`

> [MinSpecialCharsValidationStrategy](src/strategies/validation/min_special_chars.go)

### NoRepetedValidationStrategy

A regra `noRepeted` foge mais da lógica das outras regras. A solução que eu pensei foi a seguinte, comprimir todos os caracteres repetidos consecutivos da senha em um só e comparar o tamanho da senha comprimida com a senha original.

Se a senha comprimida e a senha original forem iguais, a regra passou. Se elas forem diferentes, a regra não passou.

Exemplos:
- Sucesso: A senha *"abacate123"*, depois de comprimida, continua *"abacate123"*.

- Falha: A senha *"Opaaa73"*, depois de comprimida, vira *"Opa73"* (diferente da original).

> [NoRepetedValidationStrategy](src/strategies/validation/no_repeted.go)

### PasswordValidationService

O `PasswordValidationService` vai ser o serviço responsável por chamar as strategies em ordem e retornar a validação completa para o `resolver` da query **verify**.

Ele recebe um map que atrela os nomes das regras de validação às suas respectivas *estratégias*.

> [PasswordValidationService](src/services/password_validation/password_validation_service.go)

### Injetar serviço no *resolver*

Com o serviço de validação implementado, eu criei uma [factory](src/factories/password_validation_service_factory.go) para a implementação padrão, fiz a injeção no [resolver](graph/resolver.go) e chamei o serviço no [resolver da query verify](graph/schema.resolvers.go).

Nesse ponto, todos os testes estavam passando (de integração e de unidade). Eu testei alguns casos manualmente pelo playground da aplicação e tudo funcionou.

Até o momento apenas as regras **minSize**, **minDigit**, **minSpecialChars** e **noRepeted** haviam sido implementadas, mas a estrutura já estava preparada para receber as regras restantes facilmente.

### Refatoração das regras com Regex

Olhando as regras que ainda não estavam implementadas e comparando com as lógicas de validação de `minDigit` e `minSpecialChars`, eu percebi que elas teriam certo nível de duplicação e uma refatoração podia ser feita para todas as validações que usassem Regex.

Então, antes de escrever mais testes e implementar as regras novas, eu resolvi refatorar as estratégias de Regex existentes para usar o mesmo código, mudando apenas a expressão regular.

Criei uma [struct base](src/strategies/validation/regex_validation.go) com a lógica de validação com base em uma expressão. Depois atualizei as estratégias `minSpecialChars` e `minDigit` para utilizar a implementação, cada uma passando sua própria expressão regular de validação.

### MinUppercaseStrategy e MinLowercaseStrategy

As duas estratégias restantes, *minUppercase* e *minLowerCase* se aproveitam da struct base de validação regex, porém cada uma com sua expressão:
- minUppercase: `[A-Z]`
- minLowercase: `[a-z]`

Então, escrevi mais testes de integração que incluíssem essas regras, depois escrevi testes de unidade para implementar cada estratégia e finalizar as implementações das regras.

- [MinUppercaseStrategy](src/strategies/validation/min_uppercase.go)
- [MinLowercaseStrategy](src/strategies/validation/min_lowercase.go)

## Conclusão

Nesse ponto, todas as regras estavam implementadas e todos os testes automatizados e manuais passando.

Nas seções seguintes, você encontra as instruções para rodar a aplicação, assim como alguns casos de teste de exemplo.

## Como testar a aplicação

### Localmente

Para executar localmente você precisa ter `go 1.19` instalado.

```bash
# Navegue até pasta raiz do projeto
cd <pasta-onde-está-o-projeto>/studio-sol-back-end-test

# Execute a aplicação
go run cmd/server/server.go
```

A aplicação decide a porta pela variável de ambiente `PORT`. Caso nenhuma seja fornecida, a porta padrão é a 8080. As rotas são as seguintes:

Endpoint graphql: http://localhost:8080/graphql

Playground GraphQL: http://localhost:8080

> Lembre-se de trocar a porta se tiver fornecido um valor para a variável de ambiente `PORT`.

### Dockerfile

Caso não tenha `go 1.19` instalado, pode ser mais simples utilizar um container para testar.

```bash
# Navegue até pasta raiz do projeto
cd <pasta-onde-está-o-projeto>/studio-sol-back-end-test

# Faça o build da imagem
docker build -t studio-sol-back-end-test .

# Rode o container
docker run -d -p 8080:8080 --name studio-sol-back-end-test studio-sol-back-end-test
```

As rotas são as seguintes:

Endpoint graphql: http://localhost:8080/graphql

Playground GraphQL: http://localhost:8080

Você também pode passar outra porta ser usada pela aplicação:

```bash
docker run -d -e PORT=8000 -p 8080:8000 --name studio-sol-back-end-test studio-sol-back-end-test
```

### Versão hospedada

A aplicação também está hospedada, então você pode testá-la nesses links:

Endpoint graphql: https://studio-sol-back-end-test.gabrielbrandao.net/graphql

Playground GraphQL: https://studio-sol-back-end-test.gabrielbrandao.net

## Alguns casos de teste

Caso 1:

  - Entrada:

    ```gql
    {
      verify(
        password: "ee123&"
        rules: [{rule: "minSize", value: 8}, {rule: "minSpecialChars", value: 2}, {rule: "noRepeted", value: 0}, {rule: "minDigit", value: 4}, {rule: "minUppercase", value: 7}]
      ) {
        verify
        noMatch
      }
    }
    ```

  - Saída:

    ```json
    {
      "data": {
        "verify": {
          "verify": false,
          "noMatch": [
            "minSize",
            "minSpecialChars",
            "noRepeted",
            "minDigit",
            "minUppercase"
          ]
        }
      }
    }
    ```

Caso 2:

  - Entrada:

    ```gql
    query {
      verify(password: "TesteSenhaForte!123&", rules: [
        {rule: "minSize",value: 8},
        {rule: "minSpecialChars",value: 2},
        {rule: "noRepeted",value: 0},
        {rule: "minDigit",value: 4}
      ]) {
      verify
      noMatch
      }
    }
    ```

  - Saída:

    ```json
    {
      "data": {
        "verify": {
          "verify": false,
          "noMatch": [
            "minDigit"
          ]
        }
      }
    }
    ```

Caso 3:

  - Entrada:

    ```gql
    query {
      verify(password: "M!nhaS3nh@", rules: [
        {rule: "minSize",value: 8},
        {rule: "minSpecialChars",value: 2},
        {rule: "minDigit",value: 1}
      ]) {
      verify
      noMatch
      }
    }
    ```

  - Saída:

    ```json
    {
      "data": {
        "verify": {
          "verify": true,
          "noMatch": []
        }
      }
    }
    ```