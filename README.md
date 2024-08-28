# Forum

Ce projet consiste à créer un forum web.

## Table des matières

- [Aperçu du Projet](#aperçu-du-projet)
- [Structure du Projet](#structure-du-projet)
- [Pour Commencer](#pour-commencer)

## Aperçu du Projet

**Aperçu du Projet**

L'objectif de ce projet est de créer un forum web en utilisant GoLang avec les fonctionnalités clés suivantes :

1. **Communication entre les Utilisateurs** : Les utilisateurs peuvent interagir entre eux en créant des publications et des commentaires. Seuls les utilisateurs enregistrés ont le privilège de créer des publications et des commentaires, tandis que les utilisateurs non enregistrés peuvent les consulter.

2. **Association de Catégories** : Les utilisateurs enregistrés peuvent associer une ou plusieurs catégories à leurs publications. L'implémentation et la sélection des catégories sont laissées à la discrétion du développeur.

3. **J'aime et Je n'aime pas les Publications et les Commentaires** : Les utilisateurs enregistrés ont la possibilité d'exprimer leur approbation ou leur désapprobation en aimant ou en n'aimant pas les publications et les commentaires. Le nombre de "j'aime" et de "je n'aime pas" doit être visible pour tous les utilisateurs, quel que soit leur statut d'inscription.

4. **Filtrage des Publications** : Un mécanisme de filtrage est mis en place, permettant aux utilisateurs de filtrer les publications affichées en fonction de trois critères :
   - Catégories : Les utilisateurs peuvent filtrer les publications par catégories, créant ainsi des sous-forums pour des sujets spécifiques.
   - Publications Créées : Les utilisateurs enregistrés peuvent filtrer les publications qu'ils ont créées.
   - Publications Aimées : Les utilisateurs enregistrés peuvent filtrer les publications qu'ils ont aimées.

5. **Base de données SQLite** : Le projet utilise la base de données SQLite pour stocker des données telles que les informations d'identification des utilisateurs, les publications, les commentaires, les "j'aime" et les "je n'aime pas". La base de données est gérée à l'aide de requêtes SQL, notamment des opérations SELECT, CREATE et INSERT.

6. **Authentification** : Les utilisateurs peuvent s'inscrire en tant que nouveaux utilisateurs sur le forum en fournissant leur adresse e-mail, leur nom d'utilisateur et leur mot de passe. Le système garantit que chaque adresse e-mail est unique, et les mots de passe peuvent être cryptés pour une sécurité renforcée. De plus, les sessions de connexion sont gérées à l'aide de cookies, et chaque session inclut une date d'expiration.

En résumé, ce projet vise à créer un forum web où les utilisateurs enregistrés peuvent participer aux discussions en créant des publications et des commentaires, en associant des catégories à leurs publications et en exprimant leurs opinions à travers les "j'aime" et les "je n'aime pas". Un système d'authentification assure une inscription et une connexion sécurisées des utilisateurs, tandis qu'un mécanisme de filtrage permet aux utilisateurs de personnaliser leur expérience du forum en fonction des catégories, des publications créées et des publications aimées. L'ensemble du système repose sur une base de données SQLite pour le stockage et la récupération des données.

## Structure du Projet

Ce projet suit un modèle d'architecture en couches pour maintenir l'organisation du code et la séparation des préoccupations :

```
.
├── config
├── database
│   ├── query
│   └── sql
├── encryption
├── models
├── README.md
├── server
│   ├── handler
│   ├── repository
│   ├── routes
│   └── service
├── static
│   ├── assets
│   ├── scripts
│   └── styles
├── storage
├── templates
└── utils
```

1. **config/** : Gère la configuration et lit les fichiers .env personnalisés. voir [config](./docs/config.md).

2. **database/** : Gère les opérations de base de données telles que l'insertion, la mise à jour, etc. voir [database](./docs/database.md)

    - **database/query** : Génère des requêtes en fonction des données fournies. voir [query](./docs/database-query.md)
    
3. **models/** : Contient des modèles de données ou des structures représentant des objets dans l'application.

4. **server/** : Le cœur de l'application web, responsable de la gestion des requêtes HTTP et des réponses.

    - **server/handler/** : Définit les gestionnaires de requêtes, chacun traitant des routes ou des points d'extrémité spécifiques.
    
    - **server/repository/** : Interagit avec la base de données, faisant le lien entre le serveur et la base de données.
    
    - **server/routes/** : Définit la logique de routage, faisant correspondre les points d'extrémité HTTP aux gestionnaires.
    
    - **server/service/** : Contient la logique métier ou les services utilisés par le serveur et les gestionnaires.

## Pour Commencer

Pour commencer avec le projet web de forum GoLang, suivez ces étapes :

### Prérequis

Avant de commencer, assurez-vous d'avoir les prérequis suivants installés sur votre environnement de développement :

1. GoLang : Installez GoLang sur votre système en suivant les instructions officielles d'installation : [https://golang.org/doc/install](https://golang.org/doc/install)

2. SQLite : Assurez-vous d'avoir SQLite installé ou disponible sur votre système. Vous pouvez télécharger SQLite depuis le site officiel : [https://www.sqlite.org/download.html](https://www.sqlite.org/download.html)

### Installation

Suivez ces étapes pour configurer et exécuter le projet web de forum :

1. Clonez le Projet :

   ```shell
   git clone https://learn.zone01dakar.sn/git/bcoulibal/forum.git
   cd forum/
   ```

2. Installez les Dépendances :

   ```shell
   go mod download
   ```

3. Configurez les Variables d'Environnement :

   Créez un fichier `.env` dans le répertoire racine du projet et spécifiez les variables d'environnement comme mentionné dans la documentation du projet. voir [config](./docs/config.md#configuration)

4. Chargez les Variables d'Environnement :

   Utilisez le package `config` mentionné précédemment dans votre projet pour charger les variables d'environnement à partir du fichier `env`.

   Voici un exemple :

   ```go
   import "forum/config"

   // Chargez les variables d'environnement
   err := config.LoadEnv()
   if err != nil {
       // Gérez l'erreur
       panic("Erreur lors du chargement du fichier .env : " + err.Error())
   }
   ```

5. Compilez et Exécutez :

   Utilisez la commande `go run` pour compiler et exécuter votre projet :

   ```shell
   go run main.go
   ```

6. Accédez au Forum :

   Une fois votre projet en cours d'exécution, accédez au forum en ouvrant un navigateur web et en accédant à l'URL appropriée (généralement `http://localhost:PORT` où `PORT` est le port sur lequel votre serveur web écoute).

### Contributors

- #### [bcoulibal](http://learn.zone01dakar.sn/git/bcoulibal) (Bounama Coulibaly)
- #### [nmbengue](http://learn.zone01dakar.sn/git/nmbengue) (Ndieme Mbengue)
- #### [ssock](http://learn.zone01dakar.sn/git/ssock) (Serine Saliou Sock)
- #### [mamdrame](http://learn.zone01dakar.sn/git/mamdrame) (Mamour Ousmane Drame)
