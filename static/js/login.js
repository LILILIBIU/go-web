const form = document.getElementById('form');
const username = document.getElementById('username');
const password = document.getElementById('password');

form.addEventListener('submit', e => {
    e.preventDefault();
    if( checkInputs()!==1){
        return
    }

    form.onsubmit
});

function checkInputs() {
    // trim to remove the whitespaces
    const usernameValue = username.value.trim();
    const passwordValue = password.value.trim();

    if (usernameValue === '') {
        setErrorFor(username, '用户名不能为空');
        return 0
    } else {
        setSuccessFor(username);
    }


    if (passwordValue === '') {
        setErrorFor(password, '密码不能为空');
        return 0
    } else {
        setSuccessFor(password);
    }

    function setErrorFor(input, message) {
        const formControl = input.parentElement;
        const small = formControl.querySelector('small');
        formControl.className = 'form-control error';
        small.innerText = message;
    }

    function setSuccessFor(input) {
        const formControl = input.parentElement;
        formControl.className = 'form-control success';
    }
}