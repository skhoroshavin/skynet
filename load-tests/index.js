import {signup} from "./src/auth.js";
import * as faker from "faker/locale/ru"

faker.seed(__VU + (Date.now().valueOf() / 10000))

export default function () {
    signup()
}
