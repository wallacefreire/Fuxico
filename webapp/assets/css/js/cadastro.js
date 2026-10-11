$('#formulario-cadastro').on('submit', CriarUsuario);

function CriarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val() ) {
        alert("As senhas não conferem!");
        return;
    } 

    $.ajax({
        url: '/usuarios',
        method: 'POST',
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        alert("Usuário criado com sucesso!");
    }).fail(function() {
        alert("Erro ao criar usuário!");
    });
}