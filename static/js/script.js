document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let user = document.getElementById("username").value;
    let pass = document.getElementById("password").value;
    let message = document.getElementById("message");

    // Exemplo simples de validação
    if(user === "admin" && pass === "1234") {
        message.style.color = "green";
        message.textContent = "Login realizado com sucesso!";
        // Refirecionar após login
        setTimeout(() => {
            window.location.href = "cadastra.html";
        }, 1500);
    } else {
        message.style.color = "red";
        message.textContent = "Usuário ou senha inválidos";
    }
});
