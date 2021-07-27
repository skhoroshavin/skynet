import auth from "../services/auth";
import Link from "next/link";
import useSWR from "swr"

export default function Home() {
    const userId = useSWR("me", auth.me).data

    return <>
        <div>Welcome to SkyNET, social network for cyborgs and all that stuff!</div>
        {userId == null
            ? <Link href="/signup">Signup</Link>
            : <div>Welcome back, {userId}!</div>}
    </>
}
