import {group} from "k6";
import {signup} from "./specs/signup.js";

export default function () {
    group("signup new user", () => {
        signup()
    })
}
