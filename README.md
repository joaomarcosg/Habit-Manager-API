# API de Gerenciamento de hábitos - Habit Manager

## 👨‍💻 Tecnologias e Ferramentas

| Tecnologia | Descrição |
| ---------- | --------- |
| Go         | Linguagem de programação estaticamente tipada |
| Chi        | Framework Go que facilita a criação der servidores HTTP |
| Postgres   | Banco de dados relacional |
| Docker     | Plataforma de software para implantar aplicativos em containers |
| Gorilla CSRF        | CSRF Tokens, autenticação baseada em token e um cookie pair

## 📝 Descrição do projeto

A API de gerenciamento de hábitos (Habit Manager API) é uma aplicação para o controle e gerenciamento de hábitos pessoais. Com ela podemos criar e medir o progresso dos nossos hábitos.

Toda a API foi desenvolvida em **Go** e com auxílio do framework **Chi** para acelerar o desenvolvimento do projeto. Para a persistência dos dados usei **Postgres** em conjunto com **Docker** para rodar na aplicação. Em relação a autenticação para acesso aos recursos optei pelo **CSRF Tokens** com auxílio do pacote Gorilla/CSRF.

Nesse projeto tive a oportunidade de aplicar meus conhecimentos em APIs RESTful, tratamento de erros, persistência de dados usando ferramentas como SQLC, autenticação, tratamento de JSON.