$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(e) {
    e.preventDefault();
    let senha = $("#senha").val()
    let confirmarSenha = $("#confirmar-senha").val()

    if (senha != confirmarSenha) {
        alert("A senhas são diferentes")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: senha
        },
        
    }).done(function () {
        alert("Usuário Cadastrado com sucesso")
    }).fail(function(erro) {
        console.log(erro.responseJSON);
        alert("Erro ao cadastrar Usuário")
    });
}