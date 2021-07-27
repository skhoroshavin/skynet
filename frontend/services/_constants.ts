
export const IS_SERVER = typeof window === 'undefined'

export const SKYNET_API = IS_SERVER ?
    (process.env.SKYNET_API || "http://localhost:8080") :
    "/api"
