package main

import (
	"./model"
	"./ui"
	"fmt"
)

const OPTIONS int = 7

var users []model.User

func main() {

	var option = 0

	for option != 6 {

		fmt.Println("Seleccione uma opção:\n" +
			"1 - Criar utilizador\n" +
			"2 - Listar utilizadores \n" +
			"3 - Procurar utilizador \n" +
			"4 - Actualizar utilizador \n" +
			"5 - Carregar dados \n" +
			"6 - Gravar dados \n" +
			"7 - Sair")

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			ui.CreateUserUI(&users)
		case 2:
			ui.PrintUsersUI(users)
		case 3:
			ui.FindUserUI(&users)
		case 4:

		case 5:
			ui.ReadUsersUI(&users)
		case 6:
			ui.SaveUsersUI(users)
		case 7:
			break
		default:
			fmt.Println("Opção inválida.")
		}
	}

	fmt.Println("Programa encerrado.")
}
