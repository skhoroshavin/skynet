import React, {ReactFragment} from "react";
import Link from "next/link";
import {signOut, useCurrentUser} from "../../services/auth";

const navUserClass = "ml-auto h-full px-6 flex items-center font-bold hover:bg-primary-700"

type NavUserMenuProps = {
    userId: string
}

const NavUserMenu = ({userId}: NavUserMenuProps) => {
    return <button className={`${navUserClass}`}
                   type="button" onClick={signOut}>
        {userId}
    </button>
}

const NavSignIn = () => {
    return <Link href={"/signin"} passHref>
        <a className={`${navUserClass}`}>Sign In</a>
    </Link>
}

export const NavUser = () => {
    const userId = useCurrentUser()
    return userId ? <NavUserMenu userId={userId}/> : <NavSignIn/>
}

