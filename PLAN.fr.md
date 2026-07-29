# sbx — Zero to Hero · Plan de l'atelier (FR)

> ⚠️ **Document vivant.** Ce plan décrit le périmètre et le déroulé prévus de l'atelier pratique. Le contenu est en cours d'élaboration : les modules, l'ordre, les commandes et les exemples **peuvent évoluer**, et l'atelier peut être **adapté** à la volée en fonction du public, du temps disponible et de la version de `sbx` utilisée (certaines fonctionnalités sont encore **expérimentales**). Les durées indiquées sont indicatives, pas contractuelles.
>
> En particulier :
> - Le contenu **doit être adapté en fonction de l'agent d'IA sélectionné**. Cette édition utilise l'agent **Codex** (OpenAI), mais `sbx` peut exécuter d'autres agents — le flux de connexion, les prompts et certaines commandes différeront en conséquence.
> - Le **nombre d'étapes et de modules variera selon la durée du workshop** : une session courte ne couvre que les premiers modules, une session complète va jusqu'au bout. Les formateurs choisissent, réduisent ou réordonnent les étapes selon le temps disponible.

## En quoi consiste cet atelier ?

Une introduction pratique, en mode « on apprend en faisant », aux **Docker Sandboxes (`sbx`)** — une manière d'exécuter des agents de code IA (ici l'agent **Codex** d'OpenAI) dans des VM isolées et pilotées par des politiques, plutôt que directement sur votre machine.

Tout au long de l'atelier, vous pilotez un véritable agent à l'intérieur d'une sandbox : vous publiez des ports, injectez des variables d'environnement et des secrets, contrôlez ce que la sandbox peut atteindre sur le réseau et sur le disque, puis vous empaquetez et partagez vos propres environnements de sandbox reproductibles.

> ✋ Cet atelier utilise `sbx` et l'agent **Codex** authentifié avec une **clé API OpenAI**. Enregistrez-la une fois avec `sbx secret set -g openai` ; le proxy l'injecte au moment de la requête et la sandbox ne voit jamais qu'un placeholder `proxy-managed`.

## À qui s'adresse-t-il ?

- Développeurs et ingénieurs plateforme / DevOps curieux d'exécuter des agents IA en toute sécurité.
- Aucune expertise Docker poussée requise, mais une aisance de base avec le terminal est attendue.
- L'application de démonstration est un tout petit microservice **Go** — aucune connaissance de Go nécessaire.

## Objectifs pédagogiques

À la fin, les participants savent :
1. Créer une sandbox et y exécuter un agent IA.
2. Faire construire, exécuter et exposer un petit service par l'agent, puis commiter/pousser sur GitHub.
3. Injecter des variables d'environnement et des **secrets** que l'agent ne voit jamais en clair.
4. Restreindre une sandbox avec des politiques **réseau** et **système de fichiers**.
5. Personnaliser et partager des environnements de sandbox via **kits**, **snapshots** et **templates**.

## Prérequis

- Un compte GitHub (handle) et un dépôt nommé `hello-sbx` pour travailler.
- Une **clé API OpenAI** (`sk-...`), enregistrée une fois avec `sbx secret set -g openai`.
- `sbx` installé et configuré — voir <https://docs.docker.com/ai/sandboxes/#get-started>.
- VS Code + un terminal.

## Déroulé de l'atelier

### Module 0 — Introduction & installation *(prévu / en évolution)*
- Pourquoi des sandboxes : le problème d'« évasion de conteneur » / du rayon d'impact pour les agents IA.
- Ce qu'est `sbx` et comment s'y connecter.
- Vérification des prérequis.

### Module 1 — Créer et explorer une sandbox
- `sbx run codex` — créer sa première sandbox ; la clé API OpenAI est injectée par le proxy (aucune connexion dans la sandbox).
- Entrer dans le shell de la sandbox, inspecter son système de fichiers, faire exécuter des tâches simples à l'agent.

### Module 2 — Interagir avec GitHub
- Demander à l'agent d'échafauder un petit microservice Go avec un endpoint JSON `/hello`.
- **Publier un port** vers l'hôte (`sbx ports … --publish`) et tester avec `curl`.
- Commiter & pousser le code (via le token GitHub injecté par le proxy), et créer une issue GitHub depuis l'agent avec la CLI `gh`.

### Module 3 — Variables d'environnement & secrets
- Injecter une variable d'environnement persistante dans une sandbox en cours d'exécution.
- **Injection de secret personnalisé** : la sandbox ne voit qu'un *placeholder*, tandis que le proxy substitue la vraie valeur dans les requêtes sortantes — démontré face à un `tiny-service` local qui valide un jeton bearer.

### Module 4 — Politiques : contrôler les accès
- **Politiques réseau** (`04-01`) : autoriser/refuser des domaines, portée globale vs par sandbox, observer le **HTTP 403** structuré, auditer avec `sbx policy log`.
- **Workspaces** (`04-02`) : monter des répertoires de l'hôte en lecture-écriture / lecture seule, prouver que le reste de l'hôte reste invisible.
- **Politiques système de fichiers** (`04-03`) : règles de lecture vs écriture, et comment les workspaces en lecture seule sont le moyen local de restreindre les écritures.

### Module 5 — Personnaliser & partager des environnements
- **Kits** (`05-01`) : add-ons déclaratifs de type builder/launcher écrits avec le **schéma de kit v2** ; partage via ZIP, registre OCI ou git. *(expérimental)*
- **Snapshots** (`05-02`) : figer une sandbox préparée dans une image de template réutilisable.
- **Template personnalisé depuis un Dockerfile** (`05-03`) : construire une image de base reproductible.
- **Construire un template dans une sandbox** (`05-04`) : quand l'hôte n'a pas Docker, utiliser le moteur Docker intégré de la sandbox.

### Sujets sur la feuille de route *(couverts ou non selon la session)*
Politiques de gouvernance IA, audit, **mode clone**, git worktrees et workspaces multiples — ces sujets sont en cours de rédaction et pourront être ajoutés ou adaptés selon la session.

## Format & logistique

- **Style :** 100 % pratique ; chaque participant travaille dans son propre dépôt `hello-sbx`.
- **Durée :** modulaire — une session courte peut couvrir les modules 1 à 3 ; une session complète va jusqu'au module 5. Les formateurs choisissent et réordonnent les modules selon le groupe.
- **Nettoyage :** la plupart des leçons se terminent par des commandes de nettoyage (`sbx rm …`, `sbx template rm …`).

---
*Comme `sbx` évolue rapidement, vérifiez toujours les commandes avec la documentation officielle : <https://docs.docker.com/ai/sandboxes/>. Ce plan sera mis à jour au fur et à mesure que l'outillage et l'atelier mûrissent.*
