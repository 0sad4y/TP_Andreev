document.addEventListener("DOMContentLoaded", () => {
    const tbody = document.getElementById("stats-body");
    tbody.innerHTML = "";

    const row = `
        <tr>
            <td>${employeeData.tripCount}</td>
            <td>${employeeData.moneySpent}</td>
            <td>${employeeData.avgTripCount}</td>
            <td>${employeeData.avgMoneySpent}</td>
        </tr>`;
    tbody.insertAdjacentHTML("beforeend", row);

    const labels = chartData.map(d => d.x);
    const data = chartData.map(d => d.y);

    const ctx = document.getElementById('chart').getContext('2d');
    new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels,
            datasets: [{
                label: '',
                data: data,
                borderColor: '#4cb00a',
                backgroundColor: 'rgba(60,179,113,0.2)',
                tension: 0.3,
                pointStyle: false,
            }]
        },
        options: {
            responsive: true,
            plugins: { legend: { display: false } },
            scales: { y: { beginAtZero: true } }
        }
    });
});