import {signup} from "./src/auth.js";
import {fakeUser} from "./src/fake";
import * as faker from "faker/locale/ru"
import {group} from "k6";

faker.seed(Date.now().valueOf() + __VU*1000000)

export default function () {
    const user = fakeUser()

    group("signup new user", () => {
        signup(user)
    })
}
