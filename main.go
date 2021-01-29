package main

import (
	"./model"
	"./ui"
	"fmt"
)

const OPTIONS int = 6

var users []model.User

func main() {

	var option = 1

	for option > 0 && option < OPTIONS {

		fmt.Println("Seleccione uma opção:\n" +
			"1 - Criar utilizador\n" +
			"2 - Listar utilizadores \n" +
			"3 - Actualizar utilizador \n" +
			"4 - Carregar dados \n" +
			"5 - Gravar dados \n" +
			"6 - Sair")

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			ui.CreateUserUI(&users)
		case 2:
			ui.PrintUsersUI(users)
		case 3:

		case 4:
			ui.ReadUsersUI(&users)
		case 5:
			ui.SaveUsersUI(users)
		case 6:
			break
		default:
			fmt.Println("Opção inválida.")
		}
	}

	fmt.Println("Programa encerrado.")
}
