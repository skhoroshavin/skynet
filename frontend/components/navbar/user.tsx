import React, {useEffect, useState} from "react";
import Link from "next/link";
import {signOut, useCurrentUser} from "../../services/auth";
import {ChevronDownIcon} from "@heroicons/react/solid";

type NavUserMenuProps = {
    userId: string
}

const NavUserMenu = ({userId}: NavUserMenuProps) => {
    const [menuVisible, setMenuVisible] = useState(false)

    return <div className={`relative ml-auto h-full hover:bg-primary-700 ${menuVisible && "bg-primary-700"}`}>
        <button className="h-full px-4 font-bold -mr-1 relative z-10"
                type="button" onClick={() => setMenuVisible(!menuVisible)}>
            {userId}
            <ChevronDownIcon className="inline-block w-5"/>
        </button>

        <button className={`fixed h-full w-full inset-0 cursor-default ${menuVisible ? "block" : "hidden"}`}
                type="button" onClick={() => setMenuVisible(false)}/>

        <div className={menuVisible ? "absolute block right-0 bg-white rounded-b-lg shadow-lg w-36 py-2 z-10" : "hidden"}>
            <Link href={"/settings"} passHref>
                <a className="mt-2 block w-full text-left px-4 py-1 hover:bg-primary-500">
                    Settings
                </a>
            </Link>
            <hr className="mt-2 border-grey-300"/>
            <button className="mt-2 block w-full text-left px-4 py-1 hover:bg-primary-500"
                    type="button" onClick={signOut}>
                Sign out
            </button>
        </div>
    </div>
}

const NavSignIn = () => {
    return <Link href={"/signin"} passHref>
        <a className="ml-auto block h-full font-bold px-4 pt-5 hover:bg-primary-700">Sign In</a>
    </Link>
}

export const NavUser = () => {
    const [mounted, setMounted] = useState(false);
    const userId = useCurrentUser()

    useEffect(() => {
        setMounted(true);
    })

    if (!mounted) {
        return null
    }

    return userId ? <NavUserMenu userId={userId}/> : <NavSignIn/>;
}
