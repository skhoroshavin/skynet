const fc = require('fast-check');
const auth = require('../frontend/services/auth')

it("can signup new user", async () => {
    const err = await auth.signUp("xcsdf", "234")
    expect(err).toBe(null)
})

