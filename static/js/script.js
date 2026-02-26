function changeMessage() {
    const messageElement = document.getElementById('message');
    const currentMessage = messageElement.textContent;

    const alternateMessage = [
        "Messi es el Mejor del mundo",
        "Pedri el mejor centro del mundo",
        "Vini Jr llora mucho"
    ];

    let newMEssage;

    do {
        newMEssage = alternateMessage[Math.floor(Math.random() * alternateMessage.length)];
    } while (newMEssage === currentMessage);

    messageElement.textContent = newMEssage;
}