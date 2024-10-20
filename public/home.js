// ตรวจสอบว่าผู้ใช้ได้ล็อกอินแล้วหรือยัง
const token = localStorage.getItem('token');
if (!token) {
    window.location.href = "/public/login.html";
} else {
    const decodedToken = JSON.parse(atob(token.split('.')[1]));
    document.getElementById('userName').textContent = decodedToken.email;
}

// ฟังก์ชัน Logout
document.getElementById('logoutBtn').addEventListener('click', function() {
    localStorage.removeItem('token');
    window.location.href = "/public/index.html";
});

// ดึงข้อมูลการใช้จ่าย (Read)
async function fetchExpenses() {
    const response = await fetch('/expenses', {
        method: 'GET',
        headers: { 'Authorization': `Bearer ${token}` }
    });
    const expenses = await response.json();
    const expenseTable = document.getElementById('expenseTable').querySelector('tbody');
    expenseTable.innerHTML = ''; // ล้างข้อมูลก่อนแสดงการใช้จ่าย
    expenses.forEach(expense => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${expense.category}</td>
            <td>${expense.amount}</td>
            <td>${expense.date}</td>
            <td>
                <button onclick="editExpense(${expense.id})" class="button">แก้ไข</button>
                <button onclick="deleteExpense(${expense.id})" class="button">ลบ</button>
            </td>
        `;
        expenseTable.appendChild(row);
    });
}

// เพิ่มการใช้จ่าย (Create)
document.getElementById('addExpenseBtn').addEventListener('click', async function() {
    const category = prompt('กรอกประเภทการใช้จ่าย:');
    const amount = prompt('กรอกจำนวนเงิน:');
    const date = prompt('กรอกวันที่ (YYYY-MM-DD):');
    if (category && amount && date) {
        const response = await fetch('/expenses', {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ category, amount, date })
        });
        if (response.ok) {
            alert('เพิ่มการใช้จ่ายสำเร็จ');
            fetchExpenses(); // โหลดข้อมูลการใช้จ่ายใหม่
        } else {
            alert('การเพิ่มการใช้จ่ายล้มเหลว');
        }
    }
});

// แก้ไขการใช้จ่าย (Update)
async function editExpense(id) {
    const category = prompt('แก้ไขประเภทการใช้จ่าย:');
    const amount = prompt('แก้ไขจำนวนเงิน:');
    const date = prompt('แก้ไขวันที่ (YYYY-MM-DD):');
    if (category && amount && date) {
        const response = await fetch(`/expenses/${id}`, {
            method: 'PUT',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ category, amount, date })
        });
        if (response.ok) {
            alert('แก้ไขการใช้จ่ายสำเร็จ');
            fetchExpenses();
        } else {
            alert('การแก้ไขการใช้จ่ายล้มเหลว');
        }
    }
}

// ลบการใช้จ่าย (Delete)
async function deleteExpense(id) {
    const confirmation = confirm('คุณต้องการลบการใช้จ่ายนี้ใช่หรือไม่?');
    if (confirmation) {
        const response = await fetch(`/expenses/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': `Bearer ${token}` }
        });
        if (response.ok) {
            alert('ลบการใช้จ่ายสำเร็จ');
            fetchExpenses();
        } else {
            alert('การลบการใช้จ่ายล้มเหลว');
        }
    }
}

// เรียกใช้ฟังก์ชันดึงข้อมูลการใช้จ่ายเมื่อโหลดหน้าเสร็จ
fetchExpenses();
