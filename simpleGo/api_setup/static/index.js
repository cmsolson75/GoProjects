const apiRequest = async (url, method = 'GET', data = null) => {
    const options = { method, headers: { 'Content-Type': 'application/json' } };
    if (data) options.body = JSON.stringify(data);

    try {
        const response = await fetch(url, options);
        
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        if (response.status === 204 || response.headers.get('Content-Length') === '0') {
            return null;
        }
        return await response.json();
    } catch (error) {
        console.error(`Error during ${method} request to ${url}:`, error);
        throw error;
    }
}

const getMessage = async () => {
    try {
        const titleMessage = await apiRequest('/message');
        document.getElementById("message").innerHTML = titleMessage.content;
    } catch (error) {
        console.error("Failed to fetch message:", error);
    }
}

const fetchUsers = async () => {
    try {
        const users = await apiRequest('/user');
        const rows = users
            .map(user => `
            <tr>
                <td>${user.id}</td>
                <td>${user.name}</td>
                <td>${user.age}</td>
            </tr>
        `)
            .join('');
        // what is #data-output
        document.querySelector('#data-output').innerHTML = rows;
    } catch (error) {
        console.error("Failed to fetch users:", error);
    }
}

const fetchUserByID = async () => {
    const userID = document.getElementById('userID').value;
    if (!userID) {
        alert("Please enter a valid User ID!");
        return;
    }

    try {
        const user = await apiRequest(`/user/${userID}`);
        document.getElementById('userName').innerHTML = user.name;
    } catch (error) {
        alert("User not found.");
        document.getElementById('userName').innerText = '';
    }
}

const updateUser = async () => {
    const form = document.querySelector('.update');
    const formData = new FormData(form);
    const data = Object.fromEntries(formData);
    data.id = parseInt(data.id, 10);
    data.age = parseInt(data.age, 10);

    try {
        await apiRequest('/user', 'PUT', data);
        fetchUsers();
    } catch (error) {
        alert("Failed to update user.");
    }
}

const deleteUser = async () => {
    const userID = document.getElementById('deleteID').value;
    if (!userID) {
        alert("Please enter a valid User ID!");
        return;
    }

    try {
        await apiRequest(`/user/${userID}`, 'DELETE');
        alert("User deleted successfully!");
        fetchUsers();
    } catch (error) {
        alert("Failed to delete user.");
    }

}

const addUser = async () => {
    const formAdd = document.querySelector('.add');
    const formData = new FormData(formAdd);
    const data = Object.fromEntries(formData);
    data.age = parseInt(data.age, 10);

    try {
        await apiRequest('/user', 'POST', data);
        alert('User added successfully!');
        fetchUsers();
    } catch (error) {
        alert('Failed to add user');
    }
}

document.addEventListener('DOMContentLoaded', () => {
    getMessage();
    fetchUsers();
})

// Event Listeners
document.querySelector('.update').addEventListener('submit', (e) => {
    e.preventDefault();
    updateUser();
})

document.querySelector('.add').addEventListener('submit', (e) => {
    // is this normal to add
    e.preventDefault();
    addUser();
})