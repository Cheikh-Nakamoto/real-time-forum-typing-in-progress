
# Package Database

Le package `database` gère les opérations de base de données pour votre application web de forum GoLang. Il utilise SQLite comme moteur de base de données sous-jacent et fournit des fonctions pour les opérations courantes de la base de données telles que l'insertion, la suppression, la mise à jour et la requête.

## Table des matières

- [Aperçu](#aperçu)
- [Fonctions](#fonctions)
  - [Init](#init)
  - [Insert](#insert)
  - [Delete](#delete)
  - [Update](#update)
  - [GetOneFrom](#getonefrom)
  - [GetAllFrom](#getallfrom)
  - [GetColumnsName](#getcolumnsname)

## Aperçu

Le package `database` sert d'interface entre votre application GoLang et la base de données SQLite. Il établit une connexion à la base de données, initialise les tables de la base de données et fournit des méthodes pour les opérations courantes de la base de données.

## Utilisation

Importez le package `database` dans votre projet GoLang de la manière suivante :

```go
import "forum/database"
```
Pour utiliser les fonctions fournies par le package `database`, vous devez d'abord initialiser la connexion à la base de données en utilisant la fonction `Init`. Après l'initialisation, vous pouvez utiliser les différentes méthodes pour interagir avec votre base de données SQLite.

### Fonctions

#### `Init() *Database`

Initialise la connexion à la base de données et crée les tables de la base de données en fonction du schéma SQL fourni dans un fichier.

- Renvoie un pointeur vers la base de données initialisée.

#### `Insert(table string, data any) error`

Insère des données dans la table spécifiée.

- `table` (string) : Le nom de la table dans laquelle insérer les données.
- `data` (any) : Les données à insérer dans la table.

- Renvoie une erreur si l'insertion échoue.

#### `Delete(table string, where q.WhereOption) error`

Supprime des enregistrements de la table spécifiée en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle supprimer des enregistrements.
- `where` (q.WhereOption) : Un map spécifiant les conditions de suppression.

- Renvoie une erreur si la suppression échoue.

#### `Update(table string, object any, where q.WhereOption) error`

Met à jour des enregistrements de la table spécifiée en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table dans laquelle mettre à jour les enregistrements.
- `object` (any) : Un objet représentant les données à mettre à jour.
- `where` (q.WhereOption) : Un map spécifiant les conditions de mise à jour des enregistrements.

- Renvoie une erreur si la mise à jour échoue.

#### `GetOneFrom(table string, where q.WhereOption) *sql.Row`

Récupère un seul enregistrement de la table spécifiée en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle sélectionner un enregistrement.
- `where` (q.WhereOption) : Un map spécifiant les conditions de sélection d'un enregistrement.

- Renvoie un pointeur vers une seule ligne de résultat.

#### `GetAllFrom(table string, where q.WhereOption) *sql.Rows`

Récupère tous les enregistrements de la table spécifiée ou en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle sélectionner les enregistrements.
- `where` (q.WhereOption) : Un map spécifiant les conditions de sélection des enregistrements.

- Renvoie un pointeur vers un ensemble de résultats contenant plusieurs lignes.

#### `GetColumnsName() []string`

Récupère les noms de toutes les tables de la base de données.

- Renvoie une slice (tableau dynamique) de noms de tables.

---

Avec le package `database`, vous pouvez gérer efficacement vos opérations de base de données dans votre application web de forum GoLang, y compris l'initialisation, l'insertion, la suppression, la mise à jour et la requête des enregistrements de la base de données.