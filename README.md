# http_go

Ce programme Go récupère des informations sur les modules Go depuis l'index Go.

Écrire un programme Go index_godev qui interroge en HTTP l'index officiel des modules, accumule par forge les modules et leurs versions, et produit une table par forge similaire à:

                                   Forge Modules Versions
                              github.com  235782  1464248
                                  k8s.io     332    70345
                              gitlab.com    4086    24155
                                gopkg.in    2134    11976
                                         . . . . . . .
                      gitlab.brurberg.no       1        1
                                _Totals_  253732  1692307

## Installation


go get -u https://github.com/ferielbelouch/http_go


Pour utiliser le programme, suivez ces étapes :

1. Clonez le projet : git clone https://github.com/ferielbelouch/http_go.git

2. Accédez au répertoire du projet : cd votre-projet

3. Exécutez le programme : go run main.go

## Contributions

Les contributions sont les bienvenues ! Si vous souhaitez contribuer à ce projet, veuillez suivre ces étapes :

1. Fork du projet.
2. Créez une branche pour votre fonctionnalité : git checkout -b nouvelle_fonctionnalité
3. Faites vos modifications 
4. Effectuez un commit de vos modifications : git commit -m "Ajout de ma nouvelle fonctionnalité"
5. Poussez vos modifications vers votre fork : git push origin ma-nouvelle-fonctionnalite
6. Créez une demande d'extraction sur GitHub.

## Licence : Ce projet est sous licence MIT - voir le fichier LICENSE pour plus de détails.

## Remarques supplémentaires :

1. Assurez-vous d'avoir Go installé sur votre machine.
2. Ce programme utilise les packages standard de Go et ne nécessite pas d'installation supplémentaire.
3. Pour obtenir de l'aide ou signaler des problèmes, veuillez ouvrir une issue sur GitHub.