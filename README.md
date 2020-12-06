![License](https://img.shields.io/github/license/HETIC-MT-P2021/RPGo)
![golang](https://img.shields.io/github/languages/top/HETIC-MT-P2021/RPGo)
![golang-version](https://img.shields.io/github/go-mod/go-version/HETIC-MT-P2021/RPGo)
![commit](https://img.shields.io/github/last-commit/HETIC-MT-P2021/RPGo)
![build-CI](https://img.shields.io/github/workflow/status/HETIC-MT-P2021/RPGo/CI)

# RPGo 💣

Simple Discord Bot for a School project written in GO.
This bot will allow you to create your own RPG character and fight against enormous monsters to become the strongest one.

## Project features

* ✅ Create a character
* ⏱ Have specific character classes
* ⏱ Have a random background spawned per character created

## Authors

[Athénais Dussordet](https://github.com/Araknyfe)

[Alexandre Lellouche](https://github.com/AlexandreLch)

Last but not least : [Thomas Raineau](https://github.com/Traineau)

## Usage

```git config core.hooksPath .githooks```
> Configure GitHooks

```cp docker-compose.yaml.dist docker-compose.yaml```
> Docker configuration override, don't forget to add the Token and SQL variables

``` docker-compose up --build```
> Run the project

## Resources

External libraries used : 
* [discordgo](https://github.com/bwmarrin/discordgo) : used to connect to the Discord API
* [gomock](https://github.com/golang/mock/gomock) : used to mock functions in command tests
* [sqlmock](https://github.com/DATA-DOG/go-sqlmock) : used to mock sql functions in repository test


## Pattern implemented

The command pattern is implemented throughout the project as it is well suited for an RPG Discord bot. [Here](https://refactoring.guru/design-patterns/command) is some documentation on how we implemented it.
