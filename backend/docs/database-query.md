
# database/query

## Aperçu

Le package `query` simplifie le processus de construction de requêtes SQL en fournissant un ensemble de fonctions pour générer des opérations SQL courantes. Ces fonctions se chargent de formater correctement les instructions SQL, rendant les interactions avec la base de données plus pratiques et moins sujettes aux erreurs.

## Utilisation
Importez le package `query` dans votre projet GoLang de la manière suivante :

```go
import "forum/query"
```
Pour utiliser les fonctions fournies par le package `query`, appelez-les et transmettez les paramètres requis. Ces fonctions renverront des chaînes de requête SQL que vous pourrez exécuter en utilisant votre connexion à la base de données SQLite.

### Fonctions

#### `UpdateQuery(table string, object any, where WhereOption) string`

Génère une requête SQL `UPDATE` pour mettre à jour des enregistrements dans une table en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à mettre à jour.
- `object` (any) : Un objet représentant les données à mettre à jour (doit etre une structure).
- `where` (WhereOption) : Un map spécifiant les conditions de mise à jour des enregistrements.

Renvoie une chaîne de requête SQL `UPDATE` formatée.

#### `DeleteQuery(table string, where WhereOption) string`

Génère une requête SQL `DELETE` pour supprimer des enregistrements d'une table en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle supprimer des enregistrements.
- `where` (WhereOption) : Un map spécifiant les conditions de suppression.

Renvoie une chaîne de requête SQL `DELETE` formatée.

#### `SelectOneFrom(table string, where WhereOption) string`

Génère une requête SQL `SELECT` pour récupérer un seul enregistrement d'une table en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle sélectionner un enregistrement.
- `where` (WhereOption) : Un map spécifiant les conditions de sélection d'un enregistrement.

Renvoie une chaîne de requête SQL `SELECT` formatée.

#### `SelectAllFrom(table string) string`

Génère une requête SQL `SELECT` pour récupérer tous les enregistrements d'une table.

- `table` (string) : Le nom de la table à partir de laquelle sélectionner les enregistrements.

Renvoie une chaîne de requête SQL `SELECT` formatée.

#### `SelectAllWhere(table string, where WhereOption) string`

Génère une requête SQL `SELECT` pour récupérer des enregistrements d'une table en fonction de conditions spécifiées.

- `table` (string) : Le nom de la table à partir de laquelle sélectionner les enregistrements.
- `where` (WhereOption) : Un map spécifiant les conditions de sélection des enregistrements.

Renvoie une chaîne de requête SQL `SELECT` formatée.

#### `InsertQuery(table string, object any) string`

Génère une requête SQL `INSERT` pour ajouter de nouveaux enregistrements à une table.

- `table` (string) : Le nom de la table dans laquelle insérer les enregistrements.
- `object` (any) : Un objet représentant les données à insérer.

Renvoie une chaîne de requête SQL `INSERT` formatée.

#### `AllTablesQuery() string`

Génère une requête SQL pour récupérer les noms de toutes les tables dans la base de données SQLite.

Renvoie une chaîne de requête SQL formatée.

---

Avec le package `query`, vous pouvez facilement créer des requêtes SQL pour des opérations courantes de base de données dans votre projet GoLang, ce qui rend les interactions avec la base de données plus efficaces et moins sujettes aux erreurs.