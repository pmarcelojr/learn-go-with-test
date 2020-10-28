# Olá, mundo

[**Você pode encontrar todos os códigos para esse capítulo aqui**](https://github.com/pmarcelojr/learn-go-with-test/tree/main/primeiros-passos/ola-mundo)

O primeiro programa foi um _Olá, mundo_.

### Disciplina

Vamos repassar o ciclo novamente:

* Escrever um teste
* Compilar o código sem erros
* Rodar o teste, ver o teste falhar e certificar que a mensagem de erro faz sentido
* Escrever a quantidade mínima de código para o teste passar
* Refatorar

Este ciclo pode parecer tedioso, mas se manter nesse ciclo de feedback é importante.

Ele não apenas garante que você tenha _testes relevantes_, como também ajuda a _projetar um bom software_ refatorando-o com a segurança dos testes.

Ver a falha no teste é uma verificação importante porque também permite que você veja como é a mensagem de erro. Para quem programa, pode ser muito difícil trabalhar com uma base de código que, quando há falha nos testes, não dá uma ideia clara de qual é o problema.

Assegurando que seus testes sejam rápidos e configurando suas ferramentas para que a execução de testes seja simples, você pode entrar em um estado de fluxo ao escrever seu código.

Ao não escrever testes, você está comprometendo-se a verificar manualmente seu código executando o software que interrompe seu estado de fluxo, o que não economiza tempo, especialmente a longo prazo.

## Resumindo

Quem imaginaria que você poderia tirar tanto proveito de um `Olá, mundo`?

Até agora você deve ter alguma compreensão de:

### Algumas das sintaxes utilizadas:

* Declarar funções, com argumentos e tipos de retorno
* `if`, `const` e `switch`
* Declarar variáveis e constantes
* Escrever testes

### _Por que_ as etapas são importantes no processo TDD 

* _Escreva um teste que falhe e veja-o falhar_, para que saibamos que escrevemos um teste _relevante_ para nossos requisitos e vimos que ele produz uma _descrição da falha fácil de entender_
* Escrever a menor quantidade de código para fazer o teste passar, para que saibamos que temos software funcionando
* _Em seguida_, refatorar, tendo a segurança de nossos testes para garantir que tenhamos um código bem feito e fácil de trabalhar

No nosso caso, passamos de `Ola()` para `Ola("nome")`, para `Ola ("nome"," Francês ")` em etapas pequenas e fáceis de entender.

Naturalmente, isso é trivial comparado ao software do "mundo real", mas os princípios ainda permanecem. O TDD é uma habilidade que precisa de prática para se desenvolver. No entanto, você será muito mais facilidade em escrever software sendo capaz de dividir os problemas em pedaços menores que possa testar.