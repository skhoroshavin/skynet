import http from "k6/http";
import {check} from "k6";
import {API_URL} from "./constants";

export function update(user) {
    let res = http.put(`${API_URL}/users/${user.id}`, JSON.stringify({
        first_name: user.firstName,
        last_name: user.lastName,
        gender: user.gender,
        birthday: user.birthday,
        city: user.city,
    }))
    check(res, {
        "returned status 200":
            (res) => res.status == 200
    })
    return res
}
