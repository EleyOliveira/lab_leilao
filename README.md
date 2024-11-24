**Sistema de leilão**  

**Descrição:**  
Este projeto tem como objetivo a simulação de um leilão registrando os leilões e os lances em um banco de dados MongoDB utilizando a linguagem GO.  
O encerramento do leilão é realizado de forma automatizada após um determinado período de tempo informado em um arquivo de configuração. 

**Execução**

O MongoDB e a aplicação serão inicializados com o comando docker-compose up -d.

**Acesso as funcionalidades**

* Incluir um leilão: utilize o arquivo auctions.http na pasta internal/infra/api/web, que possui um exemplo de como realizar uma requisição POST para a url http://localhost:8080/auction.
* Consultar um leilão: acesse a url http://localhost:8080/auction/(Id do leilão) e informe o Id do leilão no formato UUID.  
* Consultar todos os leilões: acesse a url http://localhost:8080/auction?status=0, informando o valor 0 no parâmetro status.
* Consultar os leilões encerrados: acesse a url http://localhost:8080/auction?status=1, informando o valor 1 no parâmetro status.
* Incluir um lance: utilize o arquivo auctions.http na pasta internal/infra/api/web, que possui um exemplo de como realizar uma requisição POST para a url http://localhost:8080/bid.
* Consultar os lances de um leilão: acesse a url http://localhost:8080/bid/(Id do leilão) e informe o Id do leilão no formato UUID.
* Consultar o lance vencedor de um leilão: acesse a url http://localhost:8080/auction/winner/(Id do leilão) e informe o Id do leilão no formato UUID.


