var ck_email = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
var ck_username = /^[A-Za-z0-9_]{1,20}$/;
var ck_password  = /^[A-Za-z0-9!@#$%^&*()_]{6,20}$/;

function validateEmail(email) {
    if (!ck_email.test(email)) {
        return false;
    }
    return true;
}

function validateUsername(username) {
    if (!ck_username.test(username)) {
        return false;
    }
    return true;
}

function validatePassword(password) {
    if (!ck_password.test(password)) {
        return false;
    }
    return true;
}