import http from "k6/http";
import {check} from "k6";
import {API_URL} from "./constants";

export function signup(user) {
    let res = http.post(`${API_URL}/auth/signup`, JSON.stringify({
        id: user.id,
        password: user.password
    }))
    check(res, {
        "returned status 200":
            (res) => res.status === 200,
        "provided session id in a cookie":
            (res) => (res.cookies.sessionid != null),

    })
    return res
}

export function me() {
    let res = http.get(`${API_URL}/auth/me`)
    check(res, {
        "me returned status 200":
            (res) => res.status === 200
    })
    return res
}