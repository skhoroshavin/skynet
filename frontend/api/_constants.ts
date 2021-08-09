import {IS_SERVER} from "../utils/constants";

export const SKYNET_API = IS_SERVER ?
    (process.env.SKYNET_API || "http://localhost:8080") :
    "/api"
