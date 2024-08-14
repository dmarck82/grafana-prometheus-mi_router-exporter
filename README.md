# Visualizador de Dados Home Server

Este mini projeto visa obter métricas de um home server e do roteador Xiaomi AX3000. Os dados podem ser visualizados através do Grafana utilizando dashboards customizados.

Os dashboards a serem utilizados são:

1. **Node Exporter Full**: ID-1860
2. **Xiaomi-Router**: ID-16557

O projeto inclui um arquivo `grafana.json` com a versão traduzida e ajustada do dashboard "Xiaomi-Router". Basta copiar o JSON para a ferramenta de importação do Grafana.

## Como Executar o Projeto

### 1. Baixe os Arquivos

Baixe os arquivos do projeto mantendo a estrutura de pastas fornecida.

### 2. Construir a Imagem do `miwifi-exporter`

1. Navegue para o diretório do `miwifi_exporter` onde está localizado o `Dockerfile`:

    ```bash
    cd /c/Repositorios/grafana-prometheus/miwifi-exporter
    ```

2. Construa a imagem Docker para o `miwifi-exporter`:

    ```bash
    docker build -t miwifi-exporter:latest .
    ```

    - `docker build` é o comando para criar uma imagem Docker.
    - `-t miwifi-exporter:latest` define a tag da imagem como `miwifi-exporter:latest`.
    - `.` indica que o Dockerfile está no diretório atual.

### 3. Configurar e Executar os Serviços com Docker Compose

1. Volte para o diretório raiz onde está localizado o `docker-compose.yml`:

    ```bash
    cd /c/Repositorios/grafana-prometheus
    ```

2. Abra o arquivo `docker-compose.yml` em um editor de texto e substitua os seguintes valores com os dados do seu roteador:

    ```yaml
    environment:
      - ROUTER_IP=XXX
      - ROUTER_PASSWORD=XXX
    ```

3. Execute o Docker Compose para construir e iniciar os serviços:

    ```bash
    docker-compose up --build
    ```

    - `docker-compose up --build` irá construir as imagens (se necessário) e iniciar os serviços definidos no `docker-compose.yml`.

### 4. Importar Dashboards no Grafana

1. Acesse o Grafana no seu navegador, geralmente em `http://localhost:3000` (ou a URL configurada).
2. Faça login com suas credenciais.
3. Vá para **Dashboards** > **Import** e cole o JSON do arquivo `grafana.json` que está incluído no projeto.

Isso deve configurar e iniciar todos os serviços necessários, permitindo que você visualize as métricas do seu home server e do roteador Xiaomi AX3000 no Grafana.
