# IsNumberEven

## Descrição
O pacote `IsNumberEven` fornece uma utilidade simples para verificar se um número dado é par.

## Instalação
Para usar este pacote em seu projeto Go, você pode instalá-lo executando:

```bash
go get github.com/tadrianonet/isnumbereven
```

## Uso
Importe o pacote e use a função IsEven para verificar se um número é par:

```bash
package main

import (
    "fmt"
    "github.com/tadrianonet/isnumbereven"
)

func main() {
    num := 4
    if isnumbereven.IsEven(num) {
        fmt.Println(num, "é par.")
    } else {
        fmt.Println(num, "é ímpar.")
    }
}
```

## Testes
Para rodar os testes e verificar se o pacote está funcionando corretamente, use:

```bash
go test
```

## Licença
Este projeto está licenciado sob a licença MIT. Para mais detalhes, consulte o arquivo LICENSE.


