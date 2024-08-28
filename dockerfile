# Utiliser une image Go basée sur Alpine Linux pour avoir une image légère
FROM golang:alpine

# Définir les métadonnées de l'image
LABEL name = "PENC MI"
LABEL maintainer=["bcoulibal, ssock, mamdrame, nmbengue"]
LABEL description="This project consists in creating a web forum"
LABEL version="1.0"

# Créer le répertoire de travail à l'intérieur de l'image
WORKDIR /app

# Copier tous les fichiers du répertoire source dans le répertoire de destination
COPY . /app

# Installer Bash (si nécessaire)
RUN apk add --no-cache bash

# Installez les dépendances nécessaires pour la compilation de Go
RUN apk update && apk add --no-cache gcc musl-dev sqlite-dev

# Compiler l'application Go pour créer un exécutable
RUN go mod download
RUN go build -o app .

EXPOSE 8000

# Commande pour démarrer l'application lorsque le conteneur est lancé
CMD ["./app"]
