const rowsPerPage = 5;
let currentPage = 1;

function renderTable(page = 1) {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    const pageData = tableData.slice(start, end);

    const tbody = document.getElementById("data-body");
    tbody.innerHTML = "";

    pageData.forEach(item => {
        const row = `
            <tr>
                <td>
                    <a class="employeeLink" href="/employee/${item.id}">
                        ${item.name}
                    </a>
                </td>
                <td>${item.destination}</td>
                <td>${item.date}</td>
                <td>${item.duration}</td>
                <td>${item.moneySpent}</td>
            </tr>`;
        tbody.insertAdjacentHTML("beforeend", row);
    });
}

function renderPagination() {
    const totalPages = Math.ceil(tableData.length / rowsPerPage);
    const pagination = document.getElementById("pagination");
    pagination.innerHTML = "";

    const pageLimit = 10;
    let startPage = Math.max(1, currentPage - Math.floor(pageLimit / 2));
    let endPage = startPage + pageLimit - 1;
    if (endPage > totalPages) {
        endPage = totalPages;
        startPage = Math.max(1, endPage - pageLimit + 1);
    }

    pagination.insertAdjacentHTML("beforeend", `
        <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
            <button class="page-link">&laquo;</button>
        </li>
    `);

    if (startPage > 1) {
        pagination.insertAdjacentHTML("beforeend", `
            <li class="page-item"><button class="page-link">1</button></li>
            <li class="page-item disabled"><span class="page-link">...</span></li>
        `);
    }

    for (let i = startPage; i <= endPage; i++) {
        pagination.insertAdjacentHTML("beforeend", `
            <li class="page-item ${i === currentPage ? 'active' : ''}">
                <button class="page-link">${i}</button>
            </li>
        `);
    }

    if (endPage < totalPages) {
        pagination.insertAdjacentHTML("beforeend", `
            <li class="page-item disabled"><span class="page-link">...</span></li>
            <li class="page-item"><button class="page-link">${totalPages}</button></li>
        `);
    }

    pagination.insertAdjacentHTML("beforeend", `
        <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
            <button class="page-link">&raquo;</button>
        </li>
    `);

    const buttons = pagination.querySelectorAll(".page-link");
    buttons.forEach(btn => {
        btn.addEventListener("click", () => {
            const text = btn.textContent;
            if (text === '«' && currentPage > 1) currentPage--;
            else if (text === '»' && currentPage < totalPages) currentPage++;
            else if (!isNaN(text)) currentPage = Number(text);

            renderTable(currentPage);
            renderPagination();
        });
    });
}

document.addEventListener("DOMContentLoaded", () => {
    renderTable();
    renderPagination();
});