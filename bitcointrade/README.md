# Bitcointrade
Módulo responsável por obter dados da API do [Bitcointrade](https://apidocs.bitcointrade.com.br/#9fe41816-3d20-e53e-9273-643c95279dc4).


# TODO
- [x] Testes da lógica de obtenção de Trades (conversão de strings e erros)
- [x] Implementação de lógica para adicionar critérios à query String de URL de busca de Trades
- [x] - Verificar a necessidade de uso das funções utils para conversão de números para Strings (strconv já possui utilidades para isso)
- [ ] - Refatorar lógica de correção de Trade inválido para melhoria Semântica
- [ ] - Implementar persistência de dados já baixados do Bitcointrade
- [ ] - Adicionar tratamento para quando range de datas for inválido (início maior que fim)