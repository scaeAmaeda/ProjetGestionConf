package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/sqlite" // Import du pilote SQLite
)

func main() {
	// db
	fournisseur1 := fournisseur{nom: "Primeur", localisation: "Paris 15, coin de la rue", typeFournisseur: "détaillant"}
	fournisseur2 := fournisseur{nom: "Leclerc", localisation: "Paris 15, Bourcicault", typeFournisseur: "Grande Surface"}
	Poireau := ingredient{nom: "Poireau", typeIngredient: "primaire", fournisseurIngredient: fournisseur1}
	Carottes := ingredient{nom: "Carottes", typeIngredient: "primaire", fournisseurIngredient: fournisseur1}
	Courgettes := ingredient{nom: "Courgettes", typeIngredient: "primaire", fournisseurIngredient: fournisseur1}
	Cubor := ingredient{nom: "Cubor", typeIngredient: "primaire", fournisseurIngredient: fournisseur2}

	platsStatuts := []string{"Open", "In Work", "Prototype Released", "Serial Released"}
	soupe1 := plat{nom: "Soupe Standarde", statut: platsStatuts[2], typage: "soupe", ingredient: []ingredient{Poireau, Carottes, Courgettes, Cubor}, ingredientQuantite: []int{1, 2, 1, 1}, ingredientUnite: []string{"nb", "nb", "nb", "nb"}}

	// view
	afficherCompositionPlat(soupe1)
	tousLesPlats := []plat{}
	tousLesPlats = append(tousLesPlats, soupe1)
	afficherTousLesPlats(tousLesPlats)

	// Interractions avec db
	db := connectDB()
	defer db.Close()
	getSouSoupes(db)

}

type plat struct {
	nom                string
	statut             string
	typage             string
	ingredient         []ingredient
	ingredientQuantite []int
	ingredientUnite    []string
}
type ingredient struct {
	nom                   string
	typeIngredient        string
	fournisseurIngredient fournisseur
}

type fournisseur struct {
	nom             string
	localisation    string
	typeFournisseur string
}

// TO COMMENT
// type ingredientCI struct {
// 	nom string
// }

// type ECP struct {
// 	ECPNumber   string
// 	Product     string
// 	date        string
// 	WasTable    ECPWasTable
// 	IsTable     ECPIsTable
// 	effectivite []plat
// }

// type ECPWasTable struct {
// 	ingredientCI []ingredientCI
// 	ingredients  []ingredient
// }

// type ECPIsTable struct {
// 	ingredientCI []ingredientCI
// 	ingredients  []ingredient
// }

// func effectivityCalculation(){
// // Prend tous les ECP pour un produit donné,
// // Considère toutes les Was Table et les Is table
// }

// TO END COMMENT

func afficherCompositionPlat(p plat) {
	fmt.Printf("Le plat choisi : %v \n", p.nom)
	fmt.Printf("Le plat est de type : %v\n", p.typage)
	fmt.Printf("Et voici sa liste d'ingrédients : \n")
	for i := range p.ingredient {
		fmt.Printf("- %v %v de %v \n", p.ingredientQuantite[i], p.ingredientUnite[i], p.ingredient[i].nom)
	}
}

func afficherTousLesPlats(ps []plat) {
	for _, el := range ps {
		fmt.Println(el.nom)
	}
}

// func afficherLaProvenance(p plat) {
// 	// mettre une slice à qui on va ajouter un élement s'il est pas dedans déjjà
// }

// fonction créée pour ouvrir une bdd
func connectDB() *sql.DB {
	// Ouvre (ou crée) une base de données SQLite
	db, err := sql.Open("sqlite", "repositoryRecettes.sqlite")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données : %v", err)
	}

	// Vérifie que la connexion est opérationnelle
	if err := db.Ping(); err != nil {
		log.Fatalf("Impossible de se connecter à la base : %v", err)
	}

	fmt.Println("Connexion réussie à la base de données.")
	return db
}

func getSouSoupes(db *sql.DB) {
	rows, err := db.Query("SELECT nom, etat, typage FROM soupes")
	if err != nil {
		log.Fatalf("Erreur lors du select : %v", err)
	}
	defer rows.Close()

	fmt.Println("Liste des sousoupes :")
	for rows.Next() {
		var nom string
		var etat string
		var typage string
		err = rows.Scan(&nom, &etat, &typage)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture des lignes : %v", err)
		}
		fmt.Printf("ID: %s, Nom: %s, Type: %s\n", nom, etat, typage)
	}
}
