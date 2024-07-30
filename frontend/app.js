document.getElementById('dataForm').addEventListener('submit', function(event) {
    event.preventDefault();

    // // Obtener los datos del formulario
    // const source = document.getElementById('source').value;
    // const destination = document.getElementById('destination').value;

    // Crear un objeto con los datos
    const data = {
        source: source,
        destination: destination
    };

    // Hacer la solicitud POST con fetch
    fetch('http://localhost:8090/get-translations', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        document.getElementById('response').innerText = `Response: ${JSON.stringify(data)}`;
    })
    .catch((error) => {
        console.error('Error:', error);
    });
});
