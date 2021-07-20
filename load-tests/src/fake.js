import * as faker from "faker/locale/ru";

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

function userName(firstName, lastName) {
    var result;
    switch (faker.datatype.number(5)) {
        case 0:
            result = firstName + faker.datatype.number(999999);
            break;
        case 1:
        case 2:
        case 3:
            result = firstName + faker.random.arrayElement([".", "_", ""]) + lastName + faker.datatype.number(9999);
            break;
        case 4:
            result = firstName[0] + lastName + faker.datatype.number(999999);
            break;
        case 5:
            result = firstName + lastName[0] + faker.datatype.number(999999);
            break;
    }
    result = result.toString().replace(/'/g, "");
    result = result.replace(/ /g, "");
    return transliterate(result.toLowerCase());
}

export function fakeUser() {
    const locale = faker.random.arrayElement(['ru', 'en', 'en', 'en'])
    faker.setLocale(locale)

    const genderId = faker.datatype.number(1)

    const firstName = faker.name.firstName(genderId)
    const lastName = faker.name.lastName(genderId)
    const gender = genderId == 0 ? "male" : "female"
    const birthday = faker.date.between('1940-01-01', '2009-01-01')
    const city = faker.address.city()

    const id = userName(firstName, lastName)
    const password = faker.random.alphaNumeric(faker.datatype.number({min: 6, max: 20}))

    return {
        id: id,
        password: password,
        firstName: faker.datatype.number(10) < 9 ? firstName : null,
        lastName: faker.datatype.number(10) < 6 ? lastName : null,
        gender: faker.datatype.number(10) < 9 ? gender : null,
        birthday: faker.datatype.number(10) < 6 ? birthday : null,
        city: faker.datatype.number(10) < 5 ? city : null
    }
}
