function DrawMoneySpentChart() {
    const labels = chartData1.map(d => d.x);
    const data = chartData1.map(d => d.y);

    const ctx = document.getElementById('chart_1').getContext('2d');
    new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels,
            datasets: [{
                label: '',
                data: data,
                borderColor: '#4cb00a',
                backgroundColor: 'rgba(46,139,87,0.2)',
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
}

function DrawTripChart() {
    const labels = chartData2.map(d => d.x);
    const data = chartData2.map(d => d.y);

    const ctx = document.getElementById('chart_2').getContext('2d');
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
            scales: { y: { 
                beginAtZero: true,
                ticks: { stepSize: 1 }
            }}
        }
    });
}

document.addEventListener("DOMContentLoaded", () => {
    DrawMoneySpentChart()
    DrawTripChart()
});
