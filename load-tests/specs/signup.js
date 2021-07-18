import http from "k6/http";
import {alnum, stringOf, varStringOf} from "./_random.js";
import {check} from "k6";
import {API_URL} from "./_constants.js";

function _signup(id, password) {
    return http.post(`${API_URL}/auth/signup`, JSON.stringify({id, password}))
}

export function signup() {
    const id = varStringOf(8, 20, alnum)
    const password = varStringOf(3, 30, alnum)
    const res = _signup(id, password)
    check(res, {
        "signup returned status 200":
            (res) => res.status == 200,
        "signup provided session id in a cookie":
            (res) => res.cookies.sessionid != null
    })
    return res.cookies.sessionid
}
