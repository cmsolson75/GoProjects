fetchUsers()

fetch('/message')
    .then(res => res.json())
    .then(data => {
        document.getElementById("message").innerHTML = data.content;
    });
const formEl = document.querySelector('.add');
formEl.addEventListener('submit', event => {
    event.preventDefault();

    const formData = new FormData(formEl);
    var object = {};
    formData.forEach((value, key) => object[key] = value);
    var json = JSON.stringify(object);
    console.log(json)

    fetch('/user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: json
    })
    fetchUsers()

})

async function fetchUsers() {
    fetch("/user")
        .then(function (response) {
            return response.json()
        })
        .then(function (users) {
            let placeholder = document.querySelector("#data-output");
            let out = "";
            for (let user of users) {
                out += `
                <tr>
                    <td>${user.id}</td>
                    <td>${user.name}</td>
                    <td>${user.age}</td>
                </tr>
            `;
            }

            placeholder.innerHTML = out;
        })
}


async function fetchUserID() {
    try {
        const userID = document.getElementById("userID").value;
        console.log(userID);
        const response = await fetch(`/user/${userID}`);
        if (!response.ok) {
            throw new Error("could not fetch resource");
        }
        const data = await response.json();
        const name = data.name;
        document.getElementById("userName").innerHTML = name;
    }
    catch (error) {
        alert("element not found, try again");
        document.getElementById("userName").innerHTML = "";
        console.error(error);
    }
}

function updateUserByID() {
    try {
        const updateForm = document.querySelector(".update");
        const formData = new FormData(updateForm);
        var object = {};
        formData.forEach((value, key) => object[key] = value);
        var data = JSON.stringify(object);
        console.log(data)
        fetch('/user', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: data
        });
        fetchUsers()
    }
    catch (error) {

    }
    fetchUsers()
}



async function deleteUserID() {
    try {
        const userID = document.getElementById("deleteID").value;
        const myDataObject = { userID: userID }
        const data = JSON.stringify(myDataObject)
        console.log(data);
        fetch("/user", {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(myDataObject)
        });
        fetchUsers()
    }
    catch (error) {
        alert("element not found, try again");
    }
}