document.addEventListener('DOMContentLoaded', () => {
    console.log("Funcion ejecutada")
    const apiUrl = 'http://localhost:8090/get-translations';
    let color = true;

    fetch(apiUrl,{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            const tableBody = document.querySelector('#data-table tbody');

            data.forEach(item => {
                const row = document.createElement('tr');
                console.log(color)
                if (color == true){
                    color = false;
                }
                else{
                    row.classList = "bg-gray-100"
                    color = true
                }

                const cell1 = document.createElement('td');
                cell1.textContent = item.Source;
                cell1.classList = "py-2 px-4 border-b";
                row.appendChild(cell1);

                const cell2 = document.createElement('td');
                cell2.textContent = item.Destination;
                cell2.classList = "py-2 px-4 border-b";
                row.appendChild(cell2);
                tableBody.appendChild(row);
                console.log("Valores obtenidos correctamente.")
            });
        })
        .catch(error => {
            console.error('Hubo un problema con la solicitud Fetch:', error);
        });
});

// document.getElementById('dataForm').addEventListener('submit', function(event) {
//     event.preventDefault();

//     const source = document.getElementById('source').value;

//     const data = {
//         source: source,
//     };

//     // Hacer la solicitud POST con fetch
//     fetch('http://localhost:8090/generate-translation', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify(data)
//     })
//     .then(response => response.json())
//     .then(data => {
//         console.log('Success:', data);
//         document.getElementById('response').innerText = `Response: ${JSON.stringify(data)}`;
//     })
//     .catch((error) => {
//         console.error('Error:', error);
//     });
// });

