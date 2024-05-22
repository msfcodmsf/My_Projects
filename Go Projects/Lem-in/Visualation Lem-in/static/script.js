// SVG çizimini yap
function drawAntFarm(roomCoordinates, antPaths) {
    const svg = document.getElementById('ant-farm-svg');
    svg.innerHTML = ''; // Önceki çizimleri temizle

    // Odaları çiz
    Object.entries(roomCoordinates).forEach(([room, coords]) => {
        const circle = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
        circle.setAttribute('cx', coords[0]);
        circle.setAttribute('cy', coords[1]);
        circle.setAttribute('r', 10);
        circle.setAttribute('fill', 'red'); // İstenilen rengi belirtin
        svg.appendChild(circle);

        const text = document.createElementNS('http://www.w3.org/2000/svg', 'text');
        text.setAttribute('x', coords[0] + 15);
        text.setAttribute('y', coords[1]);
        text.setAttribute('fill', 'black');
        text.textContent = room;
        svg.appendChild(text);
    });

    // Karınca yollarını çiz
    antPaths.forEach(path => {
        const [start, end] = path.split(' ');
        const [x1, y1] = roomCoordinates[start];
        const [x2, y2] = roomCoordinates[end];
        const line = document.createElementNS('http://www.w3.org/2000/svg', 'line');
        line.setAttribute('x1', x1);
        line.setAttribute('y1', y1);
        line.setAttribute('x2', x2);
        line.setAttribute('y2', y2);
        line.setAttribute('stroke', 'black');
        svg.appendChild(line);
    });
}

// Terminalden girilen metin dosyasından verileri oku
function readInputTextFile() {
    fetch('/user_input.txt') // Kullanıcının terminalden girdiği metin dosyasının yolu
        .then(response => response.text())
        .then(text => {
            const lines = text.trim().split('\n');
            const roomCoordinates = {};
            const antPaths = [];
            lines.forEach(line => {
                const [identifier, ...values] = line.trim().split(' ');
                if (identifier === 'Room') {
                    const [room, x, y] = values;
                    roomCoordinates[room] = [parseInt(x), parseInt(y)];
                } else if (identifier === 'Path') {
                    const [start, end] = values;
                    antPaths.push(`${start} ${end}`);
                }
            });
            drawAntFarm(roomCoordinates, antPaths);
        })
        .catch(error => console.error('Error reading input text file:', error));
}

// Sayfa yüklendiğinde metin dosyasını oku ve SVG çizimini yap
document.addEventListener('DOMContentLoaded', readInputTextFile);
