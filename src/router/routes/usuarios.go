package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AutalizarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
