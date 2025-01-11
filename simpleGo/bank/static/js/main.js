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



const submitEmail = async () => {
    const userEmail = document.getElementById("userEmail").value;
    if (!userEmail) {
        alert("Please enter a emial");
        return;
    }
    const data = { 'userEmail': userEmail};

    try {
        let response = await apiRequest('/validate', 'POST', data);
        console.log(response.id);
        console.log("Request sent");
        window.location.href = `/app/${response.id}`

    } catch (error) {
        alert("no email: need to work on this error");
    }
}