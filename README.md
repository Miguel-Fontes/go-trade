# Go Trade
Breakable toy utilizado para aprendizado de Go. O intuito da aplicação é obter dados da API pública do [Bitcointrade](https://apidocs.bitcointrade.com.br/) e exibi-la de maneira gráfica.

## Build
Execute o `go install` do pacote:

    go install github.com/miguel-fontes/trade

## Execução
Após executar o install e ter o binário gerado, é necessário posicionar o binário + o pacote viewer (resources/viewer) em um mesmo diretório para executar a aplicação. Os arquivos contidos no diretório viewer são os HTMLs e JS necessários para gerar o gráfico em uma página web. Para ilustrar, construa um diretório da seguinte forma:

    ├── trade
    └── viewer
        ├── assets
        │   ├── charts.js
        │   └── main.css
        └── index.html

Feito isto, neste diretório execute:

    ./trade <data inicial: yyyy-mm-dd> <data final: yyyy-mm-dd>
    ./trade 2018-01-01 2018-01-28

Esta instrução executará a aplicação, fazendo-a obter todos os trades entre os dias 2018-01-01 e 2018-01-28. Ao fim do procedimento, a aplicação exibirá "Listening at 8080". Acesse localhost:8080 para visualizar o gráfico gerado.
    