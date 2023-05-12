### Programa de Extração de Arquivos

Este é um programa escrito em Go para extrair arquivos de arquivos .zip e .rar. O script procurará por arquivos .zip e .rar em um diretório fornecido como argumento e extrairá os arquivos dentro deles para um diretório chamado processado.


### Pré-requisitos
Go 1.14 ou superior
Biblioteca github.com/mholt/archiver para Go. Você pode instalá-la com o seguinte comando:

- go get github.com/mholt/archiver/v3


### Compilação
Para compilar o script, você pode usar o seguinte comando na linha de comando:

- go build main.go


### Execução
Para executar o script, você precisa fornecer o nome do diretório que contém os arquivos .zip e .rar como argumento. Por exemplo:

- ./main /home/user/Downloads


### Exemplos de uso
Suponha que você tenha uma pasta chamada "arquivos" na sua área de trabalho que contém alguns arquivos .zip e .rar que você deseja extrair.
Você pode executar o script da seguinte maneira:
- ./main /home/user/Desktop/arquivos


### Verificando os arquivos extraídos
Após a execução do script, você deve ver um diretório chamado "processado" no mesmo diretório onde o script foi executado. Este diretório contém três subdiretórios: "xml", "txt" e "outros". Os arquivos .xml e .txt extraídos estão nos respectivos diretórios, enquanto todos os outros tipos de arquivos estão no diretório "outros".
