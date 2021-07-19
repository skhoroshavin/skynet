import http from "k6/http";
import {check, group} from "k6";
import {API_URL, LOCALES} from "./_constants";
import * as faker from "faker/locale/ru"


function transliterate(word){
    let answer = '';
    const converter = {
        'а': 'a', 'б': 'b', 'в': 'v', 'г': 'g', 'д': 'd',
        'е': 'e', 'ё': 'e', 'ж': 'zh', 'з': 'z', 'и': 'i',
        'й': 'y', 'к': 'k', 'л': 'l', 'м': 'm', 'н': 'n',
        'о': 'o', 'п': 'p', 'р': 'r', 'с': 's', 'т': 't',
        'у': 'u', 'ф': 'f', 'х': 'h', 'ц': 'c', 'ч': 'ch',
        'ш': 'sh', 'щ': 'sch', 'ь': '', 'ы': 'y', 'ъ': '',
        'э': 'e', 'ю': 'yu', 'я': 'ya'
    };

    for (let i = 0; i < word.length; ++i ) {
        if (converter[word[i]] === undefined){
            answer += word[i];
        } else {
            answer += converter[word[i]];
        }
    }

    return answer;
}

const simplify = (s) => {
    return transliterate(s.toLowerCase().replace("'", ""))
}

export function signup() {
    const locale = faker.random.arrayElement(LOCALES)
    faker.setLocale(locale)
    const gender = faker.datatype.number(1)
    const firstName = faker.name.firstName(gender)
    const lastName = faker.name.lastName(gender)
    const id = `${simplify(firstName)}.${simplify(lastName)}${faker.random.alphaNumeric(4)}`
    const password = faker.random.alphaNumeric(faker.datatype.number({min: 6, max: 20}))
    let res
    group("signup new user", () => {
        res = http.post(`${API_URL}/auth/signup`, JSON.stringify({id, password}))
        check(res, {
            "signup returned status 200":
                (res) => res.status == 200,
            "signup provided session id in a cookie":
                (res) => res.cookies.sessionid != null
        })
    })
    return {id, password, sessionId: res.cookies.sessionid}
}
