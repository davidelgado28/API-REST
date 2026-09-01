# API-REST

API REST simples em Go para gerenciamento de tarefas, compilada e publicada automaticamente no GitHub Container Registry (GHCR) via GitHub Actions.

**Como Executar**

Não é necessário instalar o Go ou compilar o código localmente. Com o Docker instalado, rode o comando no terminal:

```bash
docker run -p 8080:8080 ghcr.io/davidelgado28/task-api:latest
```

Com isso, a aplicação estará disponível em: **http://localhost:8080**
