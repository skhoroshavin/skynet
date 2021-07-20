import {me, signup} from "./src/auth.js";
import {fakeUser} from "./src/fake";
import * as faker from "faker/locale/ru"
import {check, group} from "k6";
import http from "k6/http";
import {API_URL} from "./src/constants";

faker.seed(Date.now().valueOf() + __VU*1000000)

export default function () {
    const user = fakeUser()

    group("signup new user", () => {
        signup(user)
    })

    group("get current user", () => {
        let res = me()
        check(res, {
            "me returned current user": res.json("id") == user.id
        })
    })
}
