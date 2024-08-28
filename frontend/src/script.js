const data = {
  data: [
    {
      author: "glinda.ebert",
      avatar:
        "https://robohash.org/inventorerecusandaemagnam.png?size=300x300&set=set1",
      duration: "2 min",
      title: "Exploration des fonctionnalités avancées de Golang 🚀",
      content: "Découvrez les secrets cachés de Golang aujourd'hui!",
      categories: ["Golang", "Programmation"],
      comments: 15,
      likes: 30,
      dislikes: 2,
    },
    {
      email: "rodger.price@email.com",
      avatar: "https://robohash.org/distinctioquiaea.png?size=300x300&set=set1",
      duration: "12 min",
      title: "Les meilleures pratiques pour optimiser votre code JavaScript 🔍",
      content:
        "Boostez les performances de votre code avec ces astuces essentielles! #JavaScript #Optimisation",
      categories: ["JavaScript", "Optimisation"],
      comments: 22,
      likes: 45,
      dislikes: 3,
    },
    {
      email: "rodger.price@email.com",
      avatar: "https://robohash.org/distinctioquiaea.png?size=300x300&set=set1",
      duration: "22 min",
      title:
        "Les étapes essentielles pour devenir un développeur Python compétent 🚀🐍",
      content:
        "Parcourez le chemin pour devenir un expert en Python avec ces conseils pratiques! #Python #DevJourney",
      categories: ["Python", "Career Development"],
      comments: 30,
      likes: 55,
      dislikes: 3,
    },
    {
      email: "rodger.price@email.com",
      avatar: "https://robohash.org/distinctioquiaea.png?size=300x300&set=set1",
      duration: "1h 20 min",
      title: "Découvrez les secrets de l'optimisation des requêtes SQL! 📊💾",
      content:
        "Améliorez les performances de votre base de données en optimisant vos requêtes SQL! #SQL #Database",
      categories: ["SQL", "Database Optimization"],
      comments: 20,
      likes: 42,
      dislikes: 5,
    },
    {
      email: "rodger.price@email.com",
      avatar: "https://robohash.org/distinctioquiaea.png?size=300x300&set=set1",
      duration: "52 min",
      title: "Les défis intrigants du développement en GoLang! 🤔🚀",
      content:
        "Explorez les aspects fascinants du développement avec le langage GoLang! #GoLang #DevChallenges",
      categories: ["GoLang", "Development Challenges"],
      comments: 22,
      likes: 48,
      dislikes: 2,
    },
    {
      username: "shantel.hettinger",
      email: "shantel.hettinger@email.com",
      avatar:
        "https://robohash.org/omnismaioresmagnam.png?size=300x300&set=set1",
      duration: "30 min",
      title: "Construire une application web avec Python et Flask 🐍🌐",
      content:
        "Un guide étape par étape pour créer une application web robuste! #Python #Flask",
      categories: ["Python", "Flask", "Web Development"],
      comments: 18,
      likes: 38,
      dislikes: 1,
    },
    {
      username: "chris.mills",
      email: "chris.mills@email.com",
      avatar: "https://robohash.org/voluptasquasqui.png?size=300x300&set=set1",
      duration: "1h 10 min",
      title: "Les bases du développement web : HTML & CSS 💻🎨",
      content:
        "Plongez dans le monde du développement web avec HTML et CSS! #HTML #CSS #WebDev",
      categories: ["HTML", "CSS", "Web Development"],
      comments: 12,
      likes: 25,
      dislikes: 4,
    },
    {
      username: "joesph.thompson",
      email: "joesph.thompson@email.com",
      avatar:
        "https://robohash.org/placeatpossimusminus.png?size=300x300&set=set1",
      duration: "10 min",
      title:
        "Les avantages de l'utilisation de C dans le développement système ⚙️🖥️",
      content:
        "Découvrez pourquoi le langage C est incontournable pour le développement système! #C #SystemProgramming",
      categories: ["C", "System Programming"],
      comments: 20,
      likes: 42,
      dislikes: 5,
    },
    {
      username: "jesus.steuber",
      email: "jesus.steuber@email.com",
      avatar: "https://robohash.org/aliquammodivel.png?size=300x300&set=set1",
      duration: "1h 30 min",
      title:
        "Trucs et astuces pour résoudre les problèmes de performance en Python 🐍⚡",
      content:
        "Optimisez vos scripts Python et améliorez les performances globales! #Python #Performance",
      categories: ["Python", "Performance"],
      comments: 25,
      likes: 50,
      dislikes: 2,
    },
    {
      username: "jesus.steuber",
      email: "jesus.steuber@email.com",
      avatar: "https://robohash.org/aliquammodivel.png?size=300x300&set=set1",
      duration: "30 min",
      title: "Démystification des algorithmes de tri en C 🔄💻",
      content:
        "Comprenez le fonctionnement des algorithmes de tri classiques en langage C! #C #Algorithms",
      categories: ["C", "Algorithms"],
      comments: 18,
      likes: 38,
      dislikes: 3,
    },
    {
      username: "jesus.steuber",
      email: "jesus.steuber@email.com",
      avatar: "https://robohash.org/aliquammodivel.png?size=300x300&set=set1",
      duration: "1h",
      title: "Les tendances émergentes en développement web en 2023! 🌐🚀",
      content:
        "Restez à jour avec les dernières tendances qui façonnent le paysage du développement web! #WebDev #Trends",
      categories: ["Web Development", "Trends"],
      comments: 23,
      likes: 47,
      dislikes: 4,
    },
    {
      username: "jesus.steuber",
      email: "jesus.steuber@email.com",
      avatar: "https://robohash.org/aliquammodivel.png?size=300x300&set=set1",
      duration: "2h 30 min",
      title: "Guide complet pour maîtriser les promesses en JavaScript 🤝💻",
      content:
        "Découvrez les promesses en JavaScript et améliorez la gestion asynchrone dans votre code! #JavaScript #Promises",
      categories: ["JavaScript", "Promises"],
      comments: 25,
      likes: 50,
      dislikes: 2,
    },
    {
      username: "dominic.conroy",
      email: "dominic.conroy@email.com",
      avatar:
        "https://robohash.org/estdignissimosest.png?size=300x300&set=set1",
      duration: "1h 20 min",
      title: "Les nouveautés passionnantes de JavaScript ES2022! 🚀",
      content:
        "Explorez les fonctionnalités étonnantes de la dernière version d'ECMAScript! #JavaScript #ES2022",
      categories: ["JavaScript", "ES2022"],
      comments: 15,
      likes: 28,
      dislikes: 3,
    },
    {
      username: "dominic.conroy",
      email: "dominic.conroy@email.com",
      avatar:
        "https://robohash.org/estdignissimosest.png?size=300x300&set=set1",
      duration: "1h",
      title: "Rétrospective sur les langages de programmation vintage 🕰️💾",
      content:
        "Replongez dans l'histoire des langages de programmation qui ont marqué leur époque! #ProgrammingHistory",
      categories: ["Programming History"],
      comments: 15,
      likes: 35,
      dislikes: 2,
    },
    {
      username: "dominic.conroy",
      email: "dominic.conroy@email.com",
      avatar:
        "https://robohash.org/estdignissimosest.png?size=300x300&set=set1",
      duration: "1h 50 min",
      title: "Les bases du développement sécurisé en Python 🛡️🐍",
      content:
        "Protégez votre code Python en suivant ces bonnes pratiques de sécurité! #Python #Security",
      categories: ["Python", "Security"],
      comments: 20,
      likes: 42,
      dislikes: 5,
    },
    {
      username: "abbey.carter",
      email: "abbey.carter@email.com",
      avatar:
        "https://robohash.org/quiaetaccusantium.png?size=300x300&set=set1",
      duration: "1h 30 min",
      title: "Créer des interfaces utilisateur modernes avec ReactJS ⚛️🌐",
      content:
        "Guide complet pour développer des interfaces utilisateur réactives avec React! #ReactJS #UI",
      categories: ["ReactJS", "UI Development"],
      comments: 30,
      likes: 55,
      dislikes: 4,
    },
    {
      username: "clark.grant",
      email: "clark.grant@email.com",
      avatar:
        "https://robohash.org/aliquampraesentiumvoluptas.png?size=300x300&set=set1",
      duration: "1h 30 min",
      title: "Les défis passionnants du développement en langage C 🤔💻",
      content:
        "Plongez dans le monde du C et relevez les défis stimulants! #C #DevChallenges",
      categories: ["C", "Development Challenges"],
      comments: 18,
      likes: 40,
      dislikes: 6,
    },
    {
      username: "maxie.bernier",
      email: "maxie.bernier@email.com",
      avatar:
        "https://robohash.org/consequunturvelitrem.png?size=300x300&set=set1",
      duration: "1h 30 min",
      title: "Introduction à l'intelligence artificielle avec Python 🤖🐍",
      content:
        "Démystifiez l'IA en apprenant à construire des modèles avec Python! #Python #AI",
      categories: ["Python", "Artificial Intelligence"],
      comments: 22,
      likes: 48,
      dislikes: 2,
    },
    {
      username: "maxie.bernier",
      email: "maxie.bernier@email.com",
      avatar:
        "https://robohash.org/consequunturvelitrem.png?size=300x300&set=set1",
      duration: "1h 50 min",
      title: "Les secrets du succès en développement Full Stack 🔍💻",
      content:
        "Explorez les compétences essentielles pour devenir un développeur Full Stack accompli! #FullStack #DevLife",
      categories: ["Full Stack", "Development"],
      comments: 28,
      likes: 60,
      dislikes: 3,
    },
    {
      username: "maxie.bernier",
      email: "maxie.bernier@email.com",
      avatar:
        "https://robohash.org/consequunturvelitrem.png?size=300x300&set=set1",
      duration: "1h 30 min",
      title: "Les frameworks JavaScript les plus populaires en 2023! 📈🔥",
      content:
        "Découvrez les frameworks qui dominent la scène du développement web cette année! #JavaScript #Frameworks",
      categories: ["JavaScript", "Frameworks"],
      comments: 25,
      likes: 52,
      dislikes: 4,
    },
  ],
};
