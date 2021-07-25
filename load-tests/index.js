import * as auth from "./src/auth.js";
import * as users from "./src/users.js";
import {fakeUser} from "./src/fake";
import * as faker from "faker/locale/ru"
import {check, group} from "k6";

faker.seed(Date.now().valueOf() + __VU*1000000)

export let options = {
  scenarios: {
    main: {
      executor: 'ramping-vus',
      stages: [
        { duration: '3s', target: 1 },
      ],
      gracefulRampDown: '0s',
    }
  }
}

export default function () {
    const user = fakeUser()

    group("signup new user", () => {
        auth.signup(user)
    })

    group("get current user", () => {
        let res = auth.me()
        check(res, {
            "me returned current user": res.json("id") === user.id
        })
    })

    group("fill in user data", () => {
        users.update(user)
    })

   group("get user data", () => {
       let res = users.get(user.id).json()

       // TODO: Check birthday
       check(res, {
           "first name is correct": res["first_name"] === user.firstName,
           "last name is correct": res["last_name"] === user.lastName,
           "gender is correct": res["gender"] === user.gender,
           "city is correct": res["city"] === user.city,
       })
   })
}
